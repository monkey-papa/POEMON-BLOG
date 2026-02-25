/**
 * 请求缓存和去重工具
 * 用于优化频繁请求的接口性能
 */

import type { CacheConfig, CacheEntry, CacheStats } from "@/types";

// ==================== 缓存配置 ====================

/**
 * 缓存配置
 * key: 接口路径关键字
 * value: 缓存时间（毫秒）
 */
const CACHE_CONFIG: CacheConfig = {
  "/webInfo/getWebInfo": 5 * 60 * 1000, // 网站信息缓存 5 分钟
  "/webInfo/getSortInfo": 5 * 60 * 1000, // 分类信息缓存 5 分钟
  "/webInfo/listSortAndLabel": 5 * 60 * 1000, // 分类标签缓存 5 分钟
  "/webInfo/getClassifyList": 3 * 60 * 1000, // 分类统计缓存 3 分钟
  "/webInfo/listCollect": 5 * 60 * 1000, // 收藏列表缓存 5 分钟
};

// ==================== 缓存存储 ====================

/**
 * 缓存存储结构
 * Map<cacheKey, { data, timestamp, expireAt }>
 */
const cache = new Map<string, CacheEntry>();

/**
 * 正在进行的请求
 * Map<requestKey, Promise>
 * 用于请求去重
 */
const pendingRequests = new Map<string, Promise<unknown>>();

// ==================== 工具函数 ====================

/**
 * 生成缓存键
 * @param url - 请求 URL
 * @param params - 请求参数
 * @returns 缓存键
 */
function generateCacheKey(url: string, params: object = {}): string {
  const paramStr = JSON.stringify(params, Object.keys(params).sort());
  return `${url}::${paramStr}`;
}

/**
 * 检查接口是否需要缓存
 * @param url - 请求 URL
 * @returns 缓存时间（毫秒），null 表示不缓存
 */
function getCacheDuration(url: string): number | null {
  for (const [key, duration] of Object.entries(CACHE_CONFIG)) {
    if (url.includes(key)) {
      return duration;
    }
  }
  return null;
}

/**
 * 获取缓存数据
 * @param cacheKey - 缓存键
 * @returns 缓存数据，已过期或不存在返回 null
 */
function getCache<T = unknown>(cacheKey: string): T | null {
  const cached = cache.get(cacheKey);
  if (!cached) return null;

  if (Date.now() > cached.expireAt) {
    cache.delete(cacheKey);
    return null;
  }

  return cached.data as T;
}

/**
 * 设置缓存数据
 * @param cacheKey - 缓存键
 * @param data - 数据
 * @param duration - 缓存时间（毫秒）
 */
function setCache<T = unknown>(cacheKey: string, data: T, duration: number): void {
  cache.set(cacheKey, {
    data,
    timestamp: Date.now(),
    expireAt: Date.now() + duration,
  });
}

// ==================== 导出函数 ====================

/**
 * 带缓存和去重的请求包装器
 * @param url - 请求 URL
 * @param params - 请求参数
 * @param requestFn - 实际请求函数
 * @returns 请求结果
 */
export async function cachedRequest<T = unknown>(
  url: string,
  params: object,
  requestFn: () => Promise<T>
): Promise<T> {
  const cacheDuration = getCacheDuration(url);
  const cacheKey = generateCacheKey(url, params);

  // 1. 检查缓存
  if (cacheDuration) {
    const cachedData = getCache<T>(cacheKey);
    if (cachedData !== null) {
      return cachedData;
    }
  }

  // 2. 检查是否有相同的请求正在进行（去重）
  if (pendingRequests.has(cacheKey)) {
    return pendingRequests.get(cacheKey) as Promise<T>;
  }

  // 3. 发起新请求
  const requestPromise = (async (): Promise<T> => {
    try {
      const result = await requestFn();

      // 缓存结果
      if (cacheDuration) {
        setCache(cacheKey, result, cacheDuration);
      }

      return result;
    } finally {
      // 请求完成后移除 pending 状态
      pendingRequests.delete(cacheKey);
    }
  })();

  // 记录正在进行的请求
  pendingRequests.set(cacheKey, requestPromise);

  return requestPromise;
}

/**
 * 清除指定接口的缓存
 * @param urlPattern - URL 匹配模式
 */
export function clearCache(urlPattern: string): void {
  for (const key of cache.keys()) {
    if (key.includes(urlPattern)) {
      cache.delete(key);
    }
  }
}

/**
 * 清除所有缓存
 */
export function clearAllCache(): void {
  cache.clear();
}

/**
 * 获取缓存统计信息
 * @returns 缓存统计
 */
export function getCacheStats(): CacheStats {
  const stats: CacheStats = {
    totalEntries: cache.size,
    entries: [],
  };

  for (const [key, value] of cache.entries()) {
    stats.entries.push({
      key,
      expireAt: new Date(value.expireAt).toISOString(),
      remainingTime: Math.max(0, value.expireAt - Date.now()),
    });
  }

  return stats;
}

export default {
  cachedRequest,
  clearCache,
  clearAllCache,
  getCacheStats,
};

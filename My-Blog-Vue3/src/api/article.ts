import http from "../utils/request";
import constant from "../utils/constant";
import apiWrapper from "./tool/apiWrapper";
import type {
  ArticleListItem,
  ArticleDetail,
  AdminArticleDetail,
  GetArticleListParams,
  GetArticleDetailParams,
  SaveArticleParams,
  UpdateArticleParams,
  ChangeArticleStatusParams,
  GenerateSummaryParams,
  GenerateSummaryResponse,
  PaginationResponse,
} from "@/types";

// ==================== 文章相关 API ====================

/**
 * 获取文章列表（前台）
 * @param params - 分页参数 { current, size, sortId, labelId, searchKey }
 * @returns 文章列表
 */
export function getArticleList(
  params: GetArticleListParams
): Promise<PaginationResponse<ArticleListItem>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<ArticleListItem>>(`${constant.baseURL}/article/list`, {
        ...params,
        viewAll: false,
      }),
    "获取文章列表失败！"
  );
}

/**
 * 获取管理员文章列表（后台）
 * @param pagination - 分页参数
 * @returns 文章列表
 */
export function getAdminArticleList(
  pagination: GetArticleListParams
): Promise<PaginationResponse<ArticleListItem>> {
  return apiWrapper(
    () =>
      http.post<PaginationResponse<ArticleListItem>>(`${constant.baseURL}/article/list`, {
        ...pagination,
        viewAll: true,
      }),
    "获取文章列表失败！"
  );
}

/**
 * 获取文章详情（前台）
 * @param params - { id }
 * @returns 文章详情
 */
export function getArticleById(params: GetArticleDetailParams): Promise<ArticleDetail> {
  return apiWrapper(
    () =>
      http.post<ArticleDetail>(`${constant.baseURL}/article/detail`, {
        id: params.id,
        isAdmin: false, // 前台模式，后端自动过滤只返回公开文章
      }),
    "获取文章详情失败！"
  );
}

/**
 * 获取管理员文章详情（后台，精简字段用于编辑）
 * @param id - 文章ID
 * @returns 管理员文章详情
 */
export function getAdminArticleById(id: number): Promise<AdminArticleDetail> {
  return apiWrapper(
    () =>
      http.post<AdminArticleDetail>(`${constant.baseURL}/article/detail`, {
        id: id,
        isAdmin: true,
      }),
    "获取文章详情失败！"
  );
}

/**
 * 文章点赞/取消点赞
 * @param articleId - 文章ID
 */
export function articleLike(articleId: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/article/articleLike`, { articleId }),
    "点赞失败！",
    { allowEmpty: true }
  );
}

/**
 * 保存文章（新增）
 * @param article - 文章数据
 */
export function saveArticle(article: SaveArticleParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/article/save`, article),
    "保存文章失败！",
    { allowEmpty: true }
  );
}

/**
 * 更新文章
 * @param article - 文章数据（包含 id）
 */
export function updateArticle(article: UpdateArticleParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/article/update`, article),
    "更新文章失败！",
    { allowEmpty: true }
  );
}

/**
 * 删除文章
 * @param id - 文章ID
 */
export function deleteArticle(id: number): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/article/delete`, { id }),
    "删除文章失败！",
    { allowEmpty: true }
  );
}

/**
 * 修改文章状态
 * @param param - 状态参数 { id, viewStatus, commentStatus, recommendStatus }
 * @returns 更新结果
 */
export function changeArticleStatus(param: ChangeArticleStatusParams): Promise<void> {
  return apiWrapper(
    () => http.post<void>(`${constant.baseURL}/article/changeStatus`, param),
    "修改状态失败！",
    { allowEmpty: true }
  );
}

/**
 * 生成文章摘要（非流式，兼容已有摘要直接返回的场景）
 * @param data - { message, articleId }
 * @returns 摘要结果
 */
export function generateSummary(data: GenerateSummaryParams): Promise<GenerateSummaryResponse> {
  return apiWrapper(
    () => http.post<GenerateSummaryResponse>(`${constant.baseURL}/summary`, data),
    "生成摘要失败！"
  );
}

/**
 * 流式生成文章摘要（SSE）
 * 后端返回 SSE 流，已有摘要时返回普通 JSON
 * @param data - { message, articleId }
 * @param onChunk - 每收到一段文本时的回调
 * @returns 完整摘要文本；如果后端返回已缓存摘要，直接返回
 */
export async function generateSummaryStream(
  data: GenerateSummaryParams,
  onChunk: (text: string) => void
): Promise<string> {
  const { useStore } = await import("@/stores");
  const store = useStore();
  const isAdmin = window.location.hash?.includes("/admin");
  const token =
    (isAdmin ? store?.currentAdmin?.accessToken : store?.currentUser?.accessToken) || "";

  const response = await fetch(`${constant.baseURL}/summary`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: token ? "Token " + token : "",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error(`摘要请求失败: ${response.status}`);
  }

  const contentType = response.headers.get("content-type") || "";

  // 后端已有缓存摘要时返回普通 JSON
  if (contentType.includes("application/json")) {
    const json = await response.json();
    const summary = json?.data?.summary || "";
    if (summary) {
      onChunk(summary);
    }
    return summary;
  }

  // SSE 流式读取
  if (!response.body) {
    throw new Error("浏览器不支持 ReadableStream");
  }

  const reader = response.body.getReader();
  const decoder = new TextDecoder();
  let fullText = "";
  let buffer = "";

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    buffer += decoder.decode(value, { stream: true });

    // 按行解析 SSE
    const lines = buffer.split("\n");
    // 保留最后一个可能不完整的行
    buffer = lines.pop() || "";

    for (const line of lines) {
      if (!line.startsWith("data: ")) continue;
      const payload = line.slice(6).trim();

      // 结束信号
      if (payload === "[DONE]") {
        return fullText;
      }

      try {
        const parsed = JSON.parse(payload) as { text?: string };
        if (parsed.text) {
          fullText += parsed.text;
          onChunk(parsed.text);
        }
      } catch {
        // 非 JSON 行，忽略
      }
    }
  }

  return fullText;
}

export default {
  getArticleList,
  getArticleById,
  articleLike,
  saveArticle,
  updateArticle,
  deleteArticle,
  getAdminArticleList,
  getAdminArticleById,
  changeArticleStatus,
  generateSummary,
  generateSummaryStream,
};

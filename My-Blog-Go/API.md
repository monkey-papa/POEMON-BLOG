# POEMON Blog API 接口文档

## 概述

- **基础路径**: `/api`
- **请求方式**: 全部使用 `POST`
- **请求格式**: `application/json`（文件上传除外）
- **认证方式**: `Authorization: Bearer <token>` 或 `Authorization: Token <token>`

## 通用响应格式

### 成功响应

```json
{
  "code": 0,
  "msg": "操作成功",
  "data": { ... },
  "error": false,
  "success": true
}
```

### 分页响应

```json
{
  "code": 0,
  "msg": "操作成功",
  "data": {
    "totalCount": 100,
    "pageSize": 10,
    "totalPage": 10,
    "page": 1,
    "list": [...],
    "empty": false
  },
  "error": false,
  "success": true
}
```

### 失败响应

```json
{
  "code": 500,
  "msg": "错误信息",
  "data": null,
  "error": true,
  "success": false
}
```

---

# 一、公开接口（无需登录）

## 1. 前台登录

**POST** `/api/user/login`

### 请求参数

| 参数     | 类型   | 必填 | 说明                       |
| -------- | ------ | ---- | -------------------------- |
| account  | string | ✅   | 账号（用户名或邮箱）       |
| password | string | ✅   | 密码                       |
| province | string |      | 省份信息（首次登录时记录） |

### 响应体

```json
{
  "data": [
    {
      "accessToken": "eyJhbGciOiJIUzI1NiIs...",
      "id": 1,
      "username": "admin",
      "phoneNumber": "13800138000",
      "email": "admin@example.com",
      "admire": "https://...",
      "userStatus": true,
      "avatar": "https://...",
      "gender": 1,
      "introduction": "个人简介",
      "userType": 0,
      "createTime": "2024-01-01T00:00:00Z",
      "qiniuDomain": "cdn.example.com",
      "qiniuBucketName": "bucket",
      "qiniuSecretKey": "***",
      "qiniuAccessKey": "***"
    }
  ]
}
```

---

## 2. 后台登录

**POST** `/api/admin/user/login`

### 请求参数

| 参数     | 类型   | 必填 | 说明         |
| -------- | ------ | ---- | ------------ |
| account  | string | ✅   | 管理员用户名 |
| password | string | ✅   | 密码         |

### 响应体

```json
{
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "avatar": "https://...",
    "userType": 0,
    "username": "admin",
    "id": 1,
    "qiniuDomain": "cdn.example.com",
    "qiniuBucketName": "bucket",
    "qiniuSecretKey": "***",
    "qiniuAccessKey": "***"
  }
}
```

---

## 3. 用户注册

**POST** `/api/user/registration`

### 请求参数

| 参数     | 类型   | 必填 | 说明                                 |
| -------- | ------ | ---- | ------------------------------------ |
| username | string | ✅   | 用户名（字母数字下划线，最大50字符） |
| password | string | ✅   | 密码（6-128字符）                    |
| email    | string | ✅   | 邮箱地址                             |
| code     | string | ✅   | 邮箱验证码（6位）                    |
| province | string |      | 省份信息                             |

### 响应体

同「前台登录」

---

## 4. 忘记密码

**POST** `/api/user/updateForForgetPassword`

### 请求参数

| 参数     | 类型   | 必填 | 说明       |
| -------- | ------ | ---- | ---------- |
| username | string | ✅   | 用户名     |
| password | string | ✅   | 新密码     |
| email    | string | ✅   | 邮箱地址   |
| code     | string | ✅   | 邮箱验证码 |

### 响应体

```json
{
  "msg": "密码重置成功"
}
```

---

## 5. 发送验证码

**POST** `/api/code`

### 请求参数

| 参数  | 类型   | 必填 | 说明     |
| ----- | ------ | ---- | -------- |
| email | string | ✅   | 邮箱地址 |

### 响应体

```json
{
  "msg": "验证码已发送"
}
```

---

## 6. 用户列表

**POST** `/api/user/list`

> 通过 `isAdmin` 参数区分前台/后台视图

### 请求参数

| 参数       | 类型   | 必填 | 说明                           |
| ---------- | ------ | ---- | ------------------------------ |
| current    | int    |      | 当前页码，默认1                |
| size       | int    |      | 每页条数，默认10               |
| searchKey  | string |      | 搜索关键词（手机号）           |
| userStatus | string |      | 状态筛选："true"/"false"       |
| userType   | string |      | 用户类型筛选                   |
| isAdmin    | bool   |      | 是否管理员视图（返回完整信息） |

### 响应体（isAdmin=false，前台视图）

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "username": "user1",
        "avatar": "https://...",
        "gender": 1
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

### 响应体（isAdmin=true，后台视图）

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "username": "user1",
        "phoneNumber": "13800138000",
        "email": "user@example.com",
        "admire": "https://...",
        "userStatus": true,
        "avatar": "https://...",
        "gender": 1,
        "introduction": "个人简介",
        "userType": 2,
        "createTime": "2024-01-01T00:00:00Z",
        "province": "广东"
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 7. 文章列表

**POST** `/api/article/list`

> 通过 `viewAll` 参数区分是否查看所有文章（包括隐藏）

### 请求参数

| 参数            | 类型   | 必填 | 说明                         |
| --------------- | ------ | ---- | ---------------------------- |
| current         | int    |      | 当前页码                     |
| size            | int    |      | 每页条数                     |
| searchKey       | string |      | 搜索关键词（标题）           |
| recommendStatus | bool   |      | 推荐状态筛选                 |
| labelId         | uint   |      | 标签ID筛选                   |
| sortId          | uint   |      | 分类ID筛选                   |
| viewAll         | bool   |      | 是否查看所有（需管理员权限） |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "userId": 1,
        "username": "admin",
        "sortId": 1,
        "labelId": 1,
        "articleCover": "https://...",
        "articleTitle": "文章标题",
        "articleContent": "文章内容...",
        "viewCount": 100,
        "likeCount": 50,
        "viewStatus": true,
        "password": null,
        "recommendStatus": true,
        "commentStatus": true,
        "createTime": "2024-01-01T00:00:00Z",
        "updateTime": "2024-01-02T00:00:00Z",
        "updateBy": "admin",
        "commentCount": 10,
        "label": {
          "id": 1,
          "labelName": "Go语言",
          "labelDescription": "Go语言相关"
        },
        "sort": {
          "id": 1,
          "sortName": "技术",
          "sortDescription": "技术文章"
        }
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 8. 文章详情

**POST** `/api/article/detail`

> 通过 `isAdmin` 参数区分前台/后台视图

### 请求参数

| 参数     | 类型   | 必填 | 说明                           |
| -------- | ------ | ---- | ------------------------------ |
| id       | uint   | ✅   | 文章ID                         |
| userId   | uint   |      | 用户ID（用于判断点赞状态）     |
| flag     | string |      | "false"表示查看隐藏文章        |
| password | string |      | 文章访问密码                   |
| isAdmin  | bool   |      | 是否管理员视图（跳过密码验证） |

### 响应体（isAdmin=false，前台视图）

```json
{
  "data": {
    "id": 1,
    "summary": "文章摘要...",
    "userId": 1,
    "sortId": 1,
    "labelId": 1,
    "articleCover": "https://...",
    "articleTitle": "文章标题",
    "articleContent": "文章内容...",
    "articleAuthor": "admin",
    "viewStatus": true,
    "password": null,
    "recommendStatus": true,
    "commentStatus": true,
    "createTime": "2024-01-01T00:00:00Z",
    "likeCount": 50,
    "updateBy": "admin",
    "updateTime": "2024-01-02T00:00:00Z",
    "username": "admin",
    "commentCount": 10,
    "viewCount": 101,
    "label": {...},
    "sort": {...},
    "articleLikeStatus": 0
  }
}
```

### 响应体（isAdmin=true，后台视图）

```json
{
  "data": {
    "id": 1,
    "userId": 1,
    "sortId": 1,
    "labelId": 1,
    "articleCover": "https://...",
    "articleTitle": "文章标题",
    "articleContent": "文章内容...",
    "articleAuthor": "admin",
    "viewStatus": true,
    "password": "secret",
    "recommendStatus": true,
    "commentStatus": true
  }
}
```

---

## 9. 获取文章摘要

**POST** `/api/summary`

### 请求参数

| 参数      | 类型   | 必填 | 说明                       |
| --------- | ------ | ---- | -------------------------- |
| message   | string | ✅   | 文章内容                   |
| articleId | uint   |      | 文章ID（已有摘要直接返回） |

### 响应体

```json
{
  "data": {
    "summary": "AI生成的文章摘要..."
  }
}
```

---

## 10. 获取分类信息

**POST** `/api/webInfo/getSortInfo`

### 请求参数

无

### 响应体

```json
{
  "data": [
    {
      "id": 1,
      "sortName": "技术",
      "sortDescription": "技术文章",
      "sortType": 0,
      "priority": 1,
      "countOfSort": 50,
      "lengthOfLabel": 5,
      "labels": [
        {
          "id": 1,
          "labelName": "Go语言",
          "labelDescription": "Go语言相关",
          "sortId": 1,
          "countOfLabel": 10
        }
      ]
    }
  ]
}
```

---

## 11. 获取网站信息

**POST** `/api/webInfo/getWebInfo`

> 通过 `isAdmin` 参数区分视图

### 请求参数

| 参数    | 类型 | 必填 | 说明                           |
| ------- | ---- | ---- | ------------------------------ |
| isAdmin | bool |      | 是否管理员视图（返回全部配置） |

### 响应体

```json
{
  "data": [
    {
      "id": 1,
      "webName": "POEMON Blog",
      "webTitle": "个人博客",
      "notices": "公告内容",
      "footer": "页脚内容",
      "backgroundImage": "https://...",
      "avatar": "https://...",
      "randomAvatar": "https://...",
      "randomName": "随机名称配置",
      "randomCover": "https://...",
      "waifuJson": "{...}",
      "status": true
    }
  ]
}
```

---

## 12. 评论列表（前台）

**POST** `/api/comment/listComment`

### 请求参数

| 参数           | 类型   | 必填 | 说明                           |
| -------------- | ------ | ---- | ------------------------------ |
| current        | int    |      | 当前页码                       |
| size           | int    |      | 每页条数                       |
| csize          | int    |      | 子评论每页条数                 |
| source         | int    |      | 来源ID（文章ID）               |
| commentType    | string |      | 评论类型                       |
| floorCommentId | uint   |      | 楼层评论ID（查询子评论时使用） |

### 响应体

```json
{
  "data": {
    "total": 100,
    "parenttotal": 80,
    "data": [{
      "id": 1,
      "source": 1,
      "type": "article",
      "parentCommentId": null,
      "userId": 1,
      "username": "user1",
      "avatar": "https://...",
      "userType": 2,
      "floorCommentId": null,
      "parentUserId": null,
      "likeCount": 5,
      "commentContent": "评论内容",
      "commentInfo": null,
      "createTime": "2024-01-01T00:00:00Z",
      "childComments": [{
        "current": 1,
        "records": [...],
        "total": 10
      }]
    }]
  }
}
```

---

## 13. 表白墙列表

**POST** `/api/family/list`

> 通过 `isAdmin` 参数区分视图

### 请求参数

| 参数    | 类型 | 必填 | 说明                            |
| ------- | ---- | ---- | ------------------------------- |
| current | int  |      | 当前页码（管理员模式）          |
| size    | int  |      | 每页条数（管理员模式）          |
| status  | bool |      | 状态筛选（管理员模式）          |
| isAdmin | bool |      | 是否管理员模式（分页+查看全部） |

### 响应体（isAdmin=false，前台视图）

```json
{
  "data": [
    {
      "id": 1,
      "userId": 1,
      "bgCover": "https://...",
      "manCover": "https://...",
      "womanCover": "https://...",
      "manName": "他",
      "womanName": "她",
      "timing": "2020-01-01",
      "countdownTitle": "结婚倒计时",
      "countdownTime": "2025-01-01",
      "status": true,
      "familyInfo": "表白内容",
      "likeCount": 100,
      "createTime": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### 响应体（isAdmin=true，后台视图）

```json
{
  "data": {
    "list": [...],
    "totalCount": 10,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 14. 获取站长表白墙

**POST** `/api/family/getAdminFamily`

### 请求参数

无

### 响应体

同「表白墙列表」单条数据

---

## 15. 微言列表

**POST** `/api/weiYan/listWeiYan`

### 请求参数

| 参数    | 类型 | 必填 | 说明                       |
| ------- | ---- | ---- | -------------------------- |
| current | int  |      | 当前页码                   |
| size    | int  |      | 每页条数                   |
| userId  | uint |      | 用户ID筛选                 |
| all     | bool |      | 是否显示所有（包括非公开） |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "userId": 1,
        "username": "admin",
        "avatar": "https://...",
        "likeCount": 10,
        "content": "微言内容",
        "type": "a",
        "source": null,
        "isPublic": true,
        "createTime": "2024-01-01T00:00:00Z"
      }
    ],
    "totalCount": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 16. 文章进展列表

**POST** `/api/weiYan/listNews`

### 请求参数

| 参数    | 类型 | 必填 | 说明         |
| ------- | ---- | ---- | ------------ |
| current | int  |      | 当前页码     |
| size    | int  |      | 每页条数     |
| source  | int  |      | 文章ID       |
| all     | bool |      | 是否显示所有 |

### 响应体

同「微言列表」

---

## 17. 资源列表

**POST** `/api/resource/listResource`

### 请求参数

| 参数         | 类型   | 必填 | 说明         |
| ------------ | ------ | ---- | ------------ |
| current      | int    |      | 当前页码     |
| size         | int    |      | 每页条数     |
| resourceType | string |      | 资源类型筛选 |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "createTime": "2024-01-01T00:00:00Z",
        "mimeType": "image/jpeg",
        "path": "https://...",
        "size": 102400,
        "status": true,
        "type": "article",
        "userId": 1,
        "username": "admin"
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 18. 资源路径列表

**POST** `/api/webInfo/listResourcePath`

> 通过 `isAdmin` 参数区分视图

### 请求参数

| 参数         | 类型   | 必填 | 说明                       |
| ------------ | ------ | ---- | -------------------------- |
| current      | int    |      | 当前页码                   |
| size         | int    |      | 每页条数                   |
| resourceType | string |      | 资源类型筛选               |
| classify     | string |      | 分类筛选                   |
| status       | bool   |      | 状态筛选（管理员模式）     |
| isAdmin      | bool   |      | 是否管理员模式（查看全部） |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "classify": "工具",
        "createTime": "2024-01-01T00:00:00Z",
        "introduction": "简介",
        "cover": "https://...",
        "friendAvatar": "https://...",
        "status": true,
        "title": "标题",
        "type": "favorites",
        "url": "https://..."
      }
    ],
    "totalCount": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 19. 分类统计

**POST** `/api/webInfo/getClassifyList`

### 请求参数

| 参数 | 类型   | 必填 | 说明     |
| ---- | ------ | ---- | -------- |
| type | string |      | 类型筛选 |

### 响应体

```json
{
  "data": [
    {
      "classify": "工具",
      "count": 10
    }
  ]
}
```

---

## 20. 收藏列表

**POST** `/api/webInfo/listCollect`

### 请求参数

无

### 响应体

```json
{
  "data": {
    "工具": [
      {
        "id": 1,
        "classify": "工具",
        "createTime": "2024-01-01T00:00:00Z",
        "introduction": "简介",
        "cover": "https://...",
        "remark": "备注",
        "status": true,
        "title": "标题",
        "type": "favorites",
        "url": "https://..."
      }
    ]
  }
}
```

---

## 21. 树洞列表

**POST** `/api/admin/treeHole/boss/list`

### 请求参数

| 参数    | 类型 | 必填 | 说明                  |
| ------- | ---- | ---- | --------------------- |
| current | int  |      | 当前页码              |
| size    | int  |      | 每页条数（0表示全部） |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "username": "匿名用户",
        "createTime": "2024-01-01T00:00:00Z",
        "message": "树洞内容",
        "avatar": "https://..."
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 22. 禁用词列表

**POST** `/api/prohibitedWords/list`

### 请求参数

| 参数      | 类型   | 必填 | 说明       |
| --------- | ------ | ---- | ---------- |
| current   | int    |      | 当前页码   |
| size      | int    |      | 每页条数   |
| searchKey | string |      | 搜索关键词 |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "username": "admin",
        "message": "禁用词",
        "avatar": "https://..."
      }
    ],
    "totalCount": 50,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 23. 获取客户端IP

**POST** `/api/ip`

### 请求参数

无

### 响应体

```json
{
  "data": {
    "ip": "192.168.1.1"
  }
}
```

---

## 24. 获取IP定位和天气

**POST** `/api/ip/location`

### 请求参数

无（通过服务端IP自动定位）

### 响应体

```json
{
  "data": {
    "ip": "192.168.1.1",
    "city": "深圳",
    "province": "广东",
    "weather": [
      {
        "type": "晴",
        "high": "28℃",
        "low": "20℃",
        "fengli": "3级",
        "fengxiang": "东南风"
      }
    ],
    "tip": "天气提示信息"
  }
}
```

---

## 25. IP访问统计

**POST** `/api/list/ip`

### 请求参数

无

### 响应体

```json
{
  "data": {
    "total_sum": 10000,
    "today_sum": 100,
    "yesterday_sum": 120,
    "data": {
      "province_all_top": [{"province": "广东", "count": 1000}],
      "ip_all_top": [{"ip": "192.168.1.1", "count": 500}],
      "today": {
        "province": [...],
        "ip": [...],
        "users": [{"userId": 1, "avatar": "...", "userName": "admin"}]
      },
      "yesterday": {...}
    }
  }
}
```

---

## 26. 记录IP访问

**POST** `/api/submit`

### 请求参数

| 参数     | 类型   | 必填 | 说明                     |
| -------- | ------ | ---- | ------------------------ |
| userId   | int    |      | 用户ID                   |
| ip       | string |      | IP地址（不传则自动获取） |
| province | string |      | 省份                     |
| city     | string |      | 城市                     |

### 响应体

```json
{
  "msg": "记录成功"
}
```

---

## 27. 地图数据

**POST** `/api/map`

### 请求参数

无

### 响应体

返回地图JSON数据

---

## 28. 用户地图数据

**POST** `/api/user/map`

### 请求参数

无

### 响应体

```json
{
  "data": [
    {
      "name": "广东省",
      "value": "100"
    }
  ]
}
```

---

# 二、需要登录的接口

## 29. 更新用户信息

**POST** `/api/user/updateUserInfo`

### 请求参数

| 参数            | 类型   | 必填 | 说明                       |
| --------------- | ------ | ---- | -------------------------- |
| userId          | uint   | ✅   | 用户ID                     |
| username        | string |      | 用户名                     |
| gender          | int    |      | 性别（0-未知，1-男，2-女） |
| introduction    | string |      | 个人简介                   |
| phoneNumber     | string |      | 手机号                     |
| email           | string |      | 邮箱                       |
| avatar          | string |      | 头像URL                    |
| code            | string |      | 验证码（修改邮箱时需要）   |
| qiniuDomain     | string |      | 七牛云域名                 |
| qiniuAccessKey  | string |      | 七牛云AccessKey            |
| qiniuSecretKey  | string |      | 七牛云SecretKey            |
| qiniuBucketName | string |      | 七牛云Bucket名称           |

### 响应体

```json
{
  "data": {
    "id": 1,
    "username": "admin",
    "phoneNumber": "13800138000",
    "email": "admin@example.com",
    "admire": "https://...",
    "userStatus": true,
    "avatar": "https://...",
    "gender": 1,
    "introduction": "个人简介",
    "userType": 0,
    "createTime": "2024-01-01T00:00:00Z",
    "qiniuDomain": "cdn.example.com",
    "qiniuBucketName": "bucket",
    "qiniuSecretKey": "***",
    "qiniuAccessKey": "***"
  }
}
```

---

## 30. 保存资源记录

**POST** `/api/resource/saveResource`

### 请求参数

| 参数     | 类型   | 必填 | 说明             |
| -------- | ------ | ---- | ---------------- |
| id       | uint   | ✅   | 用户ID           |
| type     | string | ✅   | 资源类型         |
| path     | string | ✅   | 资源路径         |
| size     | int    | ✅   | 文件大小（字节） |
| mimeType | string | ✅   | MIME类型         |

### 响应体

```json
{
  "msg": "保存成功"
}
```

---

## 31. 上传图片

**POST** `/api/resource/updateImage`

> 使用 `multipart/form-data` 格式

### 请求参数

| 参数   | 类型   | 必填 | 说明                     |
| ------ | ------ | ---- | ------------------------ |
| userId | string | ✅   | 用户ID（表单字段）       |
| image  | file   | ✅   | 图片文件                 |
| folder | string |      | 存储文件夹，默认"images" |

### 响应体

```json
{
  "data": {
    "url": "https://cdn.example.com/images/xxx.jpg"
  }
}
```

---

## 32. 添加评论

**POST** `/api/admin/comment/boss/addComment`

### 请求参数

| 参数            | 类型   | 必填 | 说明               |
| --------------- | ------ | ---- | ------------------ |
| source          | int    | ✅   | 来源ID（文章ID等） |
| type            | string |      | 评论类型           |
| parentCommentId | uint   |      | 父评论ID           |
| userId          | uint   | ✅   | 用户ID             |
| floorCommentId  | uint   |      | 楼层评论ID         |
| parentUserId    | uint   |      | 被回复用户ID       |
| commentContent  | string | ✅   | 评论内容           |
| commentInfo     | string |      | 附加信息           |

### 响应体

```json
{
  "msg": "评论成功"
}
```

---

## 33. 添加表白墙

**POST** `/api/family/addFamily`

### 请求参数

| 参数           | 类型   | 必填 | 说明         |
| -------------- | ------ | ---- | ------------ |
| userId         | uint   | ✅   | 用户ID       |
| bgCover        | string | ✅   | 背景图URL    |
| manCover       | string |      | 男方头像URL  |
| womanCover     | string |      | 女方头像URL  |
| manName        | string |      | 男方名称     |
| womanName      | string |      | 女方名称     |
| timing         | string |      | 计时开始时间 |
| countdownTitle | string |      | 倒计时标题   |
| countdownTime  | string |      | 倒计时时间   |
| familyInfo     | string |      | 表白内容     |
| status         | bool   |      | 状态         |
| likeCount      | int    |      | 点赞数       |

### 响应体

```json
{
  "msg": "添加成功"
}
```

---

## 34. 删除微言

**POST** `/api/weiYan/deleteWeiYan`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 微言ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 35. 保存微言

**POST** `/api/weiYan/save`

> 有 `source` 参数则为文章进展，否则为普通微言

### 请求参数

| 参数     | 类型   | 必填 | 说明                           |
| -------- | ------ | ---- | ------------------------------ |
| userId   | uint   | ✅   | 用户ID                         |
| content  | string | ✅   | 内容                           |
| type     | string |      | 类型，默认"a"                  |
| source   | int    |      | 关联文章ID（有值则为文章进展） |
| isPublic | bool   |      | 是否公开（文章进展默认公开）   |

### 响应体

```json
{
  "msg": "保存成功"
}
```

---

## 36. 保存资源路径

**POST** `/api/webInfo/saveResourcePath`

### 请求参数

| 参数         | 类型   | 必填 | 说明     |
| ------------ | ------ | ---- | -------- |
| title        | string | ✅   | 标题     |
| classify     | string |      | 分类     |
| cover        | string |      | 封面URL  |
| url          | string |      | 链接     |
| introduction | string |      | 简介     |
| type         | string |      | 类型     |
| friendAvatar | string |      | 友链头像 |

### 响应体

```json
{
  "msg": "保存成功"
}
```

---

## 37. 保存树洞

**POST** `/api/webInfo/saveTreeHole`

### 请求参数

| 参数     | 类型   | 必填 | 说明                    |
| -------- | ------ | ---- | ----------------------- |
| avatar   | string | ✅   | 头像URL                 |
| username | string |      | 用户名                  |
| message  | string | ✅   | 留言内容（最大500字符） |

### 响应体

```json
{
  "data": {
    "id": 1,
    "username": "匿名",
    "createTime": "2024-01-01T00:00:00Z",
    "message": "树洞内容",
    "avatar": "https://..."
  }
}
```

---

## 38. 文章点赞

**POST** `/api/article/articleLike`

> 需要登录，用户ID从Token获取，自动切换点赞状态

### 请求参数

| 参数      | 类型 | 必填 | 说明   |
| --------- | ---- | ---- | ------ |
| articleId | uint | ✅   | 文章ID |

### 响应体

```json
{
  "msg": "点赞成功" // 或 "取消点赞成功"
}
```

---

## 39. 发送评论通知

**POST** `/api/codeComment`

### 请求参数

| 参数    | 类型   | 必填 | 说明                      |
| ------- | ------ | ---- | ------------------------- |
| comment | string | ✅   | 评论内容                  |
| name    | string | ✅   | 评论者名称                |
| email   | string | ✅   | 接收者邮箱                |
| type    | string |      | 通知类型：comment/approve |

### 响应体

```json
{
  "msg": "通知已发送"
}
```

---

# 三、管理员接口

## 40. 更新网站信息

**POST** `/api/admin/webInfo/updateAdminWebInfo`

### 请求参数

| 参数            | 类型   | 必填 | 说明                       |
| --------------- | ------ | ---- | -------------------------- |
| id              | uint   | ✅   | 记录ID                     |
| webName         | string |      | 网站名称（有值则完整更新） |
| webTitle        | string |      | 网站标题                   |
| notices         | string |      | 公告                       |
| footer          | string |      | 页脚                       |
| backgroundImage | string |      | 背景图URL                  |
| avatar          | string |      | 头像URL                    |
| randomAvatar    | string |      | 随机头像配置               |
| randomName      | string |      | 随机名称配置               |
| randomCover     | string |      | 随机封面配置               |
| waifuJson       | string |      | 看板娘配置                 |
| status          | bool   |      | 状态                       |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 41. 修改用户属性

**POST** `/api/admin/user/updateAttribute`

### 请求参数

| 参数     | 类型   | 必填 | 说明                                                     |
| -------- | ------ | ---- | -------------------------------------------------------- |
| userId   | uint   | ✅   | 用户ID                                                   |
| field    | string | ✅   | 要修改的字段：status/admire/type                         |
| value    | string |      | 字符串值（用于admire）                                   |
| intValue | int    |      | 整数值（用于type：0-Boss，1-管理员，2-普通用户，3-访客） |
| boolFlag | bool   |      | 布尔值（用于status）                                     |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 42. 保存或更新分类

**POST** `/api/webInfo/saveOrUpdateSort`

> 有 `id` 则更新，无 `id` 则新增

### 请求参数

| 参数            | 类型   | 必填 | 说明                       |
| --------------- | ------ | ---- | -------------------------- |
| id              | uint   |      | 分类ID（有则更新）         |
| sortName        | string | ✅   | 分类名称                   |
| sortDescription | string | ✅   | 分类描述                   |
| sortType        | int    |      | 分类类型：0-导航栏，1-普通 |
| priority        | int    | ✅   | 优先级                     |

### 响应体

```json
{
  "msg": "保存成功" // 或 "更新成功"
}
```

---

## 43. 删除分类

**POST** `/api/webInfo/deleteSort`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 分类ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 44. 保存或更新标签

**POST** `/api/webInfo/saveOrUpdateLabel`

> 有 `id` 则更新，无 `id` 则新增

### 请求参数

| 参数             | 类型   | 必填 | 说明               |
| ---------------- | ------ | ---- | ------------------ |
| id               | uint   |      | 标签ID（有则更新） |
| sortId           | uint   | ✅   | 所属分类ID         |
| labelName        | string | ✅   | 标签名称           |
| labelDescription | string | ✅   | 标签描述           |

### 响应体

```json
{
  "msg": "保存成功" // 或 "更新成功"
}
```

---

## 45. 删除标签

**POST** `/api/webInfo/deleteLabel`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 标签ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 46. 获取分类和标签列表

**POST** `/api/webInfo/listSortAndLabel`

### 请求参数

无

### 响应体

```json
{
  "data": {
    "sorts": [
      {
        "countOfSort": 50,
        "id": 1,
        "sortName": "技术",
        "sortType": 0,
        "sortDescription": "技术文章",
        "priority": 1
      }
    ],
    "labels": [
      {
        "countOfLabel": 10,
        "id": 1,
        "labelDescription": "Go语言相关",
        "labelName": "Go语言",
        "sortId": 1
      }
    ]
  }
}
```

---

## 47. 删除文章

**POST** `/api/article/delete`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 文章ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 48. 保存文章

**POST** `/api/article/save`

### 请求参数

| 参数            | 类型   | 必填 | 说明                    |
| --------------- | ------ | ---- | ----------------------- |
| articleTitle    | string | ✅   | 文章标题（最大200字符） |
| userId          | uint   | ✅   | 作者ID                  |
| articleAuthor   | string |      | 作者名                  |
| articleContent  | string |      | 文章内容                |
| recommendStatus | bool   |      | 推荐状态                |
| viewStatus      | bool   |      | 可见状态                |
| commentStatus   | bool   |      | 评论状态                |
| password        | string |      | 访问密码                |
| articleCover    | string |      | 封面图URL               |
| sortId          | uint   |      | 分类ID                  |
| labelId         | uint   |      | 标签ID                  |

### 响应体

```json
{
  "msg": "保存成功"
}
```

---

## 49. 更新文章

**POST** `/api/article/update`

### 请求参数

| 参数            | 类型   | 必填 | 说明      |
| --------------- | ------ | ---- | --------- |
| id              | uint   | ✅   | 文章ID    |
| articleTitle    | string |      | 文章标题  |
| userId          | uint   |      | 作者ID    |
| articleAuthor   | string |      | 作者名    |
| articleContent  | string |      | 文章内容  |
| recommendStatus | bool   |      | 推荐状态  |
| viewStatus      | bool   |      | 可见状态  |
| commentStatus   | bool   |      | 评论状态  |
| password        | string |      | 访问密码  |
| articleCover    | string |      | 封面图URL |
| sortId          | uint   |      | 分类ID    |
| labelId         | uint   |      | 标签ID    |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 50. 修改文章状态

**POST** `/api/article/changeStatus`

### 请求参数

| 参数            | 类型 | 必填 | 说明     |
| --------------- | ---- | ---- | -------- |
| articleId       | uint | ✅   | 文章ID   |
| viewStatus      | bool |      | 可见状态 |
| recommendStatus | bool |      | 推荐状态 |
| commentStatus   | bool |      | 评论状态 |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 51. 删除文章图片

**POST** `/api/delArticleImage`

### 请求参数

| 参数   | 类型   | 必填 | 说明    |
| ------ | ------ | ---- | ------- |
| url    | string |      | 图片URL |
| id     | uint   |      | 资源ID  |
| userId | uint   | ✅   | 用户ID  |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 52. 修改资源状态

**POST** `/api/resource/changeResourceStatus`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 资源ID |
| flag | bool |      | 状态值 |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 53. 更新资源路径

**POST** `/api/webInfo/updateResourcePath`

### 请求参数

| 参数         | 类型   | 必填 | 说明    |
| ------------ | ------ | ---- | ------- |
| id           | uint   | ✅   | 记录ID  |
| title        | string |      | 标题    |
| classify     | string |      | 分类    |
| cover        | string |      | 封面URL |
| url          | string |      | 链接    |
| introduction | string |      | 简介    |
| type         | string |      | 类型    |
| remark       | string |      | 备注    |
| status       | bool   |      | 状态    |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 54. 删除资源路径

**POST** `/api/webInfo/deleteResourcePath`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 记录ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 55. 删除树洞

**POST** `/api/webInfo/deleteTreeHole`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 记录ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 56. 添加禁用词

**POST** `/api/prohibitedWords/add`

### 请求参数

| 参数     | 类型   | 必填 | 说明       |
| -------- | ------ | ---- | ---------- |
| avatar   | string |      | 头像URL    |
| username | string |      | 用户名     |
| message  | string | ✅   | 禁用词内容 |

### 响应体

```json
{
  "msg": "保存成功"
}
```

---

## 57. 删除禁用词

**POST** `/api/prohibitedWords/delete`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 记录ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 58. 更新禁用词

**POST** `/api/prohibitedWords/update`

### 请求参数

| 参数    | 类型   | 必填 | 说明       |
| ------- | ------ | ---- | ---------- |
| id      | uint   | ✅   | 记录ID     |
| message | string | ✅   | 禁用词内容 |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 59. 评论列表（管理）

**POST** `/api/admin/comment/boss/list`

### 请求参数

| 参数        | 类型   | 必填 | 说明     |
| ----------- | ------ | ---- | -------- |
| current     | int    |      | 当前页码 |
| size        | int    |      | 每页条数 |
| source      | int    |      | 来源ID   |
| commentType | string |      | 评论类型 |

### 响应体

```json
{
  "data": {
    "list": [
      {
        "id": 1,
        "source": 1,
        "type": "article",
        "parentCommentId": null,
        "username": "user1",
        "avatar": "https://...",
        "floorCommentId": null,
        "parentUserId": null,
        "likeCount": 5,
        "commentContent": "评论内容",
        "commentInfo": null,
        "createTime": "2024-01-01T00:00:00Z"
      }
    ],
    "totalCount": 100,
    "page": 1,
    "pageSize": 10
  }
}
```

---

## 60. 删除评论

**POST** `/api/admin/comment/boss/deleteComment`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 评论ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 61. 删除表白墙

**POST** `/api/family/deleteFamily`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 记录ID |

### 响应体

```json
{
  "msg": "删除成功"
}
```

---

## 62. 修改表白墙状态

**POST** `/api/family/changeLoveStatus`

### 请求参数

| 参数 | 类型 | 必填 | 说明   |
| ---- | ---- | ---- | ------ |
| id   | uint | ✅   | 记录ID |
| flag | bool |      | 状态值 |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

## 63. 修改微言公开状态

**POST** `/api/weiYan/changePublic`

### 请求参数

| 参数     | 类型 | 必填 | 说明     |
| -------- | ---- | ---- | -------- |
| id       | uint | ✅   | 微言ID   |
| isPublic | bool |      | 是否公开 |

### 响应体

```json
{
  "msg": "更新成功"
}
```

---

# 四、健康检查接口

> 以下接口使用 `GET` 方法，不经过 `/api` 前缀

## 64. 健康检查

**GET** `/health`

### 响应体

```json
{
  "status": "healthy"
}
```

---

## 65. 就绪检查

**GET** `/ready`

### 响应体

```json
{
  "status": "ready"
}
```

---

## 66. 系统信息

**GET** `/system`

> 仅开发环境可用

### 响应体

返回系统运行时信息

---

# 五、用户类型说明

| 类型值 | 说明               | 权限         |
| ------ | ------------------ | ------------ |
| 0      | Boss（超级管理员） | 所有权限     |
| 1      | 管理员             | 管理权限     |
| 2      | 普通用户           | 登录用户权限 |
| 3      | 访客               | 有限管理权限 |

---

# 六、错误码说明

| code | 说明                      |
| ---- | ------------------------- |
| 0    | 成功                      |
| 401  | 未授权（Token无效或过期） |
| 403  | 禁止访问（权限不足）      |
| 500  | 服务器错误                |

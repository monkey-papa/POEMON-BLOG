import requests
from django.http import JsonResponse
from rest_framework.response import Response
from rest_framework.views import APIView
from luntan.settings import summary_url, summary_authorization
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from appone.models.article import Article

class GetArticleSummaryView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        data = request.data
        # 从请求中获取文章内容
        article_content = data.get('message', '')
        article_id = data.get('articleId', 0)

        if not article_content:
            return JsonResponse({'error': '文章内容不能为空'}, status=400)
        
        try:
            # 从数据库中查找文章
            article = Article.objects.get(id=article_id)
            if article.summary:
                # 如果有摘要，直接返回数据库中的摘要
                return JsonResponse({'summary': article.summary})
        except Article.DoesNotExist:
            # 如果文章不存在，继续处理
            pass

        # AI 服务的 API URL
        api_url = summary_url
        
        # AI 服务的 API 密钥
        api_key = summary_authorization
        
        # 请求头
        headers = {
            "Authorization": f"{api_key}"
        }
        
        # 请求数据
        data = {
            "message": article_content,
            "re_chat": False,
            "stream": False,
        }
        
        try:
            # 发起请求
            response = requests.post(api_url, headers=headers, data=data)
            # 处理响应
            response.raise_for_status()  # 如果响应状态码不是 200，抛出 HTTPError 异常
            summary = response.json().get("data", {}).get("content", "")

            # 更新数据库中的摘要
            if article_id:
                Article.objects.filter(id=article_id).update(summary=summary)

            return JsonResponse({'summary': summary})
        except requests.exceptions.RequestException as e:
            # 捕获所有请求相关的异常
            return JsonResponse({'error': f'请求失败：{e}'}, status=500)
        except ValueError as e:
            # 捕获 JSON 解码错误
            return JsonResponse({'error': f'解析响应失败：{e}'}, status=500)
        except Exception as e:
            # 捕获所有其他异常
            return JsonResponse({'error': f'发生错误：{e}'}, status=500)
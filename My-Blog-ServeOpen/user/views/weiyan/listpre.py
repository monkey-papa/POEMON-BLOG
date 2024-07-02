import time

from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article
from appone.models.wei_yan import WeiYan


class WeiYanPreView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        # current 当前页码
        # size    一页显示多少内容
        # 三个搜索条件
        data = request.data
        current = data['current']
        size = data['size']
        source = data['source']
        dataall = []
        data = []
        # 通过Q来实现不同的搜索操作
        query = Q()
        query &= Q(is_public=True)
        try:
            if source and Article.objects.filter(id=source).exists():
                query &= Q(source=source)
            wei = WeiYan.objects.filter(query)
            total = wei.count()
            weis = wei.order_by('-id')[(current - 1) * size:current * size]

            for w in weis:
                data.append({
                    'id': w.id,
                    'userId': w.user_id,
                    'likeCount': w.like_count,
                    'content': w.content,
                    'type': w.type,
                    'source': w.source,
                    'isPublic': w.is_public,
                    'createTime': w.create_time,
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'total': total,
                'records': data,
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

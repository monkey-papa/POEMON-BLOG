import time

from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.client import Client
from appone.models.comment import Comment


class ListCommentView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        # current 当前页码
        # size    一页显示多少内容
        # 三个搜索条件
        data = request.data
        # print(data)
        current = int(data.get('current', 0))
        size = int(data.get('size', 0))
        source = data.get('source', 0)
        type = data.get('commentType', '').strip()
        dataall = []
        data = []
        # 通过Q来实现不同的搜索操作
        query = Q()
        try:
            if source:
                query &= Q(source=source)
            if type is not '':
                query &= Q(type=type)
            com = Comment.objects.filter(query)
            total = com.count()
            comments = com.order_by('id')[(current - 1) * size:current * size]
            for comment in comments:
                data.append({
                    'id': comment.id,
                    'source': comment.source,
                    'type': comment.type,
                    'parentCommentId': comment.parent_comment_id,
                    'username': Client.objects.get(user_id=comment.user_id).username,
                    'avatar': Client.objects.get(user_id=comment.user_id).avatar,
                    'floorCommentId': comment.floor_comment_id,
                    'parentUserId': comment.parent_user_id,
                    'likeCount': comment.like_count,
                    'commentContent': comment.comment_content,
                    'commentInfo': comment.comment_info,
                    'createTime': comment.create_time,
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'total': total,
                'data': data,
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

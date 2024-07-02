import time

from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.client import Client
from appone.models.comment import Comment


class ListClientComView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            current = int(data.get('current', 0))
            size = int(data.get('size', 0))
            csize = data.get('csize', 0)
            cs = 1
            source = data.get('source', 0)
            type = data.get('commentType', '')
            floorCommentId = data.get('floorCommentId', 0)
            dataall = []
            data = []
            ptotal = 0
            # 通过Q来实现不同的搜索操作
            query = Q()
            if source:
                query &= Q(source=source)
            if floorCommentId:
                query &= Q(floor_comment_id=floorCommentId)
            if type is not '':
                query &= Q(type=type)
            if floorCommentId:
                comm = Comment.objects.filter(query)
                comments = comm.order_by('id')[(current - 1) * size:current * size]
            else:
                comm = Comment.objects.filter(query)
                query &= Q(floor_comment_id=0)
                com = Comment.objects.filter(query)
                ptotal = com.count()
                comments = com.order_by('id')[(current - 1) * size:current * size]
            total = comm.count()
            for comment in comments:
                # print(comments)
                childComments = []
                clientdata = []
                if comment.floor_comment_id == 0:
                    childs = Comment.objects.filter(floor_comment_id=comment.id)
                    if childs.exists():
                        for child in childs:
                            if csize:
                                if cs <= csize:
                                    if comment.parent_comment_id is None:
                                        clientdata.append({
                                            'id': child.id,
                                            'source': child.source,
                                            'type': child.type,
                                            'parentCommentId': child.parent_comment_id,
                                            'userId': child.user_id,
                                            'username': Client.objects.get(user_id=child.user_id).username,
                                            'avatar': Client.objects.get(user_id=child.user_id).avatar,
                                            'userType': Client.objects.get(user_id=child.user_id).user_type,
                                            'floorCommentId': child.floor_comment_id,
                                            'parentUserId': child.parent_user_id,
                                            'likeCount': child.like_count,
                                            'commentContent': child.comment_content,
                                            'commentInfo': child.comment_info,
                                            'createTime': child.create_time,
                                            'parentUsername': Client.objects.get(
                                                user_id=Comment.objects.get(id=child.parent_comment_id).user_id).username,
                                        })
                                    else:
                                        clientdata.append({
                                            'id': child.id,
                                            'source': child.source,
                                            'type': child.type,
                                            'parentCommentId': child.parent_comment_id,
                                            'userId': child.user_id,
                                            'username': Client.objects.get(user_id=child.user_id).username,
                                            'avatar': Client.objects.get(user_id=child.user_id).avatar,
                                            'userType': Client.objects.get(user_id=child.user_id).user_type,
                                            'floorCommentId': child.floor_comment_id,
                                            'parentUserId': child.parent_user_id,
                                            'likeCount': child.like_count,
                                            'commentContent': child.comment_content,
                                            'commentInfo': child.comment_info,
                                            'createTime': child.create_time,
                                            'parentUsername': Client.objects.get(user_id=child.user_id).username,
                                        })
                                    cs = cs + 1
                            else:
                                for child in childs:
                                    if csize:
                                        if cs <= csize:
                                            if comment.parent_comment_id is None:
                                                clientdata.append({
                                                    'id': child.id,
                                                    'source': child.source,
                                                    'type': child.type,
                                                    'parentCommentId': child.parent_comment_id,
                                                    'userId': child.user_id,
                                                    'username': Client.objects.get(user_id=child.user_id).username,
                                                    'avatar': Client.objects.get(user_id=child.user_id).avatar,
                                                    'userType': Client.objects.get(user_id=child.user_id).user_type,
                                                    'floorCommentId': child.floor_comment_id,
                                                    'parentUserId': child.parent_user_id,
                                                    'likeCount': child.like_count,
                                                    'commentContent': child.comment_content,
                                                    'commentInfo': child.comment_info,
                                                    'createTime': child.create_time,
                                                    'parentUsername': Client.objects.get(
                                                        user_id=Comment.objects.get(
                                                            id=child.parent_comment_id).user_id).username,
                                                })
                                            else:
                                                clientdata.append({
                                                    'id': child.id,
                                                    'source': child.source,
                                                    'type': child.type,
                                                    'parentCommentId': child.parent_comment_id,
                                                    'userId': child.user_id,
                                                    'username': Client.objects.get(user_id=child.user_id).username,
                                                    'avatar': Client.objects.get(user_id=child.user_id).avatar,
                                                    'userType': Client.objects.get(user_id=child.user_id).user_type,
                                                    'floorCommentId': child.floor_comment_id,
                                                    'parentUserId': child.parent_user_id,
                                                    'likeCount': child.like_count,
                                                    'commentContent': child.comment_content,
                                                    'commentInfo': child.comment_info,
                                                    'createTime': child.create_time,
                                                    'parentUsername': Client.objects.get(
                                                        user_id=child.user_id).username,
                                                })
                        cs = 1
                        if floorCommentId is 0:
                            dataall.append({
                                'code': 200,
                                'message': "null",
                                'total': childs.count(),
                                'data': clientdata,
                                'currentTimeMillis': time.time(),
                            })
                            return Response({
                                'result': dataall
                            })
                    childComments.append({
                        # 'countId':,
                        'current': 1,
                        # 'hitCount':,
                        # 'maxLimit':,
                        # 'optimizeCountSql':,
                        # 'orders':,
                        # 'pages':,
                        'records': clientdata,
                        # 'searchCount':,
                        # 'size':,
                        'total': childs.count(),
                    })
                if floorCommentId is None:
                    data.append({
                        'id': comment.id,
                        'source': comment.source,
                        'type': comment.type,
                        'parentCommentId': comment.parent_comment_id,
                        'userId': comment.user_id,
                        'username': Client.objects.get(user_id=comment.user_id).username,
                        'avatar': Client.objects.get(user_id=comment.user_id).avatar,
                        'userType': Client.objects.get(user_id=comment.user_id).user_type,
                        'floorCommentId': comment.floor_comment_id,
                        'parentUserId': comment.parent_user_id,
                        'likeCount': comment.like_count,
                        'commentContent': comment.comment_content,
                        'commentInfo': comment.comment_info,
                        'createTime': comment.create_time,
                        'childComments': childComments,
                    })
                else:
                    data.append({
                        'id': comment.id,
                        'source': comment.source,
                        'type': comment.type,
                        'parentCommentId': comment.parent_comment_id,
                        'userId': comment.user_id,
                        'username': Client.objects.get(user_id=comment.user_id).username,
                        'avatar': Client.objects.get(user_id=comment.user_id).avatar,
                        'userType': Client.objects.get(user_id=comment.user_id).user_type,
                        'floorCommentId': comment.floor_comment_id,
                        'parentUserId': comment.parent_user_id,
                        'likeCount': comment.like_count,
                        'commentContent': comment.comment_content,
                        'commentInfo': comment.comment_info,
                        'createTime': comment.create_time,
                        'childComments': childComments,
                        'parentUsername': Client.objects.get(
                            user_id=Comment.objects.get(id=comment.parent_comment_id).user_id).username,
                    })
            if floorCommentId:
                total1 = total + 1
            else:
                total1 = total
            dataall.append({
                'code': 200,
                'message': "null",
                'total': total1,
                'parenttotal': total - ptotal,
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

import time

from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.comment import Comment


class AddCommentView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            source = data['source']
            type = data['type']
            parentCommentId = data['floorCommentId']
            userId = data['userId']
            floorCommentId = data['parentCommentId']
            parentUserId = data['parentUserId']
            likeCount = data['likeCount']
            commentContent = data['commentContent']
            commentInfo = data['commentInfo']
            if floorCommentId and parentCommentId:
                Comment.objects.create(source=source, type=type, parent_comment_id=floorCommentId,
                                       user_id=userId, floor_comment_id=parentCommentId,
                                       parent_user_id=parentUserId, like_count=likeCount,
                                       comment_content=commentContent, comment_info=commentInfo)
            else:
                Comment.objects.create(source=source, type=type, parent_comment_id=parentCommentId,
                                       user_id=userId, floor_comment_id=floorCommentId,
                                       parent_user_id=parentUserId, like_count=likeCount,
                                       comment_content=commentContent, comment_info=commentInfo)

            return Response({
                'result': "success"
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

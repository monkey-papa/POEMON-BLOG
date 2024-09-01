# Create your views here.
import time

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article
from appone.models.label import Label
from appone.models.sort import Sort
from appone.models.tree_hole import TreeHole


class SaveTreeView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            avatar = data['avatar']
            username = data['username']
            message = data['message']
            if avatar and message:
                tr = TreeHole.objects.create(avatar=avatar, message=message, username=username)
                tree = TreeHole.objects.get(id=tr.id)
                dataa = []
                dataall = []
                dataa.append({
                    'id': tree.id,
                    'username': tree.username,
                    'createTime': tree.create_time,
                    'message': tree.message,
                    'avatar': tree.avatar,
                    })

                dataall.append({
                    'code': 200,
                    'message': "null",
                    'records': dataa,
                    'currentTimeMillis': time.time(),
                })
                return Response({
                    'result': dataall
                })
            else:
                return Response({
                    'failure': "exists null",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

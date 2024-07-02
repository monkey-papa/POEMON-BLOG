# Create your views here.
import time

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.tree_hole import TreeHole


class ListTreeView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            current = int(data.get('current', 0))
            size = int(data.get('size', 0))
            dataall = []
            data = []
            res = TreeHole.objects.all()
            total = res.count()
            if size == 0:
                ress = res
            else:
                ress = res[(current - 1) * size:current * size]
            for res in ress:
                data.append({
                    'id': res.id,
                    'username': res.username,
                    'createTime': res.create_time,
                    'message': res.message,
                    'avatar': res.avatar,
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

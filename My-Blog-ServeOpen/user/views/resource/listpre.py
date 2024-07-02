# Create your views here.
import time

from django.db.models import Q
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.resource import Resource


class ListResourcePreView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            current = int(data.get('current', 0))
            size = int(data.get('size', 0))
            resourceType = data.get('resourceType', '')
            dataall = []
            data = []
            # 通过Q来实现不同的搜索操作
            query = Q()
            query &= Q(status=True)
            if resourceType:
                query &= Q(mime_type=resourceType)
            res = Resource.objects.filter(query)
            total = res.count()
            ress = res[(current - 1) * size:current * size]
            for res in ress:
                data.append({
                    'id': res.id,
                    'createTime': res.create_time,
                    'mimeType': res.mime_type,
                    'path': res.path,
                    'size': res.size,
                    'status': res.status,
                    'type': res.type,
                    'userId': res.user_id_id
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

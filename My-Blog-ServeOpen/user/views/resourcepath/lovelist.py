# Create your views here.
import time

from django.db.models import Q, Count
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.resource import Resource
from appone.models.resource_path import ResourcePath


class ListLoveResourcePathView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            dataall = []
            data = []
            data1 = request.GET
            type = data1.get('type', '').strip()
            # 分类，以类别分类，获取每个类别的数量
            categories = ResourcePath.objects.filter(status=True, type=type).values('classify').annotate(num_classifys=Count('id'))
            for res in categories:
                data.append({
                    'classify': res['classify'],
                    'count': res['num_classifys']
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'total': len(data),
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

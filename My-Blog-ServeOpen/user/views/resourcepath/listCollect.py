# Create your views here.
import time

from django.db.models import Q
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.resource import Resource
from appone.models.resource_path import ResourcePath


class ListCollectView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            dataall = []
            data = {}

            # 通过Q来实现不同的搜索操作
            query = Q()
            query &= Q(status=True)
            query &= Q(type='favorites')
            rs = ResourcePath.objects.filter(query).values('classify').distinct()
            total = 0
            for r in rs:
                dat = []
                r_name = r['classify']
                ress = ResourcePath.objects.filter(classify=r_name, status=True)
                total += ress.count()
                for res in ress:
                    dat.append({
                        'id': res.id,
                        'classify': res.classify,
                        'createTime': res.create_time,
                        'introduction': res.introduction,
                        'cover': res.cover,
                        'remark': res.remark,
                        'status': res.status,
                        'title': res.title,
                        'type': res.type,
                        'url': res.url
                    })
                data['{}'.format(r_name)] = dat

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

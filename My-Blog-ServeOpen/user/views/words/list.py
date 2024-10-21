# Create your views here.
import time

from django.db.models import Q
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.words import Words


class ListWordsView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        data = request.data
        current = int(data.get('current', 0))
        size = int(data.get('size', 0))
        searchKey = data.get('searchKey', '').strip()
        dataall = []
        data = []
        # 通过Q来实现不同的搜索操作
        query = Q()
        try:
            if searchKey:
                query &= Q(message__contains=searchKey)
            word = Words.objects.filter(query)
            total = word.count()
            words = word.order_by('id')[(current-1)*size:current*size]
            for wor in words:
                data.append({
                    'id': wor.id,
                    'username': wor.username,
                    'message': wor.message,
                    'avatar': wor.avatar,
                    'createTime': wor.create_time,
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

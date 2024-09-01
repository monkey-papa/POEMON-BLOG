import time

from django.db.models import Q
from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.client import Client


class frontClientView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        # current 当前页码
        # size    一页显示多少内容
        # 三个搜索条件
        data = request.data
        current = int(data.get('current', 0))
        size = int(data.get('size', 0))
        searchKey = data.get('searchKey', '').strip()
        userStatus = data.get('userStatus', '').strip()
        if userStatus == 'true':
            userStatus = 'True'
        if userStatus == 'false':
            userStatus = 'False'
        userType = data.get('userType', '').strip()
        dataall = []
        data = []

        # 通过Q来实现不同的搜索操作
        query = Q()
        try:
            if searchKey:
                query &= Q(phone_number__contains=searchKey)
            if userStatus:
                query &= Q(user_status=userStatus)
            if userType is not '':
                query &= Q(user_type=int(userType))
            cent = Client.objects.filter(query)
            total = cent.count()
            clients = cent.order_by('user_id')[(current-1)*size:current*size]
            for client in clients:
                data.append({
                    'id': client.user_id,
                    'username': client.username,
                    'avatar': client.avatar,
                    'gender': client.gender,
                    'email': client.email
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







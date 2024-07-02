from django.shortcuts import render
from datetime import timedelta, datetime

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authtoken.models import Token
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from django.contrib.auth.models import User
# Create your views here.
from appone.models.client import Client
from appone.models.code import Code


class RegretView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        data = request.data
        username = data['username']
        password = data['password']
        email = data['email']
        code = data['code']

        if not (username and password):
            return Response({'error': 'Please provide both username and password.'}, status=400)

        user_exists = User.objects.filter(username=username).exists()

        if user_exists:
            pass
        else:
            return Response({'error': 'Username does not exist.'}, status=400)

        if not email or not code:
            return Response({
                'result': "邮箱或验证码不能为空"
            })
        code_log = Code.objects.filter(email=email).order_by('-id')
        if code_log:
            newly_log = code_log[0]
            time = datetime.now() - timedelta(hours=0, minutes=5, seconds=0)
            if newly_log.create_time.replace(tzinfo=None) >= time:
                if newly_log.code == code:
                    u = User.objects.get(username=username)
                    u.set_password(password)
                    u.save()
                    return Response({'result': "success"})
                else:
                    return Response({
                        'result': "验证码错误"
                    })
            else:
                return Response({
                    'result': "验证码已过期"
                })
        else:
            return Response({
                'result': "error"
            })









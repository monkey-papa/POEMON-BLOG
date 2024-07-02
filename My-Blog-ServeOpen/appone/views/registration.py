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
from appone.models.resource import Resource


class RegisterView(APIView):
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
            return Response({'error': 'Username already exists.'}, status=400)
        if not email or not code:
            return Response({
                'result': "邮箱或验证码不能为空"
            })
        code_log = Code.objects.filter(email=email).order_by('-create_time')
        if code_log:
            newly_log = code_log[0]
            time = datetime.now() - timedelta(hours=0, minutes=5, seconds=0)
            if newly_log.create_time.replace(tzinfo=None) >= time:
                if newly_log.code == code:
                    user = User.objects.create_user(username=username, password=password)
                    Client.objects.create(user=user, username=username, email=email)
                    token, created = Token.objects.get_or_create(user=user)
                    user = User.objects.filter(username=username).first()
                    client = Client.objects.get(user_id=user.id)
                    data = []
                    dataall = []
                    avatar_a = ''
                    r = Resource.objects.filter(path=client.avatar)
                    if r.exists():
                        if r[0].status:
                            avatar_a = client.avatar
                    else:
                        avatar_a = client.avatar

                    data.append({
                        'accessToken': token.key,
                        'id': client.user_id,
                        'username': client.username,
                        'phoneNumber': client.phone_number,
                        'email': client.email,
                        'admire': client.admire,
                        'userStatus': client.user_status,
                        'avatar': avatar_a,
                        'gender': client.gender,
                        'introduction': client.introduction,
                        'userType': client.user_type,
                        'createTime': client.create_time,
                    })

                    dataall.append({
                        'code': 200,
                        'message': "null",
                        'data': data,
                        'currentTimeMillis': time.time(),
                    })
                    return Response({
                        'result': dataall
                    })
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









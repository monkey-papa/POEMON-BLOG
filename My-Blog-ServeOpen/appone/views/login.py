import time
import re
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authtoken.models import Token
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from django.contrib.auth.models import User
from appone.models.client import Client
from appone.models.resource import Resource


class LoginView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        username = request.data.get("account", "")
        password = request.data.get("password")
        p = request.data.get("province")
        str = r'^[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+){0,4}@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+){0,4}$'
        if re.match(str, username):
            client = Client.objects.get(email=username)
            if client is not None:
                user = User.objects.filter(username=client.username).first()
            else:
                user = User.objects.filter(username=username).first()
                client = Client.objects.get(user_id=user.id)
        else:
            user = User.objects.filter(username=username).first()
            client = Client.objects.get(user_id=user.id)
        if client.user_status is False or client.deleted is True:
            return Response({"result": "The account is disabled or deleted"})
        if user and user.check_password(password):
            token, created = Token.objects.get_or_create(user=user)
            data = []
            dataall = []
            avatar_a = ""
            r = Resource.objects.filter(path=client.avatar)
            if r.exists():
                if r[0].status:
                    avatar_a = client.avatar
            else:
                avatar_a = client.avatar
            # 找到，没有的话就province
            client = Client.objects.get(user_id=user.id)
            if client.province is None:
                client.province = p
                client.save()
            data.append(
                {
                    "accessToken": token.key,
                    "id": client.user_id,
                    "username": client.username,
                    "phoneNumber": client.phone_number,
                    "email": client.email,
                    "admire": client.admire,
                    "userStatus": client.user_status,
                    "avatar": avatar_a,
                    "gender": client.gender,
                    "introduction": client.introduction,
                    "userType": client.user_type,
                    "createTime": client.create_time,
                    'qiniuDomain': client.qiniu_domain,
                    'qiniuBucketName': client.qiniu_bucket_name,
                    'qiniuSecretKey': client.qiniu_secret_key,
                    'qiniuAccessKey': client.qiniu_access_key,
                }
            )
            dataall.append(
                {
                    "code": 200,
                    "message": "null",
                    "data": data,
                    "currentTimeMillis": time.time(),
                }
            )
            return Response({"result": dataall})
        else:
            return Response({"error": "用户名或密码错误"}, status=400)

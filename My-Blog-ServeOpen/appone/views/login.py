import time
import re

from django.shortcuts import render

# Create your views here.
from datetime import timedelta, datetime

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authtoken.models import Token
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from django.contrib.auth.models import User

from appone.models.client import Client
from appone.models.ip import Ip
from appone.models.resource import Resource


class LoginView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        username = request.data.get("account", "")
        password = request.data.get("password")
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
            return Response({"error": "Username or password is incorrect."}, status=400)
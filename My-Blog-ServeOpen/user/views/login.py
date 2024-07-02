from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authtoken.models import Token
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from django.contrib.auth.models import User

from appone.models.client import Client


class LoginView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        username = request.data.get('account')
        password = request.data.get('password')

        user = User.objects.filter(username=username).first()

        client = Client.objects.get(user_id=user.id)
        type = client.user_type

        if user and user.check_password(password) and type != 2 and client.user_status:
            token, created = Token.objects.get_or_create(user=user)
            return Response({"data": {'accessToken': token.key,
                                      'avatar': client.avatar,
                                      'username': username, 'id': User.objects.get(username=username).id}})
        else:
            return Response({'error': 'Username or password is incorrect or non-administrators.'}, status=400)

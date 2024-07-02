import json
import time

from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny


from appone.models.web_info import WebInfo


class UserView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        with open('static/user.json', 'r', encoding='utf-8') as file:
            data = json.load(file)
        return Response(data)

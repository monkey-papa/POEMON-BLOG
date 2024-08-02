import json
import time

from django.db.models import Count
from django.shortcuts import render

# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.client import Client
from appone.models.ip import Ip
from appone.models.web_info import WebInfo


class UserView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        with open("static/user.json", "r", encoding="utf-8") as file:
            data = json.load(file)
        queryset = Client.objects.filter(province__isnull=False).values('province').annotate(
            b=Count('*')).order_by('province')
        data1 = [{'name': item['province'], 'value': str(item['b'])} for item in queryset]
        for item in data1:
            p = item['name']
            vp = str(item['value'])
            for item1 in data:
                if p in item1["name"] and p != '' and p:
                    item1["value"] = vp
        return Response(data)

import time

from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser


from appone.models.web_info import WebInfo


class GWebInfoView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        dataall = []
        data = []
        try:
            webinfos = WebInfo.objects.all().order_by('-id')
            for webinfo in webinfos:
                data.append({
                    'id': webinfo.id,
                    'webName': webinfo.web_name,
                    'webTitle': webinfo.web_title,
                    'notices': webinfo.notices,
                    'footer': webinfo.footer,
                    'backgroundImage': webinfo.background_image,
                    'avatar': webinfo.avatar,
                    'randomAvatar': webinfo.random_avatar,
                    'randomName': webinfo.random_name,
                    'randomCover': webinfo.random_cover,
                    'waifuJson': webinfo.waifu_json,
                    'status': webinfo.status,
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'data': data,
                'currentTimeMillis': time.time(),
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

        return Response({'result': dataall})
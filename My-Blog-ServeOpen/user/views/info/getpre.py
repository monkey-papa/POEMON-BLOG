import time

from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.resource import Resource
from appone.models.web_info import WebInfo


class GWebInfoPreView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        dataall = []
        data = []
        try:
            webinfos = WebInfo.objects.filter(status=True).order_by('-id')
            for webinfo in webinfos:
                avatar_a = ''
                backgroundImage_a = ''
                randomAvatar_a = ''
                randomCover_a = ''
                avatar_a_r = Resource.objects.filter(path=webinfo.avatar)
                backgroundImage_a_r = Resource.objects.filter(path=webinfo.background_image)
                randomAvatar_a_r = Resource.objects.filter(path=webinfo.random_avatar)
                randomCover_a_r = Resource.objects.filter(path=webinfo.random_cover)
                if avatar_a_r.exists():
                    if avatar_a_r[0].status:
                        avatar_a = webinfo.avatar
                # else:
                #     avatar_a = webinfo.avatar

                if backgroundImage_a_r.exists():
                    if backgroundImage_a_r[0].status:
                        backgroundImage_a = webinfo.background_image
                # else:
                #     backgroundImage_a = webinfo.background_image

                if randomAvatar_a_r.exists():
                    if randomAvatar_a_r[0].status:
                        randomAvatar_a = webinfo.random_avatar
                # else:
                #     randomAvatar_a = webinfo.random_avatar

                if randomCover_a_r.exists():
                    if randomCover_a_r[0].status:
                        randomCover_a = webinfo.random_cover
                else:
                    randomCover_a = webinfo.random_cover

                data.append({
                    'id': webinfo.id,
                    'webName': webinfo.web_name,
                    'webTitle': webinfo.web_title,
                    'notices': webinfo.notices,
                    'footer': webinfo.footer,
                    'backgroundImage': backgroundImage_a,
                    'avatar': avatar_a,
                    'randomAvatar': randomAvatar_a,
                    'randomName': webinfo.random_name,
                    'randomCover': randomCover_a,
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
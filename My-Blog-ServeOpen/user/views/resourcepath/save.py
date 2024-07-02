# Create your views here.
from django.contrib.auth.models import User
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.label import Label
from appone.models.resource import Resource
from appone.models.resource_path import ResourcePath


class SaveResourcePathView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            title = data.get('title', '')
            classify = data.get('classify', '')
            introduction = data.get('introduction', '')
            cover = data.get('cover', '')
            remark = data.get('friendAvatar', '')
            url = data.get('url', '')
            type = data.get('type', '')
            # remark = data.get('remark', '')

            if title:
                ResourcePath.objects.create(title=title, classify=classify, introduction=introduction,
                                            cover=cover, url=url, type=type, remark=remark, status=True)
                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

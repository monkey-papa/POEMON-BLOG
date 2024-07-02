# Create your views here.

from django.db.models import Q
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.label import Label
from appone.models.resource_path import ResourcePath


class UpdateResourcePathView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            title = data.get('title', '')
            classify = data.get('classify', '')
            cover = data.get('cover', '')
            url = data.get('url', '')
            introduction = data.get('introduction', '')
            type = data.get('type', '')
            status = data.get('status', None)
            remark = data.get('remark', '')
            ress = ResourcePath.objects.filter(id=id)
            if ress.exists():
                if title:
                    ress.update(title=title)
                if classify:
                    ress.update(classify=classify)
                if cover:
                    ress.update(cover=cover)
                if url:
                    ress.update(url=url)
                if introduction:
                    ress.update(introduction=introduction)
                if type:
                    ress.update(type=type)
                if remark:
                    ress.update(remark=remark)
                if status is not None:
                    ress.update(status=status)
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

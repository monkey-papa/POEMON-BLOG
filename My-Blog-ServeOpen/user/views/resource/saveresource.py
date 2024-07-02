# Create your views here.
from django.contrib.auth.models import User
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.label import Label
from appone.models.resource import Resource


class SaveResourceView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            type = data['type']
            path = data['path']
            size = data['size']
            mime_type = data['mimeType']
            user = User.objects.filter(id=id).first()
            if user and type and path and size and mime_type:
                Resource.objects.create(type=type, path=path,
                                        size=size, mime_type=mime_type, user_id_id=id)
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

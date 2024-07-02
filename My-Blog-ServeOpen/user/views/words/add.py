# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.words import Words


class SaveWordsView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            message = data['message']
            username = data['username']
            avatar = data['avatar']
            if message and Words.objects.filter(message=message).exists() is False:
                Words.objects.create(message=message, username=username, avatar=avatar)
                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null or error",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

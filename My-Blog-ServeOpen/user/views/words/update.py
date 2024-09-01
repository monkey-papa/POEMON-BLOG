# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.words import Words


class UpdateWordsView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            message = data['message']
            if id:
                word = Words.objects.filter(id=id)
                if word.exists() and message and Words.objects.filter(message=message).exists() is False:
                    word.update(message=message)
                    return Response({
                        'result': "success",
                    })
                else:
                    return Response({
                        'failure': "exists or null",
                    })
            else:
                return Response({
                    'failure': "exists null or error",
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

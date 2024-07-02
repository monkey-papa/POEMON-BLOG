# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.words import Words


class DeleteWordsView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = data['id']
            if id:
                Words.objects.get(id=id).delete()
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

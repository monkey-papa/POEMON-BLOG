# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article

class DeleteArtView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            if id:
                Article.objects.get(id=id).delete()
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

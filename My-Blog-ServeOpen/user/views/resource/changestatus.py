# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.resource import Resource


class ChangeResourceView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            status = data.get('flag', "")
            if status == 'true':
                status = 'True'
            if status == 'false':
                status = 'False'
            if id:
                res = Resource.objects.filter(id=id)
                res.update(status=status)
                return Response({
                    'result': "sucess"
                })
            else:
                return Response({
                    'result': "failure and exist null"
                })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

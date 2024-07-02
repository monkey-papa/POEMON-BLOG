# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny


from appone.models.client import Client
from appone.models.family import Family


class ChangeFamilyStatusView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            status = data.get('flag', "").strip()
            if status == 'true':
                status = 'True'
            if status == 'false':
                status = 'False'
            Family.objects.filter(id=id).update(status=status)
            return Response({
                'result': "success",
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.client import Client


class ChangeUserStatusView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.POST
            id = int(data.get('userId', 0))
            status = data.get('flag', "").strip()
            if status == 'true':
                status = 'True'
            if status == 'false':
                status = 'False'
            Client.objects.filter(user_id=id).update(user_status=status)
            return Response({
                'result': "success",
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
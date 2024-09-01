# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser


from appone.models.client import Client


class ChangeUserAdmireView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.POST
            id = int(data.get('userId', 0))
            admire = data.get('admire', "").strip()
            Client.objects.filter(user_id=id).update(admire=admire)
            return Response({
                'result': "success",
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
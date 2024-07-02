# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny


from appone.models.client import Client


class ChangeUserTypeView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.POST
            id = int(data.get('userId', 0))
            type = int(data.get('userType', 2))
            Client.objects.filter(user_id=id).update(user_type=type)
            return Response({
                'result': "success",
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
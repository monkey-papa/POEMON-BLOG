# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser
from appone.models.client import Client
from django.contrib.auth.models import User

class ChangeUserTypeView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]
    # permission_classes = [IsAuthenticated]
    # authentication_classes = [TokenAuthentication]
    def post(self, request):
        try:
            data = request.POST
            id = int(data.get('userId', 0))
            type = int(data.get('userType', 2))
            if type < 2 or type == 3:
                User.objects.filter(id=id).update(is_staff=1)
            else:
                User.objects.filter(id=id).update(is_staff=0)
            Client.objects.filter(user_id=id).update(user_type=type)
            return Response({
                'result': "success",
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
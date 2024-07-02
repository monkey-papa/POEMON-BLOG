import time
from django.contrib.auth.models import User
from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.wei_yan import WeiYan


class AddWeiYanView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            userId = data['userId']
            content = data['content']
            isPublic = data['isPublic']
            type = data.get('type', 'a')

            user = User.objects.filter(id=userId)
            if user.exists():
                WeiYan.objects.create(user_id=userId, content=content, is_public=isPublic, type=type)
                return Response({
                    'result': "success"
                })
            else:
                return Response({
                    'result': "exist null"
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

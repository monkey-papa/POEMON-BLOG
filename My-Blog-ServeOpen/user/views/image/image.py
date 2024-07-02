from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from luntan.tools.qiniu import upload_to_qiniu


class ImageView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        image = request.FILES.get('image')
        print(image)
        image_url = upload_to_qiniu(image)
        # 返回图片的 URL
        return Response({'url': image_url})
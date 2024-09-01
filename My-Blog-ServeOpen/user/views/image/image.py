from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from luntan.tools.qiniu import upload_to_qiniu


class ImageView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        user_id = int(request.data.get("userId", 0))
        image = request.FILES.get('image')
        image_url = upload_to_qiniu(image, user_id)
        if image_url is None:
            return Response({'error': "上传七牛云错误"}, status=400)
        # 返回图片的 URL
        return Response({'url': image_url})

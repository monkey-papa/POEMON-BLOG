from rest_framework.response import Response
from rest_framework.views import APIView
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser
from appone.models.resource import Resource
from luntan.tools.qiniu import delete_image


class DeleteImageView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        image = request.data.get('url', '')
        id = int(request.data.get('id', 0))
        user_id = int(request.data.get('userId', 0))
        if user_id == 0:
            return Response({'result': "userId缺失"})
        if id:
            Resource.objects.get(id=id).delete()
        if delete_image(image, user_id):
            # 在这里可以根据需求处理删除成功后的操作，比如更新数据库等
            return Response({'result': "success"})
        else:
            return Response({'error': "七牛云删除失败"}, status=400)

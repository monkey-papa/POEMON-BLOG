from rest_framework.permissions import AllowAny
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication


from appone.models.code import Code
from luntan.tools.code.code import send_code


class CodeView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]
    # 验证码发送
    def post(self, request):
        data = request.data
        email = data.get("email").strip()
        try:
            code = send_code(email)
            Code.objects.create(email=email, code=code)
            return Response({'result': "success"})
        except Exception as error:
            return Response({'result': "failure {0}".format(error)})

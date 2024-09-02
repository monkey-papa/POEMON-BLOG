from rest_framework.permissions import IsAdminUser
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from luntan.tools.code.codeComment import send_code
from luntan.tools.code.codeApprove import send_code_approve


class CodeCommentView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]
    # 验证码发送
    def post(self, request):
        data = request.data
        comment = data.get("comment").strip()
        name = data.get("name").strip()
        email = data.get("email").strip()
        type = data.get("type").strip()
        try:
            if type == "approve":
                send_code_approve(email, comment, name)
            else:
                send_code(email, comment, name)
            return Response({'result': "success"})
        except Exception as error:
            return Response({'result': "failure {0}".format(error)})

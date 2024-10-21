import time
from datetime import timedelta, datetime
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView
from appone.models.ip import Ip


class GetIpView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = []
            dataall = []
            ip = request.META.get('REMOTE_ADDR')

            current_time = datetime.now()
            now_ip = Ip.objects.filter(ip=ip)
            gap_time = current_time - timedelta(days=1)
            if now_ip.exists():
                pass
            else:
                Ip.objects.create(ip=ip, create_time=current_time, dcount=1)

            data.append({
                'ip': ip,
            })
            dataall.append({
                'code': 200,
                'message': "null",
                'data': data,
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

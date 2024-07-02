# # Create your views here.
# import time
# from datetime import timedelta, datetime
#
# from rest_framework.authentication import TokenAuthentication
# from rest_framework.permissions import AllowAny
# from rest_framework.response import Response
# from rest_framework.views import APIView
#
# from appone.models.ip import Ip
# from user.views.ip.getprvoince import get_province
#
#
# class GetIpView(APIView):
#     permission_classes = [AllowAny]
#     authentication_classes = [TokenAuthentication]
#
#     def get(self, request):
#         try:
#             data = []
#             dataall = []
#             ip = request.META.get('REMOTE_ADDR')
#
#             current_time = datetime.now()
#
#             now_ip = Ip.objects.filter(ip=ip)
#             gap_time = current_time - timedelta(days=1)
#             if now_ip.exists():
#                 if now_ip.first().create_time.date() >= gap_time.date():
#                     count = Ip.objects.get(ip=ip).dcount
#                 else:
#                     ddcount = now_ip[0].dcount + 1
#                     b_ip = Ip.objects.get(ip=ip)
#                     b_ip.dcount = ddcount
#                     b_ip.k_time = b_ip.create_time
#                     b_ip.create_time = current_time
#                     b_ip.save()
#                     count = b_ip.dcount
#             else:
#                 Ip.objects.create(ip=ip, create_time=current_time, dcount=1, province=get_province(ip)[2], city=get_province(ip)[3])
#                 count = 1
#
#             data.append({
#                 'ip': ip,
#                 'count': count,
#             })
#             dataall.append({
#                 'code': 200,
#                 'message': "null",
#                 'data': data,
#                 'currentTimeMillis': time.time(),
#             })
#             return Response({
#                 'result': dataall
#             })
#         except Exception as error:
#             return Response({
#                 'result': "failure {0}".format(error)
#             })

# Create your views here.
import time
from datetime import timedelta, datetime

from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView

from appone.models.ip import Ip
from user.views.ip.getprvoince import get_province


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

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

from appone.models.client import Client
from appone.models.ip import Ip
from user.views.ip.getprvoince import get_province


class RcView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            ip = request.META.get('REMOTE_ADDR')
            region = request.data['province']
            city = request.data['city']
            userid = request.data.get('userId', '')

            current_time = datetime.now()
            last_time = current_time - timedelta(days=1)
            start_last = datetime(last_time.year, last_time.month, last_time.day, 0, 0, 0)
            end_last = datetime(last_time.year, last_time.month, last_time.day, 23, 59, 59)

            timeone = 0
            if userid:
                if Client.objects.filter(user_id=userid).exists():
                    if Ip.objects.filter(user_id=userid).exists():
                        ipone = Ip.objects.get(user_id=userid)
                        timeone = ipone.create_time
                        ipone.ip = ip
                        ipone.province = region
                        ipone.city = city
                        ipone.save()
                        if timeone <= end_last:
                            ipone1 = Ip.objects.get(user_id=userid)
                            ipone1.k_time = ipone1.create_time
                            ipone1.create_time = current_time
                            ipone1.dcount = ipone1.dcount + 1
                            ipone1.save()
                    else:
                        Ip.objects.filter(user_id=userid, ip=ip, dcount=1, province=region, city=city)
            else:
                now_ip = Ip.objects.filter(ip=ip)
                for i in now_ip:
                    if i.user_id:
                        pass
                    else:
                        i.province = region
                        i.city = city
                        i.save()
                        if i.create_time <= end_last:
                            i.k_time = i.create_time
                            i.create_time = current_time
                            i.dcount = i.dcount + 1
                            i.save()

            return Response({
                'result': 'success'
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

# Create your views here.
import time
from datetime import timedelta, datetime

from django.db.models import Count
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView

from appone.models.client import Client
from appone.models.ip import Ip
from django.db.models import Sum


class ListIpView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            provinceall = []
            ipall = []
            provincetoday = []
            usertoday = []
            yesterday = []
            dataall = []

            current_time = datetime.now()
            last_time = current_time - timedelta(days=1)
            start_last = datetime(last_time.year, last_time.month, last_time.day, 0, 0, 0)
            end_last = datetime(last_time.year, last_time.month, last_time.day, 23, 59, 59)

            total_sum = Ip.objects.aggregate(sum=Sum('dcount'))['sum']
            today_sum = Ip.objects.filter(create_time__range=(end_last, current_time)).count()
            yesterday_sum = Ip.objects.filter(k_time__range=(start_last, end_last)).count() + Ip.objects.filter(
                create_time__range=(start_last, end_last)).count()

            province_all_top = Ip.objects.values('province').annotate(f=Sum('dcount')).order_by('-f')[:10]
            # print(province_all_top)
            # province_all_top = Ip.objects.values('province').annotate(f=Sum('dcount')).order_by('f')[:10]

            ip_all_top = Ip.objects.values('ip').annotate(f=Sum('dcount')).order_by('-f')[:10]
            # print(ip_all_top)
            province_today = Ip.objects.filter(create_time__range=(end_last, current_time)).values(
                'province').annotate(
                f=Count('province')).order_by('-f')
            user_today = Ip.objects.filter(create_time__range=(end_last, current_time)).order_by('create_time')
            last_user1 = Ip.objects.filter(k_time__range=(start_last, end_last)).order_by('create_time')
            last_user2 = Ip.objects.filter(create_time__range=(start_last, end_last)).order_by('create_time')
            last_user = last_user1 | last_user2
            for item in province_all_top:
                if item['province']:
                    provinceall.append({
                        'province': item['province'],
                        'count': item['f'],
                    })

            for item in ip_all_top:
                if item['ip']:
                    ipall.append({
                        'ip': item['ip'],
                        'count': item['f'],
                    })

            for item in province_today:
                if item['province']:
                    provincetoday.append({
                        'province': item['province'],
                        'count': item['f'],
                    })
            for item in user_today:
                if item.user_id:
                    user = Client.objects.get(user_id=item.user_id)
                    usertoday.append({
                        'avatar': user.avatar,
                        'userName': user.username,
                    })
            for item in last_user:
                if item.user_id:
                    user = Client.objects.get(user_id=item.user_id)
                    yesterday.append({
                        'avatar': user.avatar,
                        'userName': user.username,
                    })

            dataall.append({
                'code': 200,
                'message': "null",
                'total_sum': total_sum,
                'today_sum': today_sum,
                'yesterday_sum': yesterday_sum,
                'data': {
                    'province_all_top': provinceall,
                    'ip_all_top': ipall,
                    'province_today': provincetoday,
                    'user_today': usertoday,
                    'last_user': yesterday,
                },
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

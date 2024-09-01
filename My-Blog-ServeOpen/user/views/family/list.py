import time

from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.client import Client
from appone.models.family import Family


class FamilyCommentView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        # current 当前页码
        # size    一页显示多少内容
        # 三个搜索条件
        data = request.data
        # print(data)
        current = int(data.get('current', 0))
        size = int(data.get('size', 0))
        status = data.get('status', '')

        dataall = []
        data = []
        # 通过Q来实现不同的搜索操作
        query = Q()
        try:
            if status is not None:
                query &= Q(status=status)
            fam = Family.objects.filter(query)
            total = fam.count()
            familys = fam.order_by('id')[(current - 1) * size:current * size]
            for family in familys:
                data.append({
                    'id': family.id,
                    'userId': family.user_id_id,
                    'bgCover': family.bg_cover,
                    'manCover': family.man_cover,
                    'womanCover': family.woman_cover,
                    'manName': family.man_name,
                    'womanName': family.woman_name,
                    'timing': family.timing,
                    'countdownTitle': family.countdown_title,
                    'countdownTime': family.countdown_time,
                    'status': family.status,
                    'familyInfo': family.family_info,
                    'likeCount': family.like_count,
                    'createTime': family.create_time,
                    'updateTime': family.update_time,
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'total': total,
                'records': data,
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

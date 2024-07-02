# Create your views here.
import time

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.family import Family
from appone.models.resource import Resource


class GetFamilyView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            id = 9
            data = []
            dataall = []
            family = Family.objects.get(user_id_id=id)
            bgCover = ''
            womanCover = ''
            manCover = ''
            bgCover_r = Resource.objects.filter(path=family.bg_cover)
            womanCover_r = Resource.objects.filter(path=family.woman_cover)
            manCover_r = Resource.objects.filter(path=family.man_cover)
            if bgCover_r.exists():
                if bgCover_r[0].status:
                    bgCover = family.bg_cover
            # else:
            #     bgCover = family.bg_cover

            if womanCover_r.exists():
                if womanCover_r[0].status:
                    womanCover = family.woman_cover
            # else:
            #     womanCover = family.bg_cover

            if manCover_r.exists():
                if manCover_r[0].status:
                    manCover = family.man_cover
            # else:
            #     manCover = family.bg_cover

            if family.status:
                data.append({
                    'id': family.id,
                    'userId': family.user_id_id,
                    'bgCover': bgCover,
                    'status': family.status,
                    'countdownTime': family.countdown_time,
                    'countdownTitle': family.countdown_title,
                    'createTime': family.create_time,
                    'likeCount': family.like_count,
                    'familyInfo': family.family_info,
                    'timing': family.timing,
                    'womanName': family.woman_name,
                    'manName': family.man_name,
                    'womanCover': womanCover,
                    'manCover': manCover,
                    'updateTime': family.update_time,
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

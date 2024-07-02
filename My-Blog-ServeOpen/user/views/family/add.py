import time
from django.contrib.auth.models import User
from django.db.models import Q
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.comment import Comment
from appone.models.family import Family


class AddFamilyView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            userId = data['userId']
            bgCover = data['bgCover']
            manCover = data['manCover']
            womanCover = data['womanCover']
            manName = data['manName']
            womanName = data['womanName']
            timing = data['timing']
            countdownTitle = data['countdownTitle']
            countdownTime = data['countdownTime']
            status = data.get('status', False)
            familyInfo = data['familyInfo']
            likeCount = data.get('likeCount', 0)

            user = User.objects.filter(id=userId)
            if user.exists():
                Family.objects.create(user_id=user[0], bg_cover=bgCover, man_cover=manCover,
                                      woman_cover=womanCover, man_name=manName,
                                      woman_name=womanName, timing=timing,
                                      countdown_title=countdownTitle, countdown_time=countdownTime,
                                      status=status, family_info=familyInfo, like_count=likeCount)
                return Response({
                    'result': "success"
                })
            else:
                return Response({
                    'result': "exist null"
                })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

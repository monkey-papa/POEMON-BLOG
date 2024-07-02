# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny


from appone.models.web_info import WebInfo


class UWebInfoView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.POST
            id = int(data.get('id', 0))
            web_name = data.get('webName', "").strip()
            web_title = data.get('webTitle', "").strip()
            notices = data.get('notices', "").strip()
            footer = data.get('footer', "").strip()
            background_image = data.get('backgroundImage', "").strip()
            avatar = data.get('avatar', "").strip()
            random_avatar = data.get('randomAvatar', "").strip()
            random_name = data.get('randomName', "").strip()
            random_cover = data.get('randomCover', "").strip()
            waifu_json = data.get('waifuJson', "").strip()
            status = data.get('status', "").strip()
            if status == 'true':
                status = 'True'
            if status == 'false':
                status = 'False'
            webinfo = WebInfo.objects.filter(id=id)
            if web_name == '':
                if random_avatar != '':
                    webinfo.update(random_avatar=random_avatar)
                elif random_name != '':
                    webinfo.update(random_name=random_name)
                elif random_cover != '':
                    webinfo.update(random_cover=random_cover)
                elif notices != '':
                    webinfo.update(notices=notices)
                webinfo.update(status=status)
            else:
                webinfo.update(web_name=web_name, web_title=web_title,
                               footer=footer, background_image=background_image,
                               avatar=avatar, waifu_json=waifu_json, status=status)
            return Response({
                'result': "success",
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })
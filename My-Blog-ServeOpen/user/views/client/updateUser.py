import time
from datetime import timedelta, datetime
from django.contrib.auth.models import User
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny
from appone.models.client import Client
from appone.models.code import Code
from appone.models.resource import Resource


class UpdateUserView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['userId']
            username = data.get("username", '')
            gender = data.get('gender', '')
            introduction = data.get('introduction', '')
            phoneNumber = data.get('phoneNumber', '')
            email = data.get('email', '')
            avatar = data.get('avatar', '')
            code = data.get('code', '')
            qiniuDomain = data.get('qiniuDomain', '')
            qiniuAccessKey = data.get('qiniuAccessKey', '')
            qiniuSecretKey = data.get('qiniuSecretKey', '')
            qiniuBucketName = data.get('qiniuBucketName', '')
            client = Client.objects.filter(user_id=id)
            if client.exists():
                if username and username != client[0].username:
                    user = User.objects.filter(username=username)
                    if user.exists():
                        return Response({'results': 'Username already exists.'}, status=400)
                    client.update(username=username)
                    User.objects.filter(id=id).update(username=username)
                if gender in [0, 1, 2]:
                    client.update(gender=gender)
                if introduction:
                    client.update(introduction=introduction)
                if qiniuDomain:
                    client.update(qiniu_domain=qiniuDomain)
                if qiniuAccessKey:
                    client.update(qiniu_access_key=qiniuAccessKey)
                if qiniuSecretKey:
                    client.update(qiniu_secret_key=qiniuSecretKey)
                if qiniuBucketName:
                    client.update(qiniu_bucket_name=qiniuBucketName)
                if phoneNumber:
                    client.update(phone_number=phoneNumber)
                if email:
                    code_log = Code.objects.filter(email=email).order_by('-id')
                    if code_log:
                        newly_log = code_log[0]
                        times = datetime.now() - timedelta(hours=0, minutes=5, seconds=0)
                        if newly_log.create_time.replace(tzinfo=None) >= times:
                            if newly_log.code == code:
                                client.update(email=email)
                if avatar:
                    client.update(avatar=avatar)
            ct = Client.objects.get(user_id=id)
            data = []
            dataall = []
            avatar_a = ''
            r = Resource.objects.filter(path=ct.avatar)
            if r.exists():
                if r[0].status:
                    avatar_a = ct.avatar
            # else:
            #     avatar_a = ct.avatar
            data.append({
                'id': ct.user_id,
                'username': ct.username,
                'phoneNumber': ct.phone_number,
                'email': ct.email,
                'admire': ct.admire,
                'userStatus': ct.user_status,
                'avatar': avatar_a,
                'gender': ct.gender,
                'introduction': ct.introduction,
                'userType': ct.user_type,
                'createTime': ct.create_time,
                'qiniuDomain': ct.qiniu_domain,
                'qiniuBucketName': ct.qiniu_bucket_name,
                'qiniuSecretKey': ct.qiniu_secret_key,
                'qiniuAccessKey': ct.qiniu_access_key,
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

# Create your views here.
import time

from django.contrib.auth.models import User
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article
from appone.models.article_like import Article_like


class LikeView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            userId = data['userId']
            article_id = data['id']
            user = User.objects.get(id=userId)
            article = Article.objects.get(id=article_id)
            article_like = Article_like.objects.filter(user_id_id=userId, article_id_id=article_id)
            if user and article:
                if article_like.exists():
                    article.like_count -= 1
                    article.save()
                    Article_like.objects.get(user_id_id=userId, article_id_id=article_id).delete()

                else:
                    Article_like.objects.create(user_id_id=userId, article_id_id=article_id)
                    article.like_count += 1
                    article.save()
                return Response({
                    'result': 'success'
                })
            else:
                return Response({
                    'failure': "exists null",
                })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

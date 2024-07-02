# Create your views here.
import time

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article


class GetArtView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            data = []
            dataall = []
            if id:
                article = Article.objects.get(id=id)
                data.append({
                    'id': article.id,
                    'userId': article.user_id,
                    'sortId': article.sort_id,
                    'labelId': article.label_id,
                    'articleCover': article.article_cover,
                    'articleTitle': article.article_title,
                    'articleContent': article.article_content,
                    'articleAuthor': article.update_by,
                    'viewStatus': article.view_status,
                    'password': article.password,
                    'recommendStatus': article.recommend_status,
                    'commentStatus': article.comment_status,
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
            else:
                return Response({
                    'failure': "exists null",
                })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

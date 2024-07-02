# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article


class UpdateArtView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
            id = data['id']
            article_title = data['articleTitle']
            user_id = data['userId']
            update_by = data['articleAuthor']
            article_content = data['articleContent']
            recommend_status = data['recommendStatus']
            view_status = data['viewStatus']
            password = data['password']
            article_cover = data['articleCover']
            sort_id = data['sortId']
            label_id = data['labelId']
            if id:
                article = Article.objects.filter(id=id)
                if article_title is not None:
                    article.update(article_title=article_title)
                if user_id is not None:
                    article.update(user_id=user_id)
                if update_by is not None:
                    article.update(update_by=update_by)
                if article_content is not None:
                    article.update(article_content=article_content)
                if recommend_status is not None:
                    article.update(recommend_status=recommend_status)
                if view_status is not None:
                    article.update(view_status=view_status)
                if password is not None:
                    article.update(password=password)
                if article_cover is not None:
                    article.update(article_cover=article_cover)
                if label_id is not None:
                    article.update(label_id=label_id)
                if sort_id is not None:
                    article.update(sort_id=sort_id)

                return Response({
                    'result': "success",
                })
            else:
                return Response({
                    'failure': "exists null",
                })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

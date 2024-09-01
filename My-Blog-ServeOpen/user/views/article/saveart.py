# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article


class SaveArtView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        try:
            data = request.data
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
            user_ids = Article.objects.values_list('user_id', flat=True)
            print(user_ids)
            for userid in user_ids:
                print(userid)

            Article.objects.create(article_title=article_title, update_by=update_by,
                                   article_content=article_content, recommend_status=recommend_status,
                                   view_status=view_status, password=password, article_cover=article_cover,
                                   sort_id=sort_id, label_id=label_id, user_id=user_id)

            return Response({
                'result': "success",
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article


class UpdateView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = data.get('articleId', '')
            viewStatus = data.get('viewStatus', None)
            recommendStatus = data.get('recommendStatus', None)
            commentStatus = data.get('commentStatus', None)
            if id:
                article = Article.objects.filter(id=id)
                if viewStatus is not None:
                    viewStatuss = viewStatus.capitalize()
                    article.update(view_status=viewStatuss)
                if commentStatus is not None:
                    commentStatuss = commentStatus.capitalize()
                    article.update(comment_status=commentStatuss)
                if recommendStatus is not None:
                    recommendStatuss = recommendStatus.capitalize()
                    article.update(recommend_status=recommendStatuss)
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

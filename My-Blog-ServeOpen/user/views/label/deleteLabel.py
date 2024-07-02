# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article
from appone.models.label import Label

class DeleteLabelView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            if id:
                label = Label.objects.get(id=id)
                articles = Article.objects.filter(label_id=label.id)
                if articles:
                    for article in articles:
                        article.update(label_id=23)
                Label.objects.get(id=id).delete()
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

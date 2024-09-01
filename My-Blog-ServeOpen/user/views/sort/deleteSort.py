# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article
from appone.models.label import Label
from appone.models.sort import Sort

class DeleteSortView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            id = int(data.get('id', 0))
            if id:
                sort = Sort.objects.get(id=id)
                labels = Label.objects.filter(sort_id=sort.id)
                articles = Article.objects.filter(sort_id=sort.id)
                if articles:
                    for article in articles:
                        article.update(sort_id=11)
                if labels:
                    for label in labels:
                        label.update(sort_id=11)
                Sort.objects.get(id=id).delete()
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

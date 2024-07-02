import time

from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from appone.models.article import Article
from appone.models.label import Label
from appone.models.sort import Sort


class GetSortInfoView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        dataall = []
        data = []
        try:
            sorts = Sort.objects.all().order_by('id')
            a = []
            for sort in sorts:
                labels = Label.objects.filter(sort_id_id=sort.id)
                for label in labels:
                    a.append({
                        'countOfLabel': Article.objects.filter(label_id_id=label.id).count(),
                        'id': label.id,
                        'labelDescription': label.label_name,
                        'labelName': label.label_name,
                        'sortId': label.sort_id_id,
                    })
                data.append({
                    'countOfSort': Article.objects.filter(sort_id_id=sort.id).count(),
                    'id': sort.id,
                    'labels': a,
                    'priority': sort.priority,
                    'sortDescription': sort.sort_description,
                    'sortName': sort.sort_name,
                    'sortType': sort.sort_type,
                })
                a = []
            dataall.append({
                'code': 200,
                'message': "null",
                'data': data,
                'currentTimeMillis': time.time(),
            })
        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

        return Response({'result': dataall})

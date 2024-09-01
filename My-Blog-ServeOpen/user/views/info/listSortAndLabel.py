import time

from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article
from appone.models.label import Label
from appone.models.sort import Sort


class ListSortAndLabelView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        dataall = []
        data = []
        try:
            sorts = Sort.objects.all().order_by('id')
            labels = Label.objects.all().order_by('id')
            sors = []
            labs = []
            for sort in sorts:
                sors.append({
                    'countOfSort': Article.objects.filter(sort_id=sort.id).count(),
                    'id': sort.id,
                    'sortName': sort.sort_name,
                    'sortType': sort.sort_type,
                    'sortDescription': sort.sort_description,
                    'priority': sort.priority,
                })
            for label in labels:
                labs.append({
                    'countOfLabel': Article.objects.filter(label_id=label.id).count(),
                    'id': label.id,
                    'labelDescription': label.label_description,
                    'labelName': label.label_name,
                    'sortId': label.sort_id,
                })

            data.append({
                'labels': labs,
                'sorts': sors,
            })
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

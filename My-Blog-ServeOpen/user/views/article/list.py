import time

from django.db.models import Q
from django.contrib.auth.models import User
from django.shortcuts import render
# Create your views here.

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import IsAdminUser

from appone.models.article import Article
from appone.models.client import Client
from appone.models.comment import Comment

from appone.models.label import Label
from appone.models.sort import Sort


class ListView(APIView):
    permission_classes = [IsAdminUser]
    authentication_classes = [TokenAuthentication]

    def post(self, request):
        data = request.data
        current = int(data.get('current', 0))
        size = int(data.get('size', 0))
        searchKey = data.get('searchKey', '').strip()
        recommendStatus = data['recommendStatus']
        labelId = data['labelId']
        sortId = data['sortId']
        dataall = []
        data = []
        # 通过Q来实现不同的搜索操作
        query = Q()
        try:
            if labelId:
                query &= Q(label_id=labelId)
            if sortId:
                query &= Q(sort_id=sortId)
            if recommendStatus is not None:
                query &= Q(recommend_status=recommendStatus)
            if searchKey:
                query &= Q(article_title__contains=searchKey)
            art = Article.objects.filter(query)
            total = art.count()
            articles = art[(current - 1) * size:current * size]
            for article in articles:
                labels = Label.objects.get(id=article.label_id)
                sorts = Sort.objects.get(id=article.sort_id)
                label = []
                sort = []
                label.append({
                    'countOfLabel': Article.objects.filter(label_id=article.label_id).count(),
                    'id': labels.id,
                    'labelDescription': labels.label_description,
                    'labelName': labels.label_name,
                    'sortId': labels.sort_id,
                })

                sort.append({
                    'countOfSort': Article.objects.filter(sort_id=article.sort_id).count(),
                    'id': sorts.id,
                    'labels': Label.objects.filter(sort_id=article.sort_id).count(),
                    'priority': sorts.priority,
                    'sortDescription': sorts.sort_description,
                    'sortName': sorts.sort_name,
                    'sortType': sorts.sort_type
                })

                data.append({
                    'username': User.objects.get(id=article.user_id).username,
                    'commentCount': Comment.objects.filter(source=article.id).count(),
                    'label': label,
                    'sort': sort,
                    'id': article.id,
                    'userId': article.user_id,
                    'sortId': article.sort_id,
                    'labelId': article.label_id,
                    'articleCover': article.article_cover,
                    'articleTitle': article.article_title,
                    'articleContent': article.article_content,
                    'viewCount': article.view_count,
                    'likeCount': article.like_count,
                    'viewStatus': article.view_status,
                    'password': article.password,
                    'recommendStatus': article.recommend_status,
                    'commentStatus': article.comment_status,
                    'createTime': article.create_time,
                    'updateTime': article.update_time,
                    'updateBy': article.update_by,
                    'deleted': article.deleted,
                })

            dataall.append({
                'code': 200,
                'message': "null",
                'total': total,
                'records': data,
                'currentTimeMillis': time.time(),
            })
            return Response({
                'result': dataall
            })

        except Exception as error:
            return Response({
                'result': "failure {0}".format(error)
            })

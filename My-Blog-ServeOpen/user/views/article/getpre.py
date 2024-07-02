# Create your views here.
import time

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.authentication import TokenAuthentication
from rest_framework.permissions import AllowAny

from django.contrib.auth.models import User

from appone.models.article import Article
from appone.models.article_like import Article_like
from appone.models.comment import Comment
from appone.models.label import Label
from appone.models.resource import Resource
from appone.models.sort import Sort


class GetArtPreView(APIView):
    permission_classes = [AllowAny]
    authentication_classes = [TokenAuthentication]

    def get(self, request):
        try:
            data = request.GET
            count = 0
            id = int(data.get('id', 0))
            userId = int(data.get('userId', 0))
            status = data.get('flag', 0)
            if status == 'true':
                status = True
            if status == 'false':
                status = False
            data = []
            dataall = []
            articl = Article.objects.filter(id=id)
            if Article_like.objects.filter(user_id_id=userId, article_id_id=id).exists():
                count = 1
            else:
                pass
            vc = articl[0].view_count + 1
            articl.update(view_count=vc)
            if articl.exists():
                if articl[0].view_status == status:
                    article = Article.objects.get(id=id)
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

                    avatar_a = ''
                    r = Resource.objects.filter(path=article.article_cover)
                    if r.exists():
                        if r[0].status:
                            avatar_a = article.article_cover
                    # else:
                    #     avatar_a = article.article_cover

                    data.append({
                        'id': article.id,
                        'userId': article.user_id,
                        'sortId': article.sort_id,
                        'labelId': article.label_id,
                        'articleCover': avatar_a,
                        'articleTitle': article.article_title,
                        'articleContent': article.article_content,
                        'articleAuthor': article.update_by,
                        'viewStatus': article.view_status,
                        'password': article.password,
                        'recommendStatus': article.recommend_status,
                        'commentStatus': article.comment_status,
                        'createTime': article.create_time,
                        'likeCount': article.like_count,
                        'updateBy': article.update_by,
                        'updateTime': article.update_time,
                        'username': User.objects.get(id=article.user_id).username,
                        'commentCount': Comment.objects.filter(source=article.id).count(),
                        'viewCount': article.view_count,
                        'label': label,
                        'sort': sort,
                        'articleLikeStatus': count,
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

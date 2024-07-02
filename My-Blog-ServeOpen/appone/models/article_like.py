from django.db import models
from django.contrib.auth.models import User

from appone.models.article import Article


class Article_like(models.Model):
    user_id = models.ForeignKey(User, verbose_name='用户ID', on_delete=models.CASCADE)
    article_id = models.ForeignKey(Article, verbose_name='文章ID', on_delete=models.CASCADE)

    class Meta:
        db_table = 'like'
        verbose_name = '点赞'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.user_id_id) + str(self.article_id_id)

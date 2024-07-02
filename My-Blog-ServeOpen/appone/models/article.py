from django.contrib.auth.models import User
from django.db import models

from appone.models.label import Label
from appone.models.sort import Sort

class Article(models.Model):
    user_id = models.IntegerField(verbose_name='用户ID')
    sort_id = models.IntegerField(verbose_name='分类ID')
    label_id = models.IntegerField(verbose_name='标签ID')
    article_cover = models.URLField(max_length=256, null=True, blank=True, verbose_name='封面')
    article_title = models.CharField(max_length=32, verbose_name='博文标题')
    article_content = models.TextField(verbose_name='博文内容')
    view_count = models.IntegerField(default=0, verbose_name='浏览量')
    like_count = models.IntegerField(default=0, verbose_name='点赞数')
    view_status = models.BooleanField(default=True, verbose_name='是否可见')
    password = models.CharField(max_length=128, null=True, blank=True, verbose_name='密码')
    recommend_status = models.BooleanField(default=False, verbose_name='是否推荐')
    comment_status = models.BooleanField(default=True, verbose_name='是否启用评论')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')
    update_time = models.DateTimeField(auto_now=True, verbose_name='最终修改时间')
    update_by = models.CharField(max_length=32, null=True, blank=True, verbose_name='最终修改人')
    deleted = models.BooleanField(default=False, verbose_name='是否删除')

    class Meta:
        db_table = 'article'
        verbose_name = '文章表'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.article_title)


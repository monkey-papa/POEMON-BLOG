from django.db import models

class WeiYan(models.Model):
    id = models.AutoField(primary_key=True, verbose_name='id')
    user_id = models.IntegerField(verbose_name='用户ID')
    like_count = models.IntegerField(default=0, verbose_name='点赞数')
    content = models.CharField(max_length=1024, verbose_name='内容')
    type = models.CharField(max_length=32, verbose_name='类型')
    source = models.IntegerField(null=True, blank=True, verbose_name='来源标识')
    is_public = models.BooleanField(default=False, verbose_name='是否公开[0:仅自己可见，1:所有人可见]')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    class Meta:
        db_table = 'wei_yan'
        verbose_name = '微言表'
        verbose_name_plural = verbose_name
        indexes = [
            models.Index(fields=['user_id'])
        ]

    def __str__(self):
        return str(self.id) + '-' + str(self.content) + '-' + str(self.user_id)

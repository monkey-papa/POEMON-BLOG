from django.db import models
from django.contrib.auth.models import User

class Resource(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE, verbose_name='用户ID')
    type = models.CharField(max_length=32, verbose_name='资源类型')
    path = models.CharField(max_length=256, verbose_name='资源路径')
    size = models.IntegerField(null=True, blank=True, verbose_name='资源内容的大小，单位：字节')
    mime_type = models.CharField(max_length=256, null=True, blank=True, verbose_name='资源的MIME类型')
    status = models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    class Meta:
        db_table = 'resource'
        verbose_name = '资源信息'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.path) + '-' + str(self.user_id)

from django.db import models


class ResourcePath(models.Model):
    title = models.CharField(max_length=64, verbose_name='标题')
    classify = models.CharField(max_length=32, null=True, blank=True, verbose_name='分类')
    cover = models.CharField(max_length=256, null=True, blank=True, verbose_name='封面')
    url = models.CharField(max_length=256, null=True, blank=True, verbose_name='链接')
    introduction = models.CharField(max_length=1024, null=True, blank=True, verbose_name='简介')
    type = models.CharField(max_length=32, verbose_name='资源类型')
    status = models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')
    remark = models.TextField(null=True, blank=True, verbose_name='备注')
    # avator = models.CharField(max_length=256, null=True, blank=True, verbose_name='链接')

    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    class Meta:
        db_table = 'resource_path'
        verbose_name = '资源路径'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.title)

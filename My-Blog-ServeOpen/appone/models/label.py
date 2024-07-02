from django.db import models

from appone.models.sort import Sort


class Label(models.Model):
    sort_id = models.IntegerField(verbose_name='分类ID')
    label_name = models.CharField(max_length=32, verbose_name='标签名称')
    label_description = models.CharField(max_length=256, verbose_name='标签描述')

    class Meta:
        db_table = 'label'
        verbose_name = '标签'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.label_name) + '-' + str(self.sort_id)
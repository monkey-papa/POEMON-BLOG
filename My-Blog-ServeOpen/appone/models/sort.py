from django.db import models

class Sort(models.Model):
    sort_name = models.CharField(max_length=32, verbose_name='分类名称')
    sort_description = models.CharField(max_length=256, verbose_name='分类描述')
    sort_type = models.IntegerField(default=1, verbose_name='分类类型', choices=((0, '导航栏分类'), (1, '普通分类')))
    priority = models.IntegerField(null=True, blank=True, verbose_name='导航栏分类优先级：数字小的在前面')

    class Meta:
        db_table = 'sort'
        verbose_name = '分类'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.sort_name)
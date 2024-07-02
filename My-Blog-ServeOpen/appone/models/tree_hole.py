from django.db import models

class TreeHole(models.Model):
    avatar = models.CharField(max_length=256, null=True, blank=True, verbose_name='头像')
    username = models.CharField(max_length=32, unique=True, verbose_name='用户名')
    message = models.CharField(max_length=64, verbose_name='留言')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    class Meta:
        db_table = 'tree_hole'
        verbose_name = '树洞'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.message)
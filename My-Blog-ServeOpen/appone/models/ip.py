from django.db import models

class Ip(models.Model):
    user_id = models.IntegerField(default=0, verbose_name='用户id')
    ip = models.CharField(max_length=20, null=True, blank=True, verbose_name='IP')
    province = models.CharField(max_length=20, null=True, blank=True, verbose_name='省')
    city = models.CharField(max_length=20, null=True, blank=True, verbose_name='市')
    create_time = models.DateTimeField(verbose_name='修改时间now')
    k_time = models.DateTimeField(verbose_name='修改时间', null=True, blank=True)
    mcount = models.IntegerField(default=0, verbose_name='某时间段内登录次数')
    dcount = models.IntegerField(default=0, verbose_name='登录天数')

    class Meta:
        db_table = 'ip'
        verbose_name = 'IP地址表'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.user_id) + '-' + str(self.ip)

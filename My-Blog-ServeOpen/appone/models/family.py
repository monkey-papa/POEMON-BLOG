from django.db import models
from django.contrib.auth.models import User

class Family(models.Model):
    id = models.AutoField(primary_key=True, verbose_name='id')
    user_id = models.ForeignKey(User, on_delete=models.CASCADE, verbose_name='用户ID')
    bg_cover = models.CharField(max_length=256, verbose_name='背景封面')
    man_cover = models.CharField(max_length=256, verbose_name='男生头像')
    woman_cover = models.CharField(max_length=256, verbose_name='女生头像')
    man_name = models.CharField(max_length=32, verbose_name='男生昵称')
    woman_name = models.CharField(max_length=32, verbose_name='女生昵称')
    timing = models.CharField(max_length=32, verbose_name='计时')
    countdown_title = models.CharField(max_length=32, null=True, blank=True, verbose_name='倒计时标题')
    countdown_time = models.CharField(max_length=32, null=True, blank=True, verbose_name='倒计时时间')
    status = models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')
    family_info = models.CharField(max_length=1024, null=True, blank=True, verbose_name='额外信息')
    like_count = models.IntegerField(default=0, verbose_name='点赞数')

    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')
    update_time = models.DateTimeField(auto_now=True, verbose_name='最终修改时间')

    class Meta:
        db_table = 'family'
        verbose_name = '家庭信息'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.create_time) + '-' + str(self.user_id)

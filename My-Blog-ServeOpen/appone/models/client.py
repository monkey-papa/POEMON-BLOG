from django.db import models
from django.contrib.auth.models import User

class Client(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    username = models.CharField(max_length=32, unique=True, verbose_name='用户名')
    # password = models.CharField(max_length=128, verbose_name='密码')
    phone_number = models.CharField(max_length=16, null=True, blank=True, verbose_name='手机号')
    email = models.CharField(max_length=32, null=True, blank=True, verbose_name='用户邮箱')
    user_status = models.BooleanField(default=True, verbose_name='是否启用')
    gender_choice = (
        (0, '保密'),
        (1, '男'),
        (2, '女'),
    )
    gender = models.SmallIntegerField(null=True, blank=True, choices=gender_choice, default=0, verbose_name='性别')
    open_id = models.CharField(max_length=128, null=True, blank=True, verbose_name='openId')
    avatar = models.URLField(max_length=256, null=True, blank=True, verbose_name='头像')
    admire = models.CharField(max_length=32, null=True, blank=True, verbose_name='赞赏')
    introduction = models.CharField(max_length=4096, null=True, blank=True, verbose_name='简介')
    user_type_choice = (
        (0, 'Boss'),
        (1, '管理员'),
        (2, '普通用户'),
    )
    user_type = models.SmallIntegerField(choices=user_type_choice, default=2, verbose_name='用户类型')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')
    update_time = models.DateTimeField(auto_now=True, verbose_name='最终修改时间')
    update_by = models.CharField(max_length=32, null=True, blank=True, verbose_name='最终修改人')
    deleted = models.BooleanField(default=False, verbose_name='是否删除')

    class Meta:
        db_table = 'user'
        verbose_name = '用户信息表'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.user_id) + '-' + str(self.username)
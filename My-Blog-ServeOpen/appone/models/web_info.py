from django.db import models

class WebInfo(models.Model):
    web_name = models.CharField(max_length=16, verbose_name='网站名称')
    web_title = models.CharField(max_length=512, verbose_name='网站信息')
    notices = models.CharField(max_length=512, null=True, blank=True, verbose_name='公告')
    footer = models.CharField(max_length=256, verbose_name='页脚')
    background_image = models.CharField(max_length=256, null=True, blank=True, verbose_name='背景')
    avatar = models.CharField(max_length=256, verbose_name='头像')
    random_avatar = models.TextField(null=True, blank=True, verbose_name='随机头像')
    random_name = models.CharField(max_length=4096, null=True, blank=True, verbose_name='随机名称')
    random_cover = models.TextField(null=True, blank=True, verbose_name='随机封面')
    waifu_json = models.TextField(null=True, blank=True, verbose_name='看板娘消息')
    status = models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')

    class Meta:
        db_table = 'web_info'
        verbose_name = '网站信息表'
        verbose_name_plural = verbose_name

    def __str__(self):
        return str(self.id) + '-' + str(self.web_name)

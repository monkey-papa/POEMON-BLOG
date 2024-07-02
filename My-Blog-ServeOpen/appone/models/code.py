from django.db import models
# Create your models here.

class Code(models.Model):
    email = models.EmailField(verbose_name="email", blank=True)
    code = models.CharField(max_length=10, verbose_name="验证码")
    create_time = models.DateTimeField(verbose_name='生成时间', auto_now_add=True)

    def __str__(self):
        return str(self.email)


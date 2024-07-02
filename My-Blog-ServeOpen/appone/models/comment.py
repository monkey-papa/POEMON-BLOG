from django.db import models

class Comment(models.Model):
    source = models.IntegerField(verbose_name='评论来源标识')
    type = models.CharField(max_length=32, verbose_name='评论来源类型')
    parent_comment_id = models.IntegerField(null=True, blank=True, verbose_name='父评论ID')
    user_id = models.IntegerField(verbose_name='发表用户ID')
    floor_comment_id = models.IntegerField(null=True, blank=True, verbose_name='楼层评论ID')
    parent_user_id = models.IntegerField(null=True, blank=True, verbose_name='父发表用户名ID')
    like_count = models.IntegerField(default=0, verbose_name='点赞数')
    comment_content = models.CharField(max_length=1024, verbose_name='评论内容')
    comment_info = models.CharField(max_length=256, null=True, blank=True, verbose_name='评论额外信息')
    create_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    class Meta:
        db_table = 'comment'
        verbose_name = '文章评论表'
        verbose_name_plural = verbose_name
        indexes = [
            models.Index(fields=['source'], name='idx_comment_source'),
        ]

    def __str__(self):
        return str(self.id) + '-' + str(self.comment_content)

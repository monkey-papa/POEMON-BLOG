# Generated by Django 4.1.5 on 2023-08-17 08:56

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name='Article',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('user_id', models.IntegerField(verbose_name='用户ID')),
                ('sort_id', models.IntegerField(verbose_name='分类ID')),
                ('label_id', models.IntegerField(verbose_name='标签ID')),
                ('article_cover', models.URLField(blank=True, max_length=256, null=True, verbose_name='封面')),
                ('article_title', models.CharField(max_length=32, verbose_name='博文标题')),
                ('article_content', models.TextField(verbose_name='博文内容')),
                ('view_count', models.IntegerField(default=0, verbose_name='浏览量')),
                ('like_count', models.IntegerField(default=0, verbose_name='点赞数')),
                ('view_status', models.BooleanField(default=True, verbose_name='是否可见')),
                ('password', models.CharField(blank=True, max_length=128, null=True, verbose_name='密码')),
                ('recommend_status', models.BooleanField(default=False, verbose_name='是否推荐')),
                ('comment_status', models.BooleanField(default=True, verbose_name='是否启用评论')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='最终修改时间')),
                ('update_by', models.CharField(blank=True, max_length=32, null=True, verbose_name='最终修改人')),
                ('deleted', models.BooleanField(default=False, verbose_name='是否删除')),
            ],
            options={
                'verbose_name': '文章表',
                'verbose_name_plural': '文章表',
                'db_table': 'article',
            },
        ),
        migrations.CreateModel(
            name='Article_like',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
            ],
            options={
                'verbose_name': '点赞',
                'verbose_name_plural': '点赞',
                'db_table': 'like',
            },
        ),
        migrations.CreateModel(
            name='Client',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('username', models.CharField(max_length=32, unique=True, verbose_name='用户名')),
                ('phone_number', models.CharField(blank=True, max_length=16, null=True, verbose_name='手机号')),
                ('email', models.CharField(blank=True, max_length=32, null=True, verbose_name='用户邮箱')),
                ('user_status', models.BooleanField(default=True, verbose_name='是否启用')),
                ('gender', models.SmallIntegerField(blank=True, choices=[(0, '保密'), (1, '男'), (2, '女')], default=0, null=True, verbose_name='性别')),
                ('open_id', models.CharField(blank=True, max_length=128, null=True, verbose_name='openId')),
                ('avatar', models.URLField(blank=True, max_length=256, null=True, verbose_name='头像')),
                ('admire', models.CharField(blank=True, max_length=32, null=True, verbose_name='赞赏')),
                ('introduction', models.CharField(blank=True, max_length=4096, null=True, verbose_name='简介')),
                ('user_type', models.SmallIntegerField(choices=[(0, 'Boss'), (1, '管理员'), (2, '普通用户')], default=2, verbose_name='用户类型')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='最终修改时间')),
                ('update_by', models.CharField(blank=True, max_length=32, null=True, verbose_name='最终修改人')),
                ('deleted', models.BooleanField(default=False, verbose_name='是否删除')),
            ],
            options={
                'verbose_name': '用户信息表',
                'verbose_name_plural': '用户信息表',
                'db_table': 'user',
            },
        ),
        migrations.CreateModel(
            name='Code',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('email', models.EmailField(blank=True, max_length=254, verbose_name='email')),
                ('code', models.CharField(max_length=10, verbose_name='验证码')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='生成时间')),
            ],
        ),
        migrations.CreateModel(
            name='Comment',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('source', models.IntegerField(verbose_name='评论来源标识')),
                ('type', models.CharField(max_length=32, verbose_name='评论来源类型')),
                ('parent_comment_id', models.IntegerField(blank=True, null=True, verbose_name='父评论ID')),
                ('user_id', models.IntegerField(verbose_name='发表用户ID')),
                ('floor_comment_id', models.IntegerField(blank=True, null=True, verbose_name='楼层评论ID')),
                ('parent_user_id', models.IntegerField(blank=True, null=True, verbose_name='父发表用户名ID')),
                ('like_count', models.IntegerField(default=0, verbose_name='点赞数')),
                ('comment_content', models.CharField(max_length=1024, verbose_name='评论内容')),
                ('comment_info', models.CharField(blank=True, max_length=256, null=True, verbose_name='评论额外信息')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '文章评论表',
                'verbose_name_plural': '文章评论表',
                'db_table': 'comment',
            },
        ),
        migrations.CreateModel(
            name='Family',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False, verbose_name='id')),
                ('bg_cover', models.CharField(max_length=256, verbose_name='背景封面')),
                ('man_cover', models.CharField(max_length=256, verbose_name='男生头像')),
                ('woman_cover', models.CharField(max_length=256, verbose_name='女生头像')),
                ('man_name', models.CharField(max_length=32, verbose_name='男生昵称')),
                ('woman_name', models.CharField(max_length=32, verbose_name='女生昵称')),
                ('timing', models.CharField(max_length=32, verbose_name='计时')),
                ('countdown_title', models.CharField(blank=True, max_length=32, null=True, verbose_name='倒计时标题')),
                ('countdown_time', models.CharField(blank=True, max_length=32, null=True, verbose_name='倒计时时间')),
                ('status', models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')),
                ('family_info', models.CharField(blank=True, max_length=1024, null=True, verbose_name='额外信息')),
                ('like_count', models.IntegerField(default=0, verbose_name='点赞数')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
                ('update_time', models.DateTimeField(auto_now=True, verbose_name='最终修改时间')),
            ],
            options={
                'verbose_name': '家庭信息',
                'verbose_name_plural': '家庭信息',
                'db_table': 'family',
            },
        ),
        migrations.CreateModel(
            name='Ip',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('user_id', models.IntegerField(default=0, verbose_name='用户id')),
                ('ip', models.CharField(blank=True, max_length=20, null=True, verbose_name='IP')),
                ('province', models.CharField(blank=True, max_length=20, null=True, verbose_name='省')),
                ('city', models.CharField(blank=True, max_length=20, null=True, verbose_name='市')),
                ('create_time', models.DateTimeField(verbose_name='修改时间now')),
                ('k_time', models.DateTimeField(blank=True, null=True, verbose_name='修改时间')),
                ('mcount', models.IntegerField(default=0, verbose_name='某时间段内登录次数')),
                ('dcount', models.IntegerField(default=0, verbose_name='登录天数')),
            ],
            options={
                'verbose_name': 'IP地址表',
                'verbose_name_plural': 'IP地址表',
                'db_table': 'ip',
            },
        ),
        migrations.CreateModel(
            name='Label',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('sort_id', models.IntegerField(verbose_name='分类ID')),
                ('label_name', models.CharField(max_length=32, verbose_name='标签名称')),
                ('label_description', models.CharField(max_length=256, verbose_name='标签描述')),
            ],
            options={
                'verbose_name': '标签',
                'verbose_name_plural': '标签',
                'db_table': 'label',
            },
        ),
        migrations.CreateModel(
            name='Resource',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('type', models.CharField(max_length=32, verbose_name='资源类型')),
                ('path', models.CharField(max_length=256, verbose_name='资源路径')),
                ('size', models.IntegerField(blank=True, null=True, verbose_name='资源内容的大小，单位：字节')),
                ('mime_type', models.CharField(blank=True, max_length=256, null=True, verbose_name='资源的MIME类型')),
                ('status', models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '资源信息',
                'verbose_name_plural': '资源信息',
                'db_table': 'resource',
            },
        ),
        migrations.CreateModel(
            name='ResourcePath',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('title', models.CharField(max_length=64, verbose_name='标题')),
                ('classify', models.CharField(blank=True, max_length=32, null=True, verbose_name='分类')),
                ('cover', models.CharField(blank=True, max_length=256, null=True, verbose_name='封面')),
                ('url', models.CharField(blank=True, max_length=256, null=True, verbose_name='链接')),
                ('introduction', models.CharField(blank=True, max_length=1024, null=True, verbose_name='简介')),
                ('type', models.CharField(max_length=32, verbose_name='资源类型')),
                ('status', models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')),
                ('remark', models.TextField(blank=True, null=True, verbose_name='备注')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '资源路径',
                'verbose_name_plural': '资源路径',
                'db_table': 'resource_path',
            },
        ),
        migrations.CreateModel(
            name='Sort',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('sort_name', models.CharField(max_length=32, verbose_name='分类名称')),
                ('sort_description', models.CharField(max_length=256, verbose_name='分类描述')),
                ('sort_type', models.IntegerField(choices=[(0, '导航栏分类'), (1, '普通分类')], default=1, verbose_name='分类类型')),
                ('priority', models.IntegerField(blank=True, null=True, verbose_name='导航栏分类优先级：数字小的在前面')),
            ],
            options={
                'verbose_name': '分类',
                'verbose_name_plural': '分类',
                'db_table': 'sort',
            },
        ),
        migrations.CreateModel(
            name='TreeHole',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('avatar', models.CharField(blank=True, max_length=256, null=True, verbose_name='头像')),
                ('username', models.CharField(max_length=32, unique=True, verbose_name='用户名')),
                ('message', models.CharField(max_length=64, verbose_name='留言')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '树洞',
                'verbose_name_plural': '树洞',
                'db_table': 'tree_hole',
            },
        ),
        migrations.CreateModel(
            name='WebInfo',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('web_name', models.CharField(max_length=16, verbose_name='网站名称')),
                ('web_title', models.CharField(max_length=512, verbose_name='网站信息')),
                ('notices', models.CharField(blank=True, max_length=512, null=True, verbose_name='公告')),
                ('footer', models.CharField(max_length=256, verbose_name='页脚')),
                ('background_image', models.CharField(blank=True, max_length=256, null=True, verbose_name='背景')),
                ('avatar', models.CharField(max_length=256, verbose_name='头像')),
                ('random_avatar', models.TextField(blank=True, null=True, verbose_name='随机头像')),
                ('random_name', models.CharField(blank=True, max_length=4096, null=True, verbose_name='随机名称')),
                ('random_cover', models.TextField(blank=True, null=True, verbose_name='随机封面')),
                ('waifu_json', models.TextField(blank=True, null=True, verbose_name='看板娘消息')),
                ('status', models.BooleanField(default=True, verbose_name='是否启用[0:否，1:是]')),
            ],
            options={
                'verbose_name': '网站信息表',
                'verbose_name_plural': '网站信息表',
                'db_table': 'web_info',
            },
        ),
        migrations.CreateModel(
            name='WeiYan',
            fields=[
                ('id', models.AutoField(primary_key=True, serialize=False, verbose_name='id')),
                ('user_id', models.IntegerField(verbose_name='用户ID')),
                ('like_count', models.IntegerField(default=0, verbose_name='点赞数')),
                ('content', models.CharField(max_length=1024, verbose_name='内容')),
                ('type', models.CharField(max_length=32, verbose_name='类型')),
                ('source', models.IntegerField(blank=True, null=True, verbose_name='来源标识')),
                ('is_public', models.BooleanField(default=False, verbose_name='是否公开[0:仅自己可见，1:所有人可见]')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '微言表',
                'verbose_name_plural': '微言表',
                'db_table': 'wei_yan',
            },
        ),
        migrations.CreateModel(
            name='Words',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('avatar', models.CharField(blank=True, max_length=256, null=True, verbose_name='头像')),
                ('username', models.CharField(max_length=32, unique=True, verbose_name='用户名')),
                ('message', models.CharField(max_length=64, verbose_name='留言')),
                ('create_time', models.DateTimeField(auto_now_add=True, verbose_name='创建时间')),
            ],
            options={
                'verbose_name': '禁用词',
                'verbose_name_plural': '禁用词',
                'db_table': 'word',
            },
        ),
        migrations.AddIndex(
            model_name='weiyan',
            index=models.Index(fields=['user_id'], name='wei_yan_user_id_32c565_idx'),
        ),
        migrations.AddField(
            model_name='resource',
            name='user_id',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='用户ID'),
        ),
        migrations.AddField(
            model_name='family',
            name='user_id',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='用户ID'),
        ),
        migrations.AddIndex(
            model_name='comment',
            index=models.Index(fields=['source'], name='idx_comment_source'),
        ),
        migrations.AddField(
            model_name='client',
            name='user',
            field=models.OneToOneField(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL),
        ),
        migrations.AddField(
            model_name='article_like',
            name='article_id',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='appone.article', verbose_name='文章ID'),
        ),
        migrations.AddField(
            model_name='article_like',
            name='user_id',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to=settings.AUTH_USER_MODEL, verbose_name='用户ID'),
        ),
    ]

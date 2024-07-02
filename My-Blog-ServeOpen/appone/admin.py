from django.contrib import admin


# Register your models here.
from appone.models.article import Article
from appone.models.article_like import Article_like
from appone.models.client import Client
from appone.models.code import Code
from appone.models.comment import Comment
from appone.models.family import Family
from appone.models.ip import Ip
from appone.models.label import Label
from appone.models.resource import Resource
from appone.models.resource_path import ResourcePath
from appone.models.sort import Sort
from appone.models.tree_hole import TreeHole
from appone.models.web_info import WebInfo
from appone.models.wei_yan import WeiYan
from appone.models.words import Words

admin.site.register(Client)
admin.site.register(WebInfo)
admin.site.register(Sort)
admin.site.register(Label)
admin.site.register(Article)
admin.site.register(Comment)
admin.site.register(Resource)
admin.site.register(TreeHole)
admin.site.register(Words)
admin.site.register(Family)
admin.site.register(WeiYan)
admin.site.register(ResourcePath)
admin.site.register(Ip)
admin.site.register(Code)
admin.site.register(Article_like)



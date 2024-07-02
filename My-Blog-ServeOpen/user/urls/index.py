from django.urls import path

from user.views.article.delete import DeleteArtView
from user.views.article.get import GetArtView
from user.views.article.getpre import GetArtPreView
from user.views.article.list import ListView
from user.views.article.listpre import ListPreView
from user.views.article.saveart import SaveArtView
from user.views.article.updateart import UpdateArtView
from user.views.article.updatestatus import UpdateView
from user.views.client.changeUserAdmire import ChangeUserAdmireView
from user.views.client.changeUserStatus import ChangeUserStatusView
from user.views.client.changeUserType import ChangeUserTypeView
from user.views.client.client import ClientView
from user.views.client.updateUser import UpdateUserView
from user.views.code.code import CodeView
from user.views.comment.add import AddCommentView
from user.views.comment.client_list import ListClientComView
from user.views.comment.delete import DeleteCommentView
from user.views.comment.list import ListCommentView
from user.views.family.add import AddFamilyView
from user.views.family.changestatus import ChangeFamilyStatusView
from user.views.family.delete import DeleteFamilyView
from user.views.family.get import GetFamilyView
from user.views.family.list import FamilyCommentView
from user.views.family.listrandom import ListFamilyView
from user.views.image.delete_image import DeleteImageView
from user.views.info.getpre import GWebInfoPreView
from user.views.info.listSortAndLabel import ListSortAndLabelView
from user.views.ip.getip import GetIpView
from user.views.ip.list import ListIpView
from user.views.ip.rc import RcView
from user.views.like.like import LikeView
from user.views.map.map import MapView
from user.views.map.user import UserView
from user.views.resource.changestatus import ChangeResourceView
from user.views.resource.listResource import ListResourceView
from user.views.resource.listbackground import ListResourcePreBGView
from user.views.resource.listpre import ListResourcePreView
from user.views.resource.saveresource import SaveResourceView
from user.views.resourcepath.delete import DeleteResourcePathView
from user.views.resourcepath.list import ListResourcePathView
from user.views.resourcepath.listCollect import ListCollectView
from user.views.resourcepath.listpre import ListPrResourcePathView
from user.views.resourcepath.lovelist import ListLoveResourcePathView
from user.views.resourcepath.save import SaveResourcePathView
from user.views.resourcepath.update import UpdateResourcePathView
from user.views.sort.getSortInfo import GetSortInfoView
from user.views.label.deleteLabel import DeleteLabelView
from user.views.login import LoginView
from user.views.image.image import ImageView
from user.views.info.updatewebinfo import UWebInfoView
from user.views.info.getwebinfo import GWebInfoView
from user.views.label.saveLabel import SaveLabelView
from user.views.sort.deleteSort import DeleteSortView
from user.views.sort.saveSort import SaveSortView
from user.views.label.updateLabel import UpdateLabelView
from user.views.sort.updateSort import UpadateSortView
from user.views.treehole.delete import DeleteTreeView
from user.views.treehole.list import ListTreeView
from user.views.treehole.save import SaveTreeView
from user.views.weiyan.changepublic import ChangeWeiYanView
from user.views.weiyan.delete import DeleteWeiYanView
from user.views.weiyan.list import WeiYanView
from user.views.weiyan.listpre import WeiYanPreView
from user.views.weiyan.save import AddWeiYanView
from user.views.weiyan.savenews import AddWeiYanPreView
from user.views.words.add import SaveWordsView
from user.views.words.delete import DeleteWordsView
from user.views.words.list import ListWordsView
from user.views.words.update import UpdateWordsView

urlpatterns = [
    path("api/user/login/", LoginView.as_view(), name="login"),
    path("api/admin/webInfo/getAdminWebInfo/", GWebInfoView.as_view(), name="getwebinfo"),
    path("api/admin/webInfo/updateAdminWebInfo/", UWebInfoView.as_view(), name="updatewebinfo"),
    path("api/admin/user/list/", ClientView.as_view(), name="list"),
    path("api/resource/updateImage/", ImageView.as_view(), name="test"),
    path("api/admin/user/changeUserStatus/", ChangeUserStatusView.as_view(), name="changestatus"),
    path("api/admin/user/changeUserAdmire/", ChangeUserAdmireView.as_view(), name="changeadmire"),
    path("api/admin/user/changeUserType/", ChangeUserTypeView.as_view(), name="changetype"),
    path("api/webInfo/getSortInfo/", GetSortInfoView.as_view(), name="getSortInfo"),
    path("api/webInfo/saveSort/", SaveSortView.as_view(), name="saveSort"),
    path("api/webInfo/updateSort/", UpadateSortView.as_view(), name="updateSort"),
    path("api/webInfo/saveLabel/", SaveLabelView.as_view(), name="saveLabel"),
    path("api/webInfo/updateLabel/", UpdateLabelView.as_view(), name="updateLabel"),
    path("api/webInfo/deleteLabel/", DeleteLabelView.as_view(), name="updateLabel"),
    path("api/webInfo/deleteSort/", DeleteSortView.as_view(), name="updateLabel"),
    path("api/admin/article/boss/list/", ListView.as_view(), name="updateLabel"),
    path("api/webInfo/listSortAndLabel/", ListSortAndLabelView.as_view(), name="updateLabel"),
    path("api/admin/article/changeArticleStatus/", UpdateView.as_view(), name="updateLabel"),
    path("api/admin/article/getArticleById/", GetArtView.as_view(), name="updateLabel"),
    path("api/article/deleteArticle/", DeleteArtView.as_view(), name="updateLabel"),
    path("api/article/saveArticle/", SaveArtView.as_view(), name="updateLabel"),
    path("api/article/updateArticle/", UpdateArtView.as_view(), name="updateLabel"),
    path("api/delArticleImage/", DeleteImageView.as_view(), name="updateLabel"),
    path("api/resource/saveResource/", SaveResourceView.as_view(), name="updateLabel"),
    path("api/resource/listResource/", ListResourceView.as_view(), name="updateLabel"),
    path("api/resource/clistResource/", ListResourcePreView.as_view(), name="updateLabel"),
    path("api/resource/changeResourceStatus/", ChangeResourceView.as_view(), name="updateLabel"),
    path("api/admin/treeHole/boss/list/", ListTreeView.as_view(), name="updateLabel"),
    path("api/webInfo/deleteTreeHole/", DeleteTreeView.as_view(), name="updateLabel"),
    path("api/webInfo/saveTreeHole/", SaveTreeView.as_view(), name="updateLabel"),
    path("api/prohibitedWords/add/", SaveWordsView.as_view(), name="updateLabel"),
    path("api/prohibitedWords/list/", ListWordsView.as_view(), name="updateLabel"),
    path("api/prohibitedWords/delete/", DeleteWordsView.as_view(), name="updateLabel"),
    path("api/prohibitedWords/update/", UpdateWordsView.as_view(), name="updateLabel"),
    path("api/admin/comment/boss/list/", ListCommentView.as_view(), name="updateLabel"),
    path("api/admin/comment/boss/deleteComment/", DeleteCommentView.as_view(), name="updateLabel"),
    path("api/admin/comment/boss/addComment/", AddCommentView.as_view(), name="updateLabel"),
    path("api/comment/listComment/", ListClientComView.as_view(), name="updateLabel"),
    path("api/family/listFamily/", FamilyCommentView.as_view(), name="updateLabel"),
    path("api/family/addFamily/", AddFamilyView.as_view(), name="updateLabel"),
    path("api/family/deleteFamily/", DeleteFamilyView.as_view(), name="updateLabel"),
    path("api/family/changeLoveStatus/", ChangeFamilyStatusView.as_view(), name="updateLabel"),
    path("api/family/getAdminFamily/", GetFamilyView.as_view(), name="updateLabel"),
    path("api/family/listRandomFamily/", ListFamilyView.as_view(), name="updateLabel"),
    path("api/weiYan/listWeiYan/", WeiYanView.as_view(), name="updateLabel"),
    path("api/weiYan/deleteWeiYan/", DeleteWeiYanView.as_view(), name="updateLabel"),
    path("api/weiYan/saveWeiYan/", AddWeiYanView.as_view(), name="updateLabel"),
    path("api/webInfo/listResourcePath/", ListResourcePathView.as_view(), name="updateLabel"),
    path("api/webInfo/updateResourcePath/", UpdateResourcePathView.as_view(), name="updateLabel"),
    path("api/webInfo/saveResourcePath/", SaveResourcePathView.as_view(), name="updateLabel"),
    path("api/webInfo/deleteResourcePath/", DeleteResourcePathView.as_view(), name="updateLabel"),
    path("api/webInfo/getClassifyList/", ListLoveResourcePathView.as_view(), name="updateLabel"),
    path("api/webInfo/clistResourcePath/", ListPrResourcePathView.as_view(), name="updateLabel"),
    path("api/ip/", GetIpView.as_view(), name="updateLabel"),
    path("api/list/ip/", ListIpView.as_view(), name="updateLabel"),
    path("api/submit/", RcView.as_view(), name="updateLabel"),
    path("api/map", MapView.as_view(), name="updateLabel"),
    path("api/user", UserView.as_view(), name="updateLabel"),
    path("api/webInfo/getWebInfo", GWebInfoPreView.as_view(), name="updateLabel"),
    path("api/article/listArticle", ListPreView.as_view(), name="updateLabel"),
    path("api/article/getArticleById/", GetArtPreView.as_view(), name="updateLabel"),
    path("api/weiYan/listNews/", WeiYanPreView.as_view(), name="updateLabel"),
    path("api/weiYan/saveNews/", AddWeiYanPreView.as_view(), name="updateLabel"),
    path("api/weiYan/changepublic/", ChangeWeiYanView.as_view(), name="updateLabel"),
    path("api/webInfo/allWebBackgroundImage/", ListResourcePreBGView.as_view(), name="updateLabel"),
    path("api/code/", CodeView.as_view(), name="updateLabel"),
    path("api/webInfo/listCollect/", ListCollectView.as_view(), name="updateLabel"),
    path("api/article/articleLike/", LikeView.as_view(), name="updateLabel"),
    path("api/user/updateUserInfo/", UpdateUserView.as_view(), name="updateLabel"),
]

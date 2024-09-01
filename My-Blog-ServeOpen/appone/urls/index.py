from django.urls import path

from appone.views.login import LoginView
from appone.views.registration import RegisterView
from appone.views.regret import RegretView
from appone.views.test import testView

urlpatterns = [
    path("api/appone/login/", LoginView.as_view(), name="login"),
    path("api/appone/registration/", RegisterView.as_view(), name='register'),
    path("api/appone/test/", testView.as_view(), name='test'),
    path("api/user/updateForForgetPassword/", RegretView.as_view(), name='updateForForgetPassword'),
]

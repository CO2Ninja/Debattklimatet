from django.conf.urls import patterns, include, url
from django.views.generic import ListView, TemplateView
from debattklimatet.models import twitteruser, tweet
from debattklimatet.views import TwitteruserDetail
from django.contrib import admin
admin.autodiscover()

urlpatterns = patterns('',
    # Examples:
    # url(r'^$', 'django_backend.views.home', name='home'),
    # url(r'^blog/', include('blog.urls')),

    url(r'^admin/', include(admin.site.urls)),
    (r'^$', TemplateView.as_view(template_name='debattklimatet/index.html')),
    (r'^tweets/$', ListView.as_view(model=tweet,)),
    (r'^canvas/$', TwitteruserDetail.as_view()),
    
)

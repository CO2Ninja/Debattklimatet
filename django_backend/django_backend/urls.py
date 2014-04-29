from django.conf.urls import patterns, include, url
from django.views.generic import ListView
from debattklimatet.models import twitteruser, tweet
from django.contrib import admin
admin.autodiscover()

urlpatterns = patterns('',
    # Examples:
    # url(r'^$', 'django_backend.views.home', name='home'),
    # url(r'^blog/', include('blog.urls')),

    url(r'^admin/', include(admin.site.urls)),
    (r'^$', ListView.as_view(model=twitteruser,)),
    (r'^tweets/$', ListView.as_view(model=tweet,)),
)

from django.contrib import admin
from models import twitteruser, tweet, hashtag, media

admin.site.register(twitteruser)
admin.site.register(tweet)
admin.site.register(hashtag)
admin.site.register(media)

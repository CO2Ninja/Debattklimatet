# -*- coding: utf-8 -*-

from django.db import models

#Creating table debattklimatet_twitteruser
#Creating table debattklimatet_hashtag
#Creating table debattklimatet_media
#Creating table debattklimatet_tweet_hashtags
#Creating table debattklimatet_tweet_media
#Creating table debattklimatet_tweet


#Totalpoäng
#Utträkning av procent?
#Hurvida tweets innehåller miljökommentarer eller ej
#Om tweets har gåtts igenom av parsern.

class twitteruser(models.Model):
    id                  = models.BigIntegerField(primary_key = True)
    name                = models.CharField(max_length = 200)
    screenname          = models.CharField(max_length = 200)
    profileimageurl     = models.CharField(max_length = 200)
    rating              = models.IntegerField()
    totalscore          = models.IntegerField()

    def __unicode__(self):
        return unicode(self.screenname)
    
class hashtag(models.Model):
    tag                 = models.CharField(max_length = 100, primary_key = True)


class media(models.Model):
    id                  = models.BigIntegerField(primary_key = True)
    media_url           = models.CharField(max_length = 200)
    media_url_https     = models.CharField(max_length = 200)
    url                 = models.CharField(max_length = 200)

class tweet(models.Model):
    createdat           = models.DateTimeField(db_index = True)
    favoritecount       = models.IntegerField()
    favorited           = models.BooleanField()
    id                  = models.BigIntegerField(primary_key = True)
    idstr               = models.CharField(max_length = 200)
    #InReplyToScreenName = models.CharField(max_length = 200)
    #InReplyToStatusID   = models.BigIntegerField()
    #InReplyToStatusIdStr= models.CharField(max_length = 200)
    #InReplyToUserID     = models.BigInteger(Field)
    retweetcount        = models.IntegerField()
    retweeted           = models.BooleanField()
    source              = models.CharField(max_length = 200)
    text                = models.CharField(max_length = 150)
    user                = models.ForeignKey(twitteruser)
    parsed              = models.BooleanField()
    relevant            = models.BooleanField()
    media               = models.ForeignKey(media, null = True)
    hashtags            = models.ManyToManyField(hashtag, null = True)
    
    def __unicode__(self):
        return unicode(self.text)
    
    

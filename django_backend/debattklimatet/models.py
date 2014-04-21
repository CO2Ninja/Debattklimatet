from django.db import models

#Creating table debattklimatet_user
#Creating table debattklimatet_hashtag
#Creating table debattklimatet_media
#Creating table debattklimatet_tweet_HashTags
#Creating table debattklimatet_tweet_Media
#Creating table debattklimatet_tweet


class User(models.Model):
    Id                  = models.BigIntegerField(primary_key = True)
    Name                = models.CharField(max_length = 200)
    ScreenName          = models.CharField(max_length = 200)
    ProfileImageUrl     = models.CharField(max_length = 200)
    
class HashTag(models.Model):
    tag                 = models.CharField(max_length = 100, primary_key = True)


class Media(models.Model):
    Id                  = models.BigIntegerField(primary_key = True)
    Media_url           = models.CharField(max_length = 200)
    Media_url_https     = models.CharField(max_length = 200)
    Url                 = models.CharField(max_length = 200)

class Tweet(models.Model):
    CreatedAt           = models.DateTimeField(db_index = True)
    FavoriteCount       = models.IntegerField()
    Favorited           = models.BooleanField()
    Id                  = models.BigIntegerField(primary_key = True)
    IdStr               = models.CharField(max_length = 200)
    #InReplyToScreenName = models.CharField(max_length = 200)
    #InReplyToStatusID   = models.BigIntegerField()
    #InReplyToStatusIdStr= models.CharField(max_length = 200)
    #InReplyToUserID     = models.BigInteger(Field)
    RetweetCount        = models.IntegerField()
    Retweeted           = models.BooleanField()
    Source              = models.CharField(max_length = 200)
    Text                = models.CharField(max_length = 150)
    User                = models.ForeignKey(User)
    HashTags            = models.ManyToManyField(HashTag)
    Media               = models.ManyToManyField(Media)
    
    
    

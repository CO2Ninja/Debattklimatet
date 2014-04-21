from django.db import models



#type Tweet struct {
    #Contributors         []int64     `json:"contributors"`
    #Coordinates          interface{} `json:"coordinates"`
    #CreatedAt            Char      `json:"created_at"`
    #Entities             Entities    `json:"entities"`
    #FavoriteCount        int         `json:"favorite_count"`
    #Favorited            bool        `json:"favorited"`
    #Geo                  interface{} `json:"geo"`
    #Id                   int64       `json:"id"`
    #IdStr                Char      `json:"id_str"`
    #InReplyToScreenName  Char      `json:"in_reply_to_screen_name"`
    #InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
    #InReplyToStatusIdStr Char      `json:"in_reply_to_status_id_str"`
    #InReplyToUserID      int64       `json:"in_reply_to_user_id"`
    #InReplyToUserIdStr   Char      `json:"in_reply_to_user_id_str"`
    #Place                Place       `json:"place"`
    #PossiblySensitive    bool        `json:"possibly_sensitive"`
    #RetweetCount         int         `json:"retweet_count"`
    #Retweeted            bool        `json:"retweeted"`
    #RetweetedStatus      *Tweet      `json:"retweeted_status"`
    #Source               Char      `json:"source"`
    #Text                 Char      `json:"text"`
    #Truncated            bool        `json:"truncated"`
    #User                 User        `json:"user"`
#}
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
    RetweetedStatus     = models.ForeignKeyField(Tweet)
    Source              = models.CharField(max_length = 200)
    Text                = models.CharField(max_length = 150)
    User                = models.ForeignKey(User)
    HashTags            = models.ManyToManyField(HashTag)
    Media               = models.ManyToManyField(Media)
    
#type User struct {
    #ContributorsEnabled            bool   `json:"contributors_enabled"`
    #CreatedAt                      Char `json:"created_at"`
    #DefaultProfile                 bool   `json:"default_profile"`
    #DefaultProfileImage            bool   `json:"default_profile_image"`
    #Description                    Char `json:"description"`
    #FavouritesCount                int    `json:"favourites_count"`
    #FollowRequestSent              bool   `json:"follow_request_sent"`
    #FollowersCount                 int    `json:"followers_count"`
    #Following                      bool   `json:"following"`
    #FriendsCount                   int    `json:"friends_count"`
    #GeoEnabled                     bool   `json:"geo_enabled"`
    #Id                             int64  `json:"id"`
    #IdStr                          Char `json:"id_str"`
    #IsTranslator                   bool   `json:"is_translator"`
    #Lang                           Char `json:"lang"`
    #ListedCount                    int64  `json:"listed_count"`
    #Location                       Char `json:"location"`
    #Name                           Char `json:"name"`
    #Notifications                  bool   `json:"notifications"`
    #ProfileBackgroundColor         Char `json:"profile_background_color"`
    #ProfileBackgroundImageURL      Char `json:"profile_background_image_url"`
    #ProfileBackgroundImageUrlHttps Char `json:"profile_background_image_url_https"`
    #ProfileBackgroundTile          bool   `json:"profile_background_tile"`
    #ProfileImageURL                Char `json:"profile_image_url"`
    #ProfileImageUrlHttps           Char `json:"profile_image_url_https"`
    #ProfileLinkColor               Char `json:"profile_link_color"`
    #ProfileSidebarBorderColor      Char `json:"profile_sidebar_border_color"`
    #ProfileSidebarFillColor        Char `json:"profile_sidebar_fill_color"`
    #ProfileTextColor               Char `json:"profile_text_color"`
    #ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
    #Protected                      bool   `json:"protected"`
    #ScreenName                     Char `json:"screen_name"`
    #ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
    #Status                         *Tweet `json:"status"` // Only included if the user is a friend
    #StatusesCount                  int64  `json:"statuses_count"`
    #TimeZone                       Char `json:"time_zone"`
    #URL                            Char `json:"url"`
    #UtcOffset                      int    `json:"utc_offset"`
    #Verified                       bool   `json:"verified"`
#}
class User(models.Model):
    Id                  = models.BigIntegerField(primary_key = True)
    Name                = models.CharField(max_length = 200)
    ScreenName          = models.CharField(max_length = 200)
    ProfileImageUrl     = models.CharField(max_length = 200)
    
class HashTag(models.Model):
    tag                 = models.CharField(max_length = 100, primary_key = True)


#Media []struct {
        #Id              int64
        #Id_str          Char
        #Media_url       Char
        #Media_url_https Char
        #Url             Char
        #Display_url     Char
        #Expanded_url    Char
        #Sizes           MediaSizes
        #Type            Char
        #Indices         []int
    #}
class Media(models.Model):
    Id                  = models.BigIntegerField(primary_key = True)
    Media_url           = models.CharField(max_length = 200)
    Media_url_https     = models.CharField(max_length = 200)
    Url                 = models.CharField(max_length = 200)
    
    

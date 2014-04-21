from django.db import models



#type Tweet struct {
    #Contributors         []int64     `json:"contributors"`
    #Coordinates          interface{} `json:"coordinates"`
    #CreatedAt            string      `json:"created_at"`
    #Entities             Entities    `json:"entities"`
    #FavoriteCount        int         `json:"favorite_count"`
    #Favorited            bool        `json:"favorited"`
    #Geo                  interface{} `json:"geo"`
    #Id                   int64       `json:"id"`
    #IdStr                string      `json:"id_str"`
    #InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
    #InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
    #InReplyToStatusIdStr string      `json:"in_reply_to_status_id_str"`
    #InReplyToUserID      int64       `json:"in_reply_to_user_id"`
    #InReplyToUserIdStr   string      `json:"in_reply_to_user_id_str"`
    #Place                Place       `json:"place"`
    #PossiblySensitive    bool        `json:"possibly_sensitive"`
    #RetweetCount         int         `json:"retweet_count"`
    #Retweeted            bool        `json:"retweeted"`
    #RetweetedStatus      *Tweet      `json:"retweeted_status"`
    #Source               string      `json:"source"`
    #Text                 string      `json:"text"`
    #Truncated            bool        `json:"truncated"`
    #User                 User        `json:"user"`
#}
class Tweet(models.Model):
    CreatedAt           = models.DateTimeField(db_index = True)
    FavoriteCount       = models.Integer()
    Favorited           = models.Boolean()
    Id                  = models.BigInteger(primary_key = True)
    IdStr               = models.String(max_length = 200)
    #InReplyToScreenName = models.String(max_length = 200)
    #InReplyToStatusID   = models.BigInteger()
    #InReplyToStatusIdStr= models.String(max_length = 200)
    #InReplyToUserID     = models.BigInteger()
    RetweetCount        = models.Integer()
    Retweeted           = models.Boolean()
    RetweetedStatus     = models.ForeignKey(Tweet)
    Source              = models.String(max_length = 200)
    Text                = models.String(max_length = 150)
    User                = models.ForeignKey(User)
    HashTags            = models.ManyToManyField(HashTag)
    Media               = models.ManyToManyField(Media)
    
#type User struct {
    #ContributorsEnabled            bool   `json:"contributors_enabled"`
    #CreatedAt                      string `json:"created_at"`
    #DefaultProfile                 bool   `json:"default_profile"`
    #DefaultProfileImage            bool   `json:"default_profile_image"`
    #Description                    string `json:"description"`
    #FavouritesCount                int    `json:"favourites_count"`
    #FollowRequestSent              bool   `json:"follow_request_sent"`
    #FollowersCount                 int    `json:"followers_count"`
    #Following                      bool   `json:"following"`
    #FriendsCount                   int    `json:"friends_count"`
    #GeoEnabled                     bool   `json:"geo_enabled"`
    #Id                             int64  `json:"id"`
    #IdStr                          string `json:"id_str"`
    #IsTranslator                   bool   `json:"is_translator"`
    #Lang                           string `json:"lang"`
    #ListedCount                    int64  `json:"listed_count"`
    #Location                       string `json:"location"`
    #Name                           string `json:"name"`
    #Notifications                  bool   `json:"notifications"`
    #ProfileBackgroundColor         string `json:"profile_background_color"`
    #ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
    #ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
    #ProfileBackgroundTile          bool   `json:"profile_background_tile"`
    #ProfileImageURL                string `json:"profile_image_url"`
    #ProfileImageUrlHttps           string `json:"profile_image_url_https"`
    #ProfileLinkColor               string `json:"profile_link_color"`
    #ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
    #ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
    #ProfileTextColor               string `json:"profile_text_color"`
    #ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
    #Protected                      bool   `json:"protected"`
    #ScreenName                     string `json:"screen_name"`
    #ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
    #Status                         *Tweet `json:"status"` // Only included if the user is a friend
    #StatusesCount                  int64  `json:"statuses_count"`
    #TimeZone                       string `json:"time_zone"`
    #URL                            string `json:"url"`
    #UtcOffset                      int    `json:"utc_offset"`
    #Verified                       bool   `json:"verified"`
#}
class User(models.Model):
    Id                  = models.BigInteger(primary_key = True)
    Name                = models.String(max_length = 200)
    ScreenName          = models.String(max_length = 200)
    ProfileImageUrl     = models.String(max_length = 200)
    
class HashTag(models.Model):
    tag                 = models.String(max_length = 100, primary_key = True)


#Media []struct {
        #Id              int64
        #Id_str          string
        #Media_url       string
        #Media_url_https string
        #Url             string
        #Display_url     string
        #Expanded_url    string
        #Sizes           MediaSizes
        #Type            string
        #Indices         []int
    #}
class Media(models.Model):
    Id                  = models.BigInteger(primary_key = True)
    Media_url           = models.String(max_length = 200)
    Media_url_https     = models.String(max_length = 200)
    Url                 = models.String(max_length = 200)
    
    

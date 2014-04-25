// Fetches tweets from Debattklimatet's Home Timeline
// and inserts the tweets into a database
// @CO2Ninja
// Version 0.1
package main

import (
	"database/sql"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	_ "github.com/lib/pq"
	"log"
	"net/url"
	"runtime"
)

//Twitter Id's
const (
	Moderaterna         = 19226961
	Vansterpartiet      = 17233550
	Centerpartiet       = 3796501
	Folkpartiet         = 18687011
	Miljopartiet        = 18124359
	Socialdemokraterna  = 3801501
	Kristdemokraterna   = 19014898
	Sverigedemokraterna = 97878686
)

//config variables
var CONSUMER_KEY = "xswE9V0Xjlsvzf14P7Mk7LOg5"
var CONSUMER_SECRET = "nHX2KIFZA4dFmUEOAIxA1msyOpydEtCyp13VREFcKjpVX8saHs"
var ACCESS_TOKEN = "2447607758-tsYHayIaChAAZ7JMlBcZ5SN86J0qXq9WqpO8xXP"
var ACCESS_TOKEN_SECRET = "KPCjLAQDhochBZm8Ggyw0c9U2V1Rv4LO7kiYDIvxmMWSj"
var dbURL = "user=co2ninjas dbname=co2ninjas password=co2ninjas12345 host=django-db.cyyapufsikx9.eu-west-1.rds.amazonaws.com port=5432"

var api *anaconda.TwitterApi

func init() {
	numcpu := runtime.NumCPU()
	fmt.Println("CPU count:", numcpu)
	runtime.GOMAXPROCS(numcpu)
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api = anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
}

//Fetches recent tweets from the Home timeline
//Count sets the ammount of tweets to retriev
func getHome(count string) []anaconda.Tweet {
	v := url.Values{}
	v.Set("count", count)
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}
	return tweets
}

/*
func dataStructurer(dataMap map[int]anaconda.Tweet) {

	for _ , tweet := range dataMap {
		//alla fält

		fmt.Println(tweet.Id)
		CreatedAt
		FavoriteCount //int64
		Favorited	//bool
		Retweeted	//bool
		RetweetCount	//int64
		Text	//string
		RetweetedStatus
		Source
		// Structs
		User
			    Id
    			Name
    			ScreenName
    			ProfileImageUrl


		Entities
			//array
			Hashtag
				 tag
			//array
			Media
				Id
    			Media_url
    			Media_url_https
    			Url
	}

}
*/

// Connects to a specified DB with specified paramters
func dbConnect(database string, parameters string) *sql.DB {
	db, err := sql.Open(database, parameters)
	if err != nil {
		log.Fatal(err)
	}
	return db

}

// Inserts the tweets into the right tables in the database
// table debattklimatet_user
// table debattklimatet_hashtag
// table debattklimatet_media
// table debattklimatet_tweet_HashTags
// table debattklimatet_tweet_Media
// table debattklimatet_tweet
func insertTweets(tweet []anaconda.Tweet) {
	db := dbConnect("postgres", dbURL)
	for _, tweets := range tweet {
		//debattklimatet_twitteruser  Parameters: Id int64, Name string, ScreenName string, ProfileImageUrl string
		_, error := db.Exec(
			"INSERT INTO debattklimatet_twitteruser(id, name, ScreenName, ProfileImageUrl, rating) VALUES(?, ?, ?, ?, ?)", 123, "tweets.User.Name", "tweets.User.ScreenName", "tweets.User.ProfileImageURL", 0)
		if error != nil {	
        	panic(error)
        	fmt.Println(tweets.Id)
    	}

	}

	/*
	result, err := db.Exec(
        "INSERT INTO users (name, age) VALUES (?, ?)",
        "gopher",
        27,
	)
	*/


	/*
	//debattklimatet_hashtag Parameters: Indices []int, Text   string
	_, error = db.Exec(`INSERT INTO debattklimatet_hashtag(id, name, ScreenName, ProfileImageUrl)
	VALUES(tweet.User.Id, tweet.User.Name, tweet.User.ScreenName, tweet.User.ProfileImageUrl,  ) RETURNING id`)

	//debattklimatet_media  Parameters: Id int64, Media_url string, Media_url_https string, Url string
	_, error = db.Exec(`INSERT INTO debattklimatet_user(id, name, ScreenName, ProfileImageUrl)
	VALUES(tweet.User.Id, tweet.User.Name, tweet.User.ScreenName, tweet.User.ProfileImageUrl ) RETURNING id`)

	//debattklimatet_tweet  Parameters: Id int64, CreatedAt string, FavoriteCount int64, Favorited bool, Retweeted bool, RetweetCount int64, Text string, Source string
	_, error = db.Exec(`INSERT INTO debattklimatet_user(id, name, ScreenName, ProfileImageUrl)
	VALUES(tweet.User.Id, tweet.User.Name, tweet.User.ScreenName, tweet.User.ProfileImageUrl ) RETURNING id`)
	*/

	db.Close()
}

//GO!
func main() {
	insertTweets(getHome("50"))
}

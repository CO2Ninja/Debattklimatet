// https://godoc.org/github.com/ChimeraCoder/anaconda
// Fetches tweets from Debattklimatet's Home Timeline
// And sorts them based on User Id
// @CO2Ninja
// Version 0.1
package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"runtime"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
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

//maps
var moderaterna = make(map[int]anaconda.Tweet)
var miljopartiet = make(map[int]anaconda.Tweet)
var vansterpartiet = make(map[int]anaconda.Tweet)
var socialdemokraterna = make(map[int]anaconda.Tweet)
var folkpartiet = make(map[int]anaconda.Tweet)
var kristdemokraterna = make(map[int]anaconda.Tweet)
var centerpartiet = make(map[int]anaconda.Tweet)
var sverigedmokraterna = make(map[int]anaconda.Tweet)

//config variables
var CONSUMER_KEY = "xswE9V0Xjlsvzf14P7Mk7LOg5"
var CONSUMER_SECRET = "nHX2KIFZA4dFmUEOAIxA1msyOpydEtCyp13VREFcKjpVX8saHs"
var ACCESS_TOKEN = "2447607758-tsYHayIaChAAZ7JMlBcZ5SN86J0qXq9WqpO8xXP"
var ACCESS_TOKEN_SECRET = "KPCjLAQDhochBZm8Ggyw0c9U2V1Rv4LO7kiYDIvxmMWSj"

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

// Sorts the tweet's
func sorter() {
	tweets := getHome("50")
	a, b, c, d, e, f, g, h := 0,0,0,0,0,0,0,0

	//sort tweets
	for _, tweet := range tweets {
		if tweet.User.Id == Moderaterna {
			moderaterna[a] = tweet
			a++
		}
		if tweet.User.Id == Miljopartiet {
			miljopartiet[b] = tweet
			b++
		}
		if tweet.User.Id == Vansterpartiet {
			vansterpartiet[c] = tweet
			c++
		}
		if tweet.User.Id == Socialdemokraterna {
			socialdemokraterna[d] = tweet
			d++
		}
		if tweet.User.Id == Folkpartiet {
			folkpartiet[e] = tweet
			e++
		}
		if tweet.User.Id == Kristdemokraterna {
			kristdemokraterna[f] = tweet
			f++
		}
		if tweet.User.Id == Centerpartiet {
			centerpartiet[g] = tweet
			g++
		}
		if tweet.User.Id == Sverigedemokraterna {
			sverigedmokraterna[h] = tweet
			h++
		}
	}
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

func main() {
	sorter()
	for i , tweets := range moderaterna {
		fmt.Println("index: ", i, tweets.Text)
	}
	
	db, err := sql.Open("postgres", "user=co2ninjas dbname=co2ninjas password=co2ninjas12345 host=django-db.cyyapufsikx9.eu-west-1.rds.amazonaws.com port=5432 sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}

	/*
	#Creating table debattklimatet_user
	#Creating table debattklimatet_hashtag
	#Creating table debattklimatet_media
	#Creating table debattklimatet_tweet_HashTags
	#Creating table debattklimatet_tweet_Media
	#Creating table debattklimatet_tweet
	*/
	fmt.Println("DB:")
	//id := 1234
	rows, err := db.Query("SELECT * FROM debattklimatet_user")


	fmt.Println(rows)

	//dataStructurer(moderaterna)


}



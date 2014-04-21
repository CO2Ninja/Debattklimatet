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

func main() {
	insertTweets(getHome("50"))

}
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
func insertTweets(tweets []anaconda.Tweet) {
	db := dbConnect("postgres", dbURL)

	rows, err := db.Query("SELECT * FROM debattklimatet_user")
	if err != nil {
            log.Fatal(err)
    }
    strings, _ := rows.Columns()

	for _, s := range strings {
		fmt.Println(s)
	}
	for rows.Next() {
	    var id int
	    var name string
	    err = rows.Scan(&id, &name)
	    fmt.Println(id, name)
	    fmt.Println(name)
	}


	/*
	var userid int
	error := db.QueryRow(`INSERT INTO debattklimatet_user(id, name, ScreenName, ProfileImageUrl)
	VALUES(123456, 'testsson', 'testtest', 'http://test.test' ) RETURNING id`).Scan(&userid)

	fmt.Println(error)
	*/

    db.Close()

}


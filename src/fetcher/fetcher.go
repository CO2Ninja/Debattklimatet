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

var partyArray []int64 = []int64{19226961, 17233550, 3796501, 18687011, 18124359, 3801501, 19014898, 97878686}

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

//Checks if the tweet.userId is a valid one
func chechParty(tweet anaconda.Tweet) bool {
	for _, id := range partyArray {
		if tweet.User.Id == id {
			return true
		}
	}
	return false
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

// Connects to a specified DB with specified paramters
func dbConnect(database string, parameters string) *sql.DB {
	db, err := sql.Open(database, parameters)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//checks if the User.Id exists
func userExists(db *sql.DB, id int64) bool {
	fmt.Println("test")
	rows, err := db.Query("SELECT COUNT(1) FROM debattklimatet_twitteruser WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
	return rows.Next()

}

// Inserts the tweets into the right tables in the database
func insertTweets(tweet []anaconda.Tweet) {
	db := dbConnect("postgres", dbURL)

	for _, tweets := range tweet {

		if !chechParty(tweets) {
			continue
		}
		//add user
		//if !userExists(db, tweets.User.Id) {
		if true {
			_, err := db.Exec(
				"INSERT INTO debattklimatet_twitteruser (id, name, screenname, profileimageurl, rating, totalscore) VALUES ($1, $2, $3, $4, $5, $6)",
				tweets.User.Id,
				tweets.User.Name,
				tweets.User.ScreenName,
				tweets.User.ProfileImageURL,
				0,
				0,
			)
			if err != nil {
				log.Println(err)
			}
		}

		//add tweet
		//createdat | favoritecount | favorited | id | idstr | retweetcount | retweeted | source | text | user_id
		_, err := db.Exec(
			"INSERT INTO debattklimatet_tweet (createdat, favoritecount, favorited, id, idstr, RetweetCount, Retweeted, Text, Source, user_id, parsed, relevant ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, (SELECT id FROM debattklimatet_twitteruser WHERE id=$10), $11, $12)",
			tweets.CreatedAt,
			tweets.FavoriteCount,
			tweets.Favorited,
			tweets.Id,
			tweets.IdStr,
			tweets.RetweetCount,
			tweets.Retweeted,
			tweets.Text,
			tweets.Source,
			tweets.User.Id,
			false,
			false,
		)
		if err != nil {
			log.Println(err)
		}

		//add tweet media
		//Id Media_url Media_url_https Url
		for _, i := range tweets.Entities.Media {
			fmt.Println(i.Id)

			_, err := db.Exec(
				"INSERT INTO debattklimatet_media (id, media_url, media_url_https, url) VALUES ($1, $2, $3, $4)",
				i.Id,
				i.Media_url,
				i.Media_url_https,
				i.Url,
			)
			if err != nil {
				log.Println(err)
			}
		}
	}

	db.Close()
}

//GO!
func main() {
	insertTweets(getHome("10"))
}

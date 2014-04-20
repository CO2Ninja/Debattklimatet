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
)

//Twitter Id's
const (
	Carl_Bildt          = 18549724
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

var api *anaconda.TwitterApi


func init() {
	numcpu := runtime.NumCPU()
	fmt.Println("CPU count:", numcpu)
	runtime.GOMAXPROCS(numcpu)
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api = anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
}

func Test_GetTweet() {
	const tweetId = 456755660504305664
	tweet, err := api.GetTweet(tweetId, nil)
	if err != nil {
		fmt.Errorf("GetTweet returned error: %s", err.Error())
	}
	fmt.Println(tweet.Text)
	fmt.Println(tweet.Id)
	fmt.Println(tweet.User.Id)
}

//Fetches tweets based on supplied query string
func testSearch() {
	search_result, err := api.GetSearch("miljö", nil)
	if err != nil {
		panic(err)
	}

	for _, tweet := range search_result {
		fmt.Print(tweet.Text)
	}
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

func parserTest() {

}

//Sorts tweets based on userId
//fix: add []userid
func sortTweets(id int64, tweets []anaconda.Tweet) {
	i := 0
	for _, tweet := range tweets {
		if tweet.User.Id == id {
			//moderat[i] = tweet   , add map or array
			i++
		}
	}
}

//Split in several functions and perhaps files/"classes"
func main() {
	i, o := 0, 0
	moderat := make(map[int]anaconda.Tweet)
	miljop := make(map[int]anaconda.Tweet)

	fmt.Println("")
	tweets := getHome("50")

	//sort tweets
	for _, tweet := range tweets {
		if tweet.User.Id == Moderaterna {
			moderat[i] = tweet
			i++
		}
		if tweet.User.Id == Miljopartiet {
			miljop[i] = tweet
			o++
		}
		//fmt.Print(tweet.User.Id, ": ", tweet.Text)
		//fmt.Println("")
	}


	//make embeded tweets(kör som separat gorutiner sen)
	for _, tweet := range moderat {
		//fmt.Println("Nya Moderaterna", ": ", tweet.Text)
		fmt.Println("")
		time, _ := tweet.CreatedAtTime()
		embeded, _ := api.GetOEmbedId(tweet.Id, nil)
		fmt.Println(embeded)
		fmt.Println(tweet.User.ProfileImageURL, " ", time, tweet.InReplyToScreenName)

	}
	for _, tweet := range miljop {
		fmt.Println("")
		time, _ := tweet.CreatedAtTime()
		embeded, _ := api.GetOEmbedId(tweet.Id, nil)
		fmt.Println(embeded)
		fmt.Println(tweet.User.ProfileImageURL, " ", time, tweet.InReplyToScreenName, tweet.Entities.Media)

	}
}

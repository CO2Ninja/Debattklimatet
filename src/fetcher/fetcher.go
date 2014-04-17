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

// Test_TwitterCredentials tests that non-empty Twitter credentials are set
// Without this, all following tests will fail
func Test_TwitterCredentials() {
	if CONSUMER_KEY == "" || CONSUMER_SECRET == "" || ACCESS_TOKEN == "" || ACCESS_TOKEN_SECRET == "" {
		fmt.Errorf("Credentials are invalid: at least one is empty")
	}
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

func testSearch() {
	search_result, err := api.GetSearch("miljö", nil)
	if err != nil {
		panic(err)
	}

	for _, tweet := range search_result {
		fmt.Print(tweet.Text)
	}
}

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

func main() {
	i := 0
	moderat := make(map[int]anaconda.Tweet)
	Test_TwitterCredentials()
	fmt.Println("")
	tweets := getHome("50")
	for _, tweet := range tweets {
		if tweet.User.Id == Moderaterna {
			moderat[i] = tweet
			i++
		}
		//fmt.Print(tweet.User.Id, ": ", tweet.Text)
		//fmt.Println("")
	}
	//fmt.Println(moderat)
	for _, tweet := range moderat {
		//fmt.Println("Nya Moderaterna", ": ", tweet.Text)
		fmt.Println("")
		time, _ := tweet.CreatedAtTime()
		embeded, _ := api.GetOEmbedId(tweet.Id, nil)
		fmt.Println(embeded)
		fmt.Println(tweet.User.ProfileImageURL, " ", time)

	}
	//testSearch()
}

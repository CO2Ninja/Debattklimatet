// https://godoc.org/github.com/ChimeraCoder/anaconda
// Version 0.1
package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

var CONSUMER_KEY = "xswE9V0Xjlsvzf14P7Mk7LOg5"
var CONSUMER_SECRET = "nHX2KIFZA4dFmUEOAIxA1msyOpydEtCyp13VREFcKjpVX8saHs"
var ACCESS_TOKEN = "2447607758-tsYHayIaChAAZ7JMlBcZ5SN86J0qXq9WqpO8xXP"
var ACCESS_TOKEN_SECRET = "KPCjLAQDhochBZm8Ggyw0c9U2V1Rv4LO7kiYDIvxmMWSj"

var api *anaconda.TwitterApi

func init() {
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
	search_result, err := api.GetSearch("a", nil)
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

func main() {

	Test_TwitterCredentials()
	fmt.Println("")
	tweets := getHome("10")
	for _, tweet := range tweets {
		fmt.Print(tweet.User.Id, ": ", tweet.Text)
		fmt.Println("")
	}
}

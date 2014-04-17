//https://dev.twitter.com/docs/auth/application-only-auth
//http://tweeterid.com/ convert from username to unique id
package main

import (
  "bitbucket.org/georgebaev/twitterfetcher"
  "fmt"
  "sync"
)

func main() {
  consumerKey := "xswE9V0Xjlsvzf14P7Mk7LOg5"
  consumerSecret := "nHX2KIFZA4dFmUEOAIxA1msyOpydEtCyp13VREFcKjpVX8saHs"

  tp := []twitterfetcher.Tweet{}

  token, err := twitterfetcher.GetBearerToken(consumerKey, consumerSecret)

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  urls := []string{
    "https://twitter.com/813286/",
    "https://twitter.com/3801501/",
  }
  wg := new(sync.WaitGroup)
  wg.Add(len(urls))
  for i, posturl := range urls {
    go func(posturl string, i int) {
      //err = twitterfetcher.HttpGet(token, "statuses/show", &tp)  // statuses/user_timeline borde ge en de senaste tweetsen
      err = twitterfetcher.HttpGetTweets(token, "statuses/user_timeline", &tp)

      if err != nil {
        fmt.Println(err.Error())
        return
      }

      fmt.Println(tp.Tweet.Id)

      for _, tweet := range tp {
        fmt.Print(tweet.User.Id, ": ", tweet.Text)
        fmt.Println("")
      }

      wg.Done()
    }(posturl, i)
  }
  wg.Wait()
  twitterfetcher.InvalidateBearerToken(token, consumerKey, consumerSecret)
}

func parser(s []string) (m map[string]string) {
  m = make(map[string]string)
  for _, url := range s {
    m["Obama"] = url
  }

  return 
}


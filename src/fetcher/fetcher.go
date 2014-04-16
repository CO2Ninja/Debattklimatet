//https://dev.twitter.com/docs/auth/application-only-auth

package main

import (
  "bitbucket.org/georgebaev/twitterfetcher"
  "fmt"
)

func main() {
  consumerKey := "xswE9V0Xjlsvzf14P7Mk7LOg5"
  consumerSecret := "nHX2KIFZA4dFmUEOAIxA1msyOpydEtCyp13VREFcKjpVX8saHs"

  tp := twitterfetcher.TwitterPost{}

  token, err := twitterfetcher.GetBearerToken(consumerKey, consumerSecret)

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  urls := []string{
    "https://twitter.com/BarackObama/status/456506067258597376",
  }

  for i, posturl := range urls {
    go func(posturl string, i int) {
      if err := tp.ValidateURL(posturl); err != nil {
        fmt.Println(err.Error())
        return
      }

      err = twitterfetcher.HttpGet(token, "statuses/show", &tp)

      if err != nil {
        fmt.Println(err.Error())
        return
      }

      fmt.Println(i, tp)
    }(posturl, i)
  }

  twitterfetcher.InvalidateBearerToken(token, consumerKey, consumerSecret)
}
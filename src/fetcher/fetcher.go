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

  tp := twitterfetcher.TwitterPost{}

  token, err := twitterfetcher.GetBearerToken(consumerKey, consumerSecret)

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  urls := []string{
    "https://twitter.com/813286/status/456506067258597376",
    "https://twitter.com/BarackObama/status/456529033300168706",
  }
  m := parser(urls)
  wg := new(sync.WaitGroup)
  wg.Add(len(urls))
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
      wg.Done()
    }(posturl, i)
  }
  wg.Wait()
  twitterfetcher.InvalidateBearerToken(token, consumerKey, consumerSecret)
  fmt.Print(m)
}

func parser(s []string) (m map[string]string) {
  m = make(map[string]string)
  for _, url := range s {
    m["Obama"] = url
  }

  return 
}


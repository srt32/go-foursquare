package main

import (
  "fmt"
  fs "github.com/elbuo8/4square-venues"
  "log"
  "os"
)

func main() {
  fmt.Println("foo")
  fmt.Println(os.Getenv("FOURSQUARE_API_SECRET"))
  fmt.Println(os.Getenv("FOURSQUARE_API_ID"))
  client := fs.NewFSVenuesClient(os.Getenv("FOURSQUARE_API_ID"),
    os.Getenv("FOURSQUARE_API_SECRET"))
  params := make(map[string]string)
  params["ll"] = "34,-122"
  venues, err := client.GetVenue(params)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(venues)
}

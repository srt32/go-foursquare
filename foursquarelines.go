package main

import (
  "fmt"
  fs "github.com/elbuo8/4square-venues"
  "log"
  "os"
)

func main() {
  venues, err := getVenues()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(venues)
}

func getVenues() (interface{}, error) {
  client := fs.NewFSVenuesClient(os.Getenv("FOURSQUARE_API_ID"),
    os.Getenv("FOURSQUARE_API_SECRET"))
  params := make(map[string]string)
  params["ll"] = "37.769,-122.430"
  return client.GetVenues(params)
}

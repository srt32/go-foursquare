package main

import(
  "testing"
)

func Test_SearchVenues(t *testing.T){
  venues, err := getVenues()

  if err != nil {
    t.Error("Failed", err)
  } else {
    t.Log("Passed", venues)
  }
}

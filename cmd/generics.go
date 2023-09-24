package main 

import (
  "fmt"
  "encoding/json"
  "io/ioutil"
)

type contactInfo struct {
  Name string `json:"Name"`
  Email string `json:"Email"`
}

type purchaseInfo struct {
  Name string `json:"Name"`
  Price float32 `json:"Price"`
  Amount int `json:"Amount"`
}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T {
  data, _ := ioutil.ReadFile(filepath)
  var loaded = []T{}
  json.Unmarshal([]byte(data), &loaded)
  return loaded
}

func main() {
  var contacts []contactInfo = loadJSON[contactInfo]("./contacts.json")
  var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchases.json")
  fmt.Printf("%+v\n", contacts)
  fmt.Printf("%+v\n", purchases)
}

package main

import (
  "fmt"
  "math/rand"
  "time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func checkPrices(wesite string, c chan string, maxPrice float32) {
  for {
    time.Sleep(time.Second+1);
    var price = rand.Float32()*20
    if price <= maxPrice { 
      c <- wesite
      break
    }
  }
}

func sendMessage(chicken chan string, tofu chan string) {
  select {
    case wesite := <- chicken:
      fmt.Printf("Texted: Found a deal @ %v\n", wesite)
    case wesite := <- tofu:
      fmt.Printf("Emailed: Found a deal @ %v\n", wesite)
  }
}

func main() {
  then := time.Now()
  var chicken = make(chan string)
  var tofu = make(chan string)
  var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
  for i := range websites {
    go checkPrices(websites[i], chicken, MAX_CHICKEN_PRICE)
    go checkPrices(websites[i], tofu, MAX_TOFU_PRICE)
  }
  sendMessage(chicken, tofu)
  fmt.Println(time.Since(then))
}



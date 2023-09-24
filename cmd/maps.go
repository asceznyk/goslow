package main

import "fmt"

func main() {
  var myMap map[string]uint8 = make(map[string]uint8);
  fmt.Println(myMap);
  var myMap2 = map[string]uint8{"Adam":23, "Catpiss":22};
  fmt.Println(myMap2["Adam"]);
  for name, age := range myMap2{ fmt.Printf("Name %v, Age:%v\n", name, age); }
}

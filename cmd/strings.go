package main 

import (
  "fmt"
  "strings"
)

func main() {
  var myStr = []rune("resume");
  var indexed = myStr[1];
  fmt.Printf("%v, %T\n", indexed, indexed);
  for i, v := range myStr { fmt.Println(i,v); }
  var strSlice = []string{"s","u","b"};
  var strBuilder strings.Builder
  for i := range strSlice { strBuilder.WriteString(strSlice[i]) }
  var catStr = strBuilder.String();
  fmt.Println(catStr);
}

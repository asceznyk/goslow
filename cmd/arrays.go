package main

import "fmt"

func main() {
  intArr := [3]int32{1,2,3};
  fmt.Println(intArr[0], intArr[0:2]);
  var intSlice []int32 = []int32{4,5,6};
  intSlice = append(intSlice, 7);
  fmt.Println(intSlice);
  var intSlice2 []int32 = []int32{8,9};
  fmt.Println(intSlice2);
  intSlice = append(intSlice, intSlice2...);
  fmt.Println(intSlice);
  var intSlice3 []int32 = make([]int32, 3, 8);
  fmt.Println(intSlice3);
  for i, v := range intArr{ fmt.Printf("Index: %v, Value: %v\n", i, v); }
}

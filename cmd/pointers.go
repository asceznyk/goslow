package main

import "fmt"

func square(arr *[5]int32) [5]int32 {
  fmt.Printf("arr addr = %p\n", arr)
  for i,v := range arr {
    arr[i] = v*v
  }
  return *arr
}

func main() {
  var p *int32 = new(int32)
  var i int32 = 30
  fmt.Printf("p addr = %v, p value = %v\n", p, *p)
  p = &i
  fmt.Printf("p addr = %v, i addr = %v\n", p, &i)
  fmt.Printf("p value = %v\n", *p)
  *p = 1
  fmt.Printf("i value = %v, p value = %v\n", i, *p);
  var slice = [5]int32{1,2,3,4,5}
  fmt.Printf("slice addr = %p\n", &slice)
  var sqslice [5]int32 = square(&slice)
  fmt.Println(sqslice, slice)
}

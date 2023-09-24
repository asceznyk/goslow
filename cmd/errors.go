package main

import (
  "errors"
  "fmt"
)

func intDiv(num int32, den int32) (int32, int32, error) {
  var err error;
  if den == 0 {
    err = errors.New("Cannot divide by Zero");
    return 0,0,err;
  }
  var q int32 = num/den;
  var r int32 = num%den;
  return q,r,err;
}

func main() {
  var q, r, err = intDiv(16, 0);
  if err != nil { fmt.Println(err); }
  fmt.Printf("quotient = %v, remainder = %v\n", q, r);
}

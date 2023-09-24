package main

import "fmt"

type gasEngine struct {
  mpg uint16
  gallons uint16
  owner
}

type electricEngine struct {
  mpkwh uint16
  kwh uint16
  owner
}

type owner struct {
  name string
}

func (e gasEngine) milesLeft() uint16 { return e.gallons*e.mpg }

func (e electricEngine) milesLeft() uint16 { return e.kwh*e.mpkwh }

type engine interface {
  milesLeft() uint16
}

func canMakeIt(e engine, miles uint16) {
  if miles <= e.milesLeft() { 
    fmt.Println("You can make it bitch!")
  } else {
    fmt.Println("Fuck no!")
  }
}

func main() {
  var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
  var otherEngine electricEngine = electricEngine{34, 10, owner{"Jennifer"}}
  canMakeIt(myEngine, 500)
  canMakeIt(otherEngine, 20)
}

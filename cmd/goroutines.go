package main

import (
  "fmt"
  "math/rand"
  "time"
  "sync"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData  = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func save(result string) {
  m.Lock()
  results = append(results, result)
  m.Unlock()
}

func log() {
  m.RLock()
  fmt.Printf("Result @ time %v\n", results);
  m.RUnlock()
}

func dbCall(i int) {
  var delay float32 = rand.Float32()*2000
  time.Sleep(time.Duration(delay)*time.Millisecond)
  fmt.Printf("dbCall Result = %v\n", dbData[i])
  save(dbData[i])
  log()
  wg.Done()
}

func main() {
  then := time.Now()
  for i:=0; i<len(dbData); i++ {
    wg.Add(1)
    go dbCall(i)
  }
  wg.Wait()
  fmt.Printf("Total execution time %v\n", time.Since(then))
  fmt.Printf("Results = %v\n", results);
}

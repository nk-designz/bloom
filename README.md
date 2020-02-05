# Bloom
a simple library for implementing bloom-filters into go-projects
# Example
```go
package main

import(
  "fmt"
  "os"

  "github.com/nk-designz/bloom"
)

func main() {
                             // Bitsize, Hashfunctions 
  vinylCollection := bloom.Init(4096, bloom.Murmur, bloom.Blake3, bloom.Spooky)
  for {
    var cmd, record string
    fmt.Print("What to do? ")
    fmt.Scan(&cmd)
    fmt.Scan(&record)
    if(cmd == "add") {
      vinylCollection.Add(record)
    } else if(cmd == "probe") {
      if(vinylCollection.Probe(record)) {
        fmt.Println("maybe have it")
      } else {
        fmt.Println("no.")
      } 
    } else if(cmd == "exit") {
        os.Exit(0)
    } else {
      fmt.Println("What?.")
    }
  }
}
```
```text
$ go run . 
What to do? add 2112_Rush
What to do? probe Sgt_Peppers-The_Beatles
no.
What to do? add Sgt_Peppers-The_Beatles
What to do? probe Sgt_Peppers-The_Beatles
maybe have it
```

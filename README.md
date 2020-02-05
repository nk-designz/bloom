# Bloom
a simple library for implementing bloom-filters into go-projects
# Usage
```go
package main

import(
  "bufio"
  "fmt"
  "os"
 
  "https://github.com/nk-designz/bloom"
)

func main() {
                             // Bitsize, Hashfunctions 
  vinylCollection := bloom.Init(4096, bloom.Murmur, bloom.Blake3, Bloom.Spooky)
  reader := bufio.NewReader(os.Stdin)
  for {
    cmd, _ := reader.ReadString('\n')
    record, _ := reader.ReadString('\n')
    if(cmd == "add") {
      vinylCollection.Add(record)
    } else {
      if(vinylCollection.Probe(record)) {
        fmt.Println("chance of success")
      } else {
        fmt.Println("no.")
      }
    }
  }
}
```

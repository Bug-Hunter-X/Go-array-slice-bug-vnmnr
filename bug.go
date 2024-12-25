```go
package main

import "fmt"

func main() {


    var a [5]int
    fmt.Println(a)
    fmt.Println(len(a))
    b := a[:2] // slice of a
    fmt.Println(b)
    fmt.Println(len(b))
    b[0] = 100
    fmt.Println(a)

}
```
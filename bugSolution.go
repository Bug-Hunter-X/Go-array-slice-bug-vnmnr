```go
package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println(a)
    fmt.Println(len(a))
    b := make([]int, len(a[:2]))
    copy(b, a[:2]) // Create a copy of the slice
    fmt.Println(b)
    fmt.Println(len(b))
    b[0] = 100
    fmt.Println(a)

}
```
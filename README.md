# runningvariance
A package for accurately computing running (online) mean, variance, and standard
deviation. Based on [code by John D
Cook](http://www.johndcook.com/blog/standard_deviation/).

`go get github.com/carbocation/runningvariance`

To test (100% coverage)

`go test -cover`

Demo:

```go
import "fmt"

import "github.com/carbocation/runningvariance"

func main() {
    var rv runningvariance.Stat

    rv.Push(1)
    rv.Push(2)
    rv.Push(3)
    rv.Push(4)
    rv.Push(5)

    fmt.Printf("Mean: %.4f | SD %.4f", rv.Mean(), rv.StandardDeviation())
}
```
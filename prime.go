package main
 
import (
    "fmt"
    "math"
    "time"
)

var n int
var prevTime time.Time

func IsPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
        n+=1
        if n%100000000 == 0 {
            fmt.Printf("%2.3f Iteration %d, examining %d\n", time.Since(prevTime).Seconds(), n, value)
            prevTime = time.Now()
        }
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func main() {
    prevTime = time.Now()
    for i:=1; i<10000000; i++ {
        IsPrime(i)
    }
}

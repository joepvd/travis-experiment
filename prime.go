package main
 
import (
    "fmt"
    "math"
    "time"
    "os"
    "runtime"
)

var n int
var prevTime time.Time
var startTime time.Time
var c1 chan string
var c2 chan string

func IsPrime(value int, identifyer string, channel chan string) bool {
    for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
        n+=1
        if n%100000000 == 0 {
            elapsedSinceStart := time.Since(startTime).Seconds()
            elapsedSincePrev := time.Since(prevTime).Seconds()
            channel <- fmt.Sprintf("%2.3f %2.3f %s Iteration %d, examining %d", elapsedSinceStart, elapsedSincePrev, identifyer, n, value)
            prevTime = time.Now()
        }
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func runTests(limit int, identifyer string, channel chan string) {
    prevTime = time.Now()
    startTime = time.Now()
    for i:=1; i<limit; i++ {
        IsPrime(i, identifyer, channel)
    }
}

func main() {
    runtime.GOMAXPROCS(4)
    c1 := make(chan string)
    c2 := make(chan string)
    go runTests(100000, "stuffs", c1)
    go runTests(100000, "niceness", c2)
    for {
        select {
        case channel1 := <- c1:
            fmt.Println(channel1)
        case channel2 := <- c2:
            fmt.Println(channel2)
        case <- time.After(time.Second * 10):
            fmt.Println("Time is up")
            os.Exit(0)
        }
    }
}

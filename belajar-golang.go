package	main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())
    var randomValue int

    randomValue = randomWithRange(100, 300)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(300, 100)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(200, 300)
    fmt.Println("random number:", randomValue)
}

func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}

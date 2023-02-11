package main
import (
        "os"
        "time"
        "strings"
)

func printlogs() {
    time.Sleep(50 * time.Millisecond)
    balloon := 
    `
    . . . . . . . . . . . .
    . . . . . . . . . . . .
    . . . ._. === ._. . . .
    . . .'***'.|**\..'. . .
    . ./***/...|***\..\\. .
    . |*****...|****...|| .
    . :****|...|***|...;; .
    . .\****...|***...//. .
    . . \***\..|**/..// . .
    . . . \**\.|*/..' . . .
    . . . . '*\|/.' . . . .
    . . . . . /x\ . . . . .
    . . . . .((_)). . . . .
    . . . . . ''' . . . . .                 
    `

    for _, line := range strings.Split(strings.TrimRight(balloon, "\n"), "\n") {
        time.Sleep(70 * time.Millisecond)
        println(line)
     }
    time.Sleep(100 * time.Millisecond)
    println("   ==============================")
    time.Sleep(100 * time.Millisecond)
    println("   Hello from", os.Getenv("FLY_REGION"))
    time.Sleep(100 * time.Millisecond)
    println("   I'm a Fly Machine in app", os.Getenv("FLY_APP_NAME"))
    time.Sleep(100 * time.Millisecond)
    println("   Goodbye! Stopping Machine")
    time.Sleep(100 * time.Millisecond)
    println("   ==============================")
}

func main() {
    defer time.Sleep(100 * time.Millisecond)
    printlogs()
}
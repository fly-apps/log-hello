package main
import (
        "os"
        "time"
)

func printlogs() {
    time.Sleep(100 * time.Millisecond)
    println("=========================")
    println("I'm a Fly Machine in app", os.Getenv("FLY_APP_NAME"))
    println("Hello from", os.Getenv("FLY_REGION"))
    println("Goodbye! Stopping Machine")
    println("=========================")
}

func main() {
    defer time.Sleep(100 * time.Millisecond)
    printlogs()
}
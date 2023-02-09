package main
import (
        "os"
)
func main() {
    println("======================")
    println("")
    println("Hello from", os.Getenv("FLY_REGION"))
    println("Goodbye! Stopping Machine")
    println("")
    println("======================")
	os.Exit(0)
}
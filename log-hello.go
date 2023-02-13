package main
import (
        "os"
        "time"
        "strings"
)

func printlogs() {
    time.Sleep(50 * time.Millisecond)
    balloon := 
    ` . . . . . . . . . . . . . . . .
 . /\/\. . . . . . . . . . . . .
 . . . . . . . . . . .^^ . . . .
 . . . . . _. === ._ . . . . . .
 . . . . ,'*;'.|**\.'\ . . . . .
 . . . ./**/...|***\.\\. . . . .
 . . . .|**|...|***|.||. . . . .
 . . . . \**\..|**/.// . . . . .
 ^^. . . .'**..|*/./'. . . . . .
 . . . . . .'*\|/.'. . . . . . .
 . . . . . . . X . . . . . . . .
 . . . . . . .(O). . . . . . . .
 . . . . . . . . . . . /¯\v/¯\ .
 . . . . . . . . . . . . . . . .`

    println("3...")
    time.Sleep(time.Second)
    println("2...")
    time.Sleep(time.Second)
    println("1...")
    time.Sleep(time.Second)
    for _, line := range strings.Split(strings.TrimRight(balloon, "\n"), "\n") {
        println(line)
        time.Sleep(50 * time.Millisecond)
     }
    println(" ======= Hello from", os.Getenv("FLY_REGION"), "========")
    time.Sleep(50 * time.Millisecond) 
    println(" I'm a Fly Machine in app", os.Getenv("FLY_APP_NAME"))
    time.Sleep(50 * time.Millisecond)
    println(" I'm about to shut down!")
    time.Sleep(50 * time.Millisecond)
    println(" Run this Machine again:")
    time.Sleep(50 * time.Millisecond)
    println(" ---")
    time.Sleep(50 * time.Millisecond)
    println(" fly m start", os.Getenv("FLY_ALLOC_ID"))
    time.Sleep(50 * time.Millisecond)
    println(" ---")
    time.Sleep(50 * time.Millisecond)
    println(" Goodbye")
    time.Sleep(50 * time.Millisecond)
    println(" ===============================")
}

func main() {
    // defer time.Sleep(100 * time.Millisecond)
    printlogs()
    
}
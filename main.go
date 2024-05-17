package main

import (
    "log"
    "os"
)

func main() {
    // Create a log file
    file, err := os.OpenFile("example.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Set the output of the logger to the file
    log.SetOutput(file)

    // Set log prefix and flags
    log.SetPrefix("LOG: ")
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    // Log messages
    log.Println("This is a test log message")
}

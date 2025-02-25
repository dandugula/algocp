package main
import "fmt"

func deferArg() bool {
    fmt.Println("Defer argument is evaluated now..")
    return true;
}

func deferFunc(v bool) {
    fmt.Printf("Defer function: %v\n", v)
}

func deferExp() {
    for range(2) {
        defer deferFunc(deferArg())
    }
    fmt.Println("Within deferExp func")
}

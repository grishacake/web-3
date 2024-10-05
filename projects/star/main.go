package main
import (
    "fmt"
    "strings"
)

func main() {
    var str string
    fmt.Scan(&str)
    fmt.Print(strings.Join(strings.Split(str, ""), "*"))
}
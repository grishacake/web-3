package main
import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    for i := 0; i < len(str); i++ {
        digit := int(str[i]) - '0'
        fmt.Print(digit*digit)
    }
}
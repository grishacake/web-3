package main
import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    var max int
    for i := 0; i < len(str); i++ {
        digit := int(str[i]) - '0'
        if digit > max {
            max = digit
        }
    }
    fmt.Print(max)
}
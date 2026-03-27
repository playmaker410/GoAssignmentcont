package main
import "fmt"

func main (){
var ch string
fmt.Println("Enter a Character:")

fmt.Scan(&ch)

switch ch {

case "a", "e", "i", "o","u":
fmt.Println("Vowel")

default:
fmt.Println("Consonant")
}
}

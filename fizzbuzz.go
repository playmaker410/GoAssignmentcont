
package main
import "fmt"

func main(){

for i:=1;i<=50;i++{

if i % 3 == 0 && i % 5 == 0{
fmt.Println("FizzBuzz")

}else if   i % 3 == 0 {
fmt.Println("fizz")
} else if  i % 5 == 0 {
fmt.Println("buzz")
} else{
fmt.Println(i)
}
}



//SWITCH REPRESSENTATION OF IT

for i := 1; i <= 50;i++{

switch {

case i %  3== 0 && i % 5 == 0:
fmt.Println("fizzBuzz")
 case i % 3 == 0:
fmt.Println("fizz")

case i % 5 == 0:
fmt.Println("buzz")

default :
fmt.Println(i)

}

}

}

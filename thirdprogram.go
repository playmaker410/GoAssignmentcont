package main
import "fmt"


func main(){
x:= 10

for i :=0; i < 10; i++{
x:=  x+ 1
fmt.Println("value inside",x)
}

fmt.Println("value outside the loop",x)

}

//The value outside the loop did not change because the operator does not  eassigna value to the excistingouter variable
//instead it created a new variable and named it x an this only exist inside the loop

// variable shadowing occur when you declare a new variable with same name  as a variable in an outer scope.this can happen in a 
// loop or even in an if satement. one should be outside the statement and one inside the statement.

//To fix the code so that the outer variable is modified by the loop, you must use the standard assignment operator (=)
// instead of the short variable declaration operator (:=) inside the loop.




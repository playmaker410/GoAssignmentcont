package main
import "fmt"



func main(){
// int flow

var num int8 = 127
fmt.Println(num)
num = num + 1
fmt.Println(num)

//division

y := 10
x := 2
fmt.Println("int division",y/x)

//bitwise operator

p:= 5
z:= 3

fmt.Println("Using AND",p&z)// return 1 if both bit are 1 but if either  bit is 0 return 0
fmt.Println("usiin OR",p|z)// 
fmt.Println("using XOR",p^z) // following the rule if it has same value print 0 and if it is  different print 1. note it is being 
//converted into variable
fmt.Println("using Zero fill left shift",  p<<1)
fmt.Println("using Zero fill right shift",p>>1)


//modules
m := -5
o := -3

fmt.Println("modules",m%o);


}

package main 
import "fmt"

func main(){
const MaxUser = 5

Username := []string{}

NewUsers :=[]string{"golang","Go", "", "ruben","Luckify", "joshua","playmaker","Oshimodinakaka"}

for i:=0; i < len(NewUsers);i++{
name:=NewUsers[i]

if name == ""{
fmt.Println("You cannot live this box unfiled")
continue
}
if len(name) <= 3{
fmt.Println("this is too short must be at leat 5 characters",name)
continue
}
if len(Username) == MaxUser{
fmt.Println("Cant exceed registration limit")
break
}

Username=append(Username,name)
fmt.Println("User added",name)

}
}

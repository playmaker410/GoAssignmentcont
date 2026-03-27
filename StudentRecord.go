package main
import (
	"fmt"
	)



type Student struct{
name string
age int
score int
}

func  Grades(score int) string {
if score >= 70{
return "A"
} else if score >= 60{
return "B"
}else if score >= 50{
return "C"
} else{
return "F"
}
}

func Evendeterminer(score int) string{
if score % 2 == 0{
return "Even"
}
return "Odd"
}

func highestScore(students []Student)Student{
top := students[0]

for _, s := range students{
if s.score > top.score{
top = s
}
}
return top

}


func averageScore(students []Student)float64{
total := 0
for _, s := range students{
total += s.score 
}
return float64(len(students)) /
float64(len(students))

}



func main(){
students := []Student{
{"JOHN", 18,75},
{"PETER", 17,62},
{"DIVINE", 19,48},
{"DANIEL", 18,55},
{"PAUL", 20, 80},
}

for _, s := range students {
fmt.Println("Name:", s.name)
fmt.Println("Age",s.age)
fmt.Println("Score", s.age)
fmt.Println("Grade:", Grades(s.score))
fmt.Println("-----------------------")

}

top := highestScore(students)
fmt.Println("Top Score",top.name, "Score", top.score)
avg := averageScore(students)
fmt.Println("Average score", avg)


studentMap := make(map[string]Student)
for _, s := range students {
studentMap[s.name]= s
}

name := "mary"
if student, found := studentMap[name];
found {
fmt.Println("Found:", student)
} else{
fmt.Println("student not found")
}
}


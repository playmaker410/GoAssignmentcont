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
return float64(total) /
float64(len(students))

}



func main(){
students := []Student{
{"JOHN", 18,18},
{"PETER", 10,80},
{"DIVINE", 79,21},
{"DANIEL", 68,31},
{"PAUL", 17, 50},
}

for _, s := range students {
fmt.Println("Name:", s.name)
fmt.Println("Age",s.age)
fmt.Println("Score", s.score)
fmt.Println("Grade:", Grades(s.score))
fmt.Println("Even/odd",  Evendeterminer(s.score))
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

name := "PAUL"
if student, found := studentMap[name];
found {
fmt.Println("Found:", student)
} else{
fmt.Println("student not found")
}
}


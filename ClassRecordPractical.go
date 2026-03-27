package main
import "fmt"




type Student struct{
name string
age int
score int

}

func Grading(score int)string{
if score >= 70{
return "A"
} else if score >= 60 {
return "B"

}else if score >= 50 {
return "C"
} else {
return "F"
}
}

//even or odd function

func evenchecker(score int) string{
if score % 2 == 0{
return "Even"
}else{
return "Odd"
}

}

//highest score

func topscore(Students []Student) Student {
highestScore :=Students[0]

for _, s := range Students{
if s.score > highestScore.score{
highestScore = s
}
}
return highestScore
}

//Average score
func AverageScore(Students []Student)int{
total := 0
for _, s := range Students{
total += s.score
}

return (total) /
  (len(Students))
}


func main(){

Students := []Student{
{"Jenifer", 12,60},
{"Victor", 14, 65},
{"Luckyfy", 20,72},
{"Charlse", 21, 50},
{"Elon", 22, 40},
}


for _, s := range Students{
fmt.Println("Name", s.name)
fmt.Println("Age", s.age)
fmt.Println("Score", s.score)
fmt.Println(Grading(s.score))
fmt.Println(evenchecker(s.score))
fmt.Println("============================")
}

highestScore := topscore(Students)
fmt.Println("Top name", highestScore.name, "Score", highestScore.score)
avg := AverageScore(Students)
fmt.Println("Average Score", avg)

ListofS := make(map[string]Student)

for _, s := range Students{
 ListofS[s.name] = s
}

name := "Jenifer"
if Student, found  := ListofS[name];
found{
fmt.Println("Found", Student)   
}else{
fmt.Println("Not our student")
}

}

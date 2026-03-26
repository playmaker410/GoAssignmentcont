package main
import "fmt"

type Student struct{
Name string
Age int
Score int
}


func main(){
var student1 Student
var  student2 Student
var student3 Student

//specifing my student

student1.Name = "Joshua OKonkwo"
student1.Age = 70
student1.Score =  100


//specifing his own

student2.Name = "kashmak Lucify"
student2.Age = 30
student2.Score = 80

//3rd one

student3.Name = "Angle Nnamani"
student3.Age = 50
student3.Score = 80

fmt.Println(highestScore(student1,student2,student3).Name)
fmt.Println(highestScore(student1,student2,student3).Score)
}

func  highestScore(student1 Student,student2 Student,student3 Student)Student{
maxScore := student1

if student2.Score > maxScore.Score{
maxScore = student2
}

if  student3.Score > maxScore.Score{
maxScore = student3
}

return maxScore
}

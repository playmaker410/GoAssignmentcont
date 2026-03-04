package main
import "fmt"

func main(){

intergers := []int {1,2,3,4,5}

for val := range intergers{
val = val * 8
}

fmt.Println("outside the loop",intergers)



//1) Why modifying range variable does not change the slice is because value is said to be a copy and am modifying the
//      copy not t he original

//2) How range works internally: it takes copy of a slice header which  contain a pointer to the underlying array,len and capacity.
//    In each iteration, it calculates the index (i) and makes a new copy of the value stored at
//    that index in the underlying array.

//3)  To correctly  modify slice value we just need to do that using the index to access the  slice


  nums := []int{1, 2, 3, 4, 5}

    // We only take the index 'i' from the range
    for i := range nums {
        // Access the slice directly at index 'i'
        nums[i] = nums[i] + 10 
    }

    fmt.Println("Fixed slice values:", nums)
}

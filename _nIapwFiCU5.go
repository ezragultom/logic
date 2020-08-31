package main

// you can also use imports, for example:
import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")


func main(){
	fmt.Println(Solution(276))
}

func Solution(N int) int {
    insN := int(5)
    if N == 0{
        return insN * 10
    }
    
    absN := int(math.Abs(float64(N)))
    count := 0
    newN := absN
	
    for newN > 0 {
        newN = newN / 10
	
        count ++
	
    }

    maxNumber := -8000
    digitPos := 1
    isNegative := N / absN
    
    for i:=0; i<= count; i++ {
        newNumber := ((absN/digitPos)*(digitPos*10))+(insN*digitPos)+ (absN%digitPos)
        if (newNumber * isNegative) > maxNumber{
            maxNumber = newNumber * isNegative
        }
	
        
        digitPos = digitPos * 10
    }
    
    return maxNumber
}
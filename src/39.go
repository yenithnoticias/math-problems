package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}
    fmt.Println("The highest number in the list is:", max(numbers))
}

func max(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }
    
    maxNum := numbers[0]
    for _, num := range numbers[1:] {
        if num > maxNum {
            maxNum = num
        }
    }
    return maxNum
}

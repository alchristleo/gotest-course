package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
        "strings"
)

const (
		Size            int    = 10
)

func Swap(numbers []int, i int) {
		tmp := numbers[i]
		numbers[i] = numbers[i+1]
		numbers[i+1] = tmp
}

func BubbleSort(numbers []int) {
        Swapped := true
        for Swapped {
                Swapped = false
                for i := 0; i < len(numbers)-1; i++ {
                        if numbers[i] > numbers[i+1] {
                                Swap(numbers, i)
                                Swapped = true
                        }
                }
        }
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        if err := scanner.Err(); err != nil {
                log.Println(err)
        }
        fmt.Printf("Please input list of integers: ")

        for (scanner.Scan()) {
                var strVal string = scanner.Text()
                var numlist = strings.Split(strVal, "")
                var intList = make([]int, 0)
                for _, i := range numlist {
                        j, err := strconv.Atoi(i)
                        if err != nil {
                            fmt.Printf("Please input list of integers: ")
                            continue
                        }
                        intList = append(intList, j)
                }
                fmt.Printf("length: %d; cap: %d\n", len(intList), cap(intList))
                if (len(intList) > Size) {
                        fmt.Printf("Please enter a maximum of %d integers\n", Size)
                        fmt.Printf("List of integers: ")
                        continue
                }
                BubbleSort(intList)
                fmt.Println(intList)
                fmt.Printf("List of integers: ")
        }
}
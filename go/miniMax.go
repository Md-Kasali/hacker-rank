package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    // Write your code here
    // Creates an array to store the sum values of each iteration
    var valArr []int64
    var sum int64
    // Calculate the number of elements in the array
    n := len(arr)
    // Looping over the array to calculate the sum
    for i := 0; i < n; i++ {
	// Initialize the sum value to zero at every iteration to clear the previous sum value
        sum = 0
        for j := 0; j < n; j++ {
	    // Check the index to ignore adding the value to sum. In each iteration one of the elements in the array will be ignored and the other elements are added
            if i != j {
                sum += int64(arr[j])
            }
        }
	// After each iteration the calculated sum is added to the valArr (array)
        valArr = append(valArr, sum)
    }
    // Initialize the minimum and maximum value in the array to the first element in the valArr (array)
    var min int64 = valArr[0]
    var max int64 = valArr[0]
    // Iterate over the valArr (array) to find the minimum and maximum value in the array
    for i := 1; i < n; i++ {
        if valArr[i] < min {
            min = valArr[i]
        }
        if valArr[i] > max {
            max = valArr[i]
        }
    }
    // Print the minimum and maximum value in the valArr (array)
    fmt.Printf("%d %d\n", min, max)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    // Calculate the number of elements in the array
    n := len(arr)
    // Initialize the variables to track number of positive, negative and zero in the array
    positiveNum := 0
    negativeNum := 0
    numOfZero := 0
    
    // Iterate over the array to find positive, negative and zero
    for i := 0; i < n; i++ {
        if arr[i] > 0 {
            positiveNum++
        } else if arr[i] < 0 {
            negativeNum++
        } else {
          numOfZero++  
        }
    }
    // Print the ratio of number of positive, negative and zero values present in the array 
    fmt.Printf("%.6f\n", float64(positiveNum)/float64(n)) 
    fmt.Printf("%.6f\n", float64(negativeNum)/float64(n))
    fmt.Printf("%.6f\n", float64(numOfZero)/float64(n))  
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    plusMinus(arr)
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

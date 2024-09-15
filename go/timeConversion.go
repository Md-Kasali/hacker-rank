package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Write your code here
    
    // Split the time string (s) into a slice with : as a seperator ex. ["12", "24", "00AM"]
    strArr := strings.Split(s, ":")
    
    // Initializing time suffix to "AM"
    timeSuffix := "AM"
    // Store the the hour of the time in the hour variable
    hour, _ := strconv.Atoi(strArr[0])
    
    // Check the time suffix to change the value of hour   
    if b := strings.HasSuffix(strArr[2], "PM"); b == true {
        // Update the timeSuffix variable when it is PM
        timeSuffix = "PM"
    }
    if timeSuffix == "PM" {
        // Condition to check that the time has PM suffix and it is 12'o clock
        // If it is the case, doesn't update the hour otherwise update the hour with the else block
        if hour == 12 {
            strArr[0] = strconv.Itoa(hour)
        } else {
            hour = hour + 12
            strArr[0] = strconv.Itoa(hour)
            // fmt.Println(strArr[0]) 
        }
    } else {
        // Check the hour is 12 and it is AM, Update the time to 00 hour - 12'o clock at night
        // in 24 hour format is -> 00:00:00AM 
        if val, _ := strconv.Atoi(strArr[0]); val == 12 {
            strArr[0] = "00"
        }
    }
    // Format the time stored in the slice to string
    fTime := strings.Join(strArr, ":")
    // fmt.Println(fTime)
    
    // Return the 24 hour formatted time string
    return strings.Trim(fTime, "PM AM")

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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

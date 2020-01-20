package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strings"
        "strconv"
       )

// Error checking function that will panic if an error is present
func checkError(err error) {
    if err != nil {
        log.Panicf("Error: %v", err)
    }
}

// Read input file
func readFile(filename string) {
    file, err := os.Open(filename)
    checkError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    // Iterate over each line and reverse them if it is after the first line and not an int, since we have a defined format
    // We keep track of the iteration for the formatted output of "Case #x: "
    count := 0
    for scanner.Scan() {
        text := scanner.Text()
        if count != 0 {
            if _, err := strconv.Atoi(text); err != nil {
                fmt.Println(fmt.Sprintf("Case #%d: %s", count, reverseWords(text)))
            }
        }
        count++
    }
}

// Returns a string that has been split on a whitespace character (word), reversed, and joined back together with a space
func reverseWords(str string) string {
    words := strings.Fields(str)

    // Replace the contents of the "words" slice with same elements in reverse order
    // https://github.com/golang/go/wiki/SliceTricks#reversing
    for left, right := 0, len(words) - 1; left < right; left, right = left + 1, right - 1 {
        words[left], words[right] = words[right], words[left]
    }
    return strings.Join(words, " ")
}

// Main Entry Point
func main() {
    readFile("./B-small-practice.in")
    readFile("./B-large-practice.in")
}

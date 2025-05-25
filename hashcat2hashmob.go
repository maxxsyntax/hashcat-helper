//take output from hashcat --show and put it into a format that hashmob accesspets.  only tested on 22000
package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func grepInFile(filename string, searchTerm string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var matches []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, searchTerm) {
            matches = append(matches, line)
        }
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return matches, nil
}

func main() {
    otherFilename := "otherfile.txt" // File to grep in

    scanner := bufio.NewScanner(os.Stdin)
    pattern := `([a-f0-9]+:){3}[A-Za-z0-9_\.\ ]+(-[A-Za-z0-9_\.\ ]+)?:[0-9]{8}`
    re := regexp.MustCompile(pattern)

    for scanner.Scan() {
        line := scanner.Text()
        matches := re.FindAllString(line, -1)
        for _, match := range matches {
            parts := strings.Split(match, ":")
            if len(parts) > 0 {
                searchTerm := parts[0]
                lines, err := grepInFile(otherFilename, searchTerm)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Error searching in file: %v\n", err)
                    continue
                }
                for _, foundLine := range lines {
                    fmt.Printf("%s:%s\n", foundLine, parts[4])
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
    }
}


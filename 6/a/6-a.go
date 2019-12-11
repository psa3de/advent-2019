package main

import (
    "strings"
    "bufio"
    "os"
    "log"
    "fmt"
)

type orbit struct {
    center string
    children []orbit
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    var txtlines []string

    for scanner.Scan() {
        txtlines = append(txtlines, scanner.Text())
    }
    file.Close()

    var orbits = []orbit{}

    for _, eachline := range txtlines {
        stringValues := strings.Split(eachline, ")")
        foundMatch := false
        for index, value := range orbits {
            if value.center == stringValues[0] {
                orbits[index].children = append(orbits[index].children, orbit{stringValues[0], []orbit{orbit{stringValues[1],nil}}})
                foundMatch = true
            }
        }
        if(!foundMatch) {
            orbits = append(orbits, orbit{stringValues[0], []orbit{orbit{stringValues[1],nil}}})
        }
        fmt.Println(orbits[0].center)

    }
}
package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strconv"
)

func mathFunction(Input int) (int, error) {
    return Input/3-2, nil
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

        var x int = 0
    	for _, eachline := range txtlines {
    //	    fmt.Printf(eachline)
    	    z, _ := strconv.Atoi(eachline)
    	    var y = z
    	    for y > 0 {
    	        y, _ = mathFunction(y)
    	        if y > 0 {
                  //  fmt.Printf("Adding %d", y)
    		        x += y
    		    }
            }
    	}

    fmt.Printf("%d ",x)
}
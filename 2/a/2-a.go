package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
)

func opcodeOne(Index int, Input []int) ([]int, error) {
    firstIndex := Input[Index+1]
    secondIndex := Input[Index+2]
    thirdIndex := Input[Index+3]
    firstValue := Input[firstIndex]
    secondValue := Input[secondIndex]
    calculatedValue := firstValue + secondValue
    modifiedInput := Input[0:thirdIndex]
    modifiedInput = append(modifiedInput, calculatedValue)
    modifiedInput = append(modifiedInput, Input[thirdIndex+1:len(Input)]...)
    return Input, nil
}

func opcodeTwo(Index int, Input []int) ([]int, error) {
    firstIndex := Input[Index+1]
    secondIndex := Input[Index+2]
    thirdIndex := Input[Index+3]
    firstValue := Input[firstIndex]
    secondValue := Input[secondIndex]
    calculatedValue := firstValue * secondValue
    modifiedInput := Input[0:thirdIndex]
    modifiedInput = append(modifiedInput, calculatedValue)
    modifiedInput = append(modifiedInput, Input[thirdIndex+1:len(Input)]...)
    return Input, nil
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

    	for _, eachline := range txtlines {
    	    stringValues := strings.Split(eachline, ",")
    	    var intValues = []int{}
    	    for _, i := range stringValues {
    	        j, err := strconv.Atoi(i)
    	        if err != nil {
    	            panic(err)
    	        }
    	        intValues = append(intValues, j)
    	    }

            var initialValues = []int{}
            initialValues = append(initialValues, intValues[0])
            initialValues = append(initialValues, 12, 2)
            initialValues = append(initialValues, intValues[3:len(intValues)]...)
            fmt.Println(initialValues)

            EvaluationLoop:
            for i := 0; i<len(initialValues); i+=4 {
                opcode := initialValues[i]
                switch opcode {
                case 1:
                    initialValues, _ = opcodeOne(i, initialValues)
                case 2:
                    initialValues, _ = opcodeTwo(i, initialValues)
                case 99:
                    break EvaluationLoop
                default:
                    fmt.Printf("%d", opcode)
                    fmt.Println(" not implemented")
                }
                fmt.Println(initialValues)
                }
    	}

}
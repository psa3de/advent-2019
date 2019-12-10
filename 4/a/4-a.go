package main

import (
    "fmt"
    "strconv"
    "strings"
    "sort"
)

func sorted(a, b []int) bool {
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func main() {
    numPasswords := 0
    for i:= 245318; i <= 765747; i++ {
        stringValue := strconv.Itoa(i)
        characters := strings.Split(stringValue, "")
        intValues := []int{}
    	for _, j := range characters {
    	    k, err := strconv.Atoi(j)
    	    if err != nil {
    	        panic(err)
    	    }
    	    intValues = append(intValues, k)
    	}
    	intValuesCopy := append([]int(nil), intValues...)
        sort.Ints(intValuesCopy)
        if(sorted(intValues, intValuesCopy)) {
            if(len(unique(intValues)) != 6) {
                numPasswords ++
            }
        }
    }
    fmt.Println(numPasswords)
}
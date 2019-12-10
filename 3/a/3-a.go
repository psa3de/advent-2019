package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
    "strconv"
    "math"
    "sort"
)

// Strategy:
// Read input
// Move counter
// Convert each input segment into a line segment
// Store segments in two separate arrays
// Iterate through segments in array 1, check for overlaps with each element of array 2

// Make a struct for line segment
// x1, y1, x2, y2
// vertical? boolean

// Overlap:
// If array 2 segment is vertical
//  If array 1 segment is horizontal
//      If min(1.x) <= 2.x1/2.x2 <= max(1.x)
//          Intersect
// Else
//  If array 1 is vertical
//      If min(1.y) <= 2.y1/2.y2 <= max(1.y2)
//          Intersect
//

func (this lineSegment) String() string {
    return "{" + strconv.Itoa(this.x1) + ", " + strconv.Itoa(this.y1) + "}, {" + strconv.Itoa(this.x2) + ", " + strconv.Itoa(this.y2) + "}"
}

func getManhattanDistance(x float64, y float64) (int, error) {
    xDiff := math.Abs(x)
    yDiff := math.Abs(y)
    return int(xDiff + yDiff), nil
}

func getOverlaps(wire1 []lineSegment, wire2 []lineSegment) ([]int, error){
    intersectionDists := []int{}
    for _, element := range wire1 {
        for _, element2 := range wire2 {
            if (!element.vertical && element2.vertical) {
                firstWireX1 := float64(element.x1)
                firstWireX2 := float64(element.x2)
                firstWireY := float64(element.y1)
                secondWireY1 := float64(element2.y1)
                secondWireY2 := float64(element2.y2)
                secondWireX := float64(element2.x1)
                if(math.Min(firstWireX1, firstWireX2) < secondWireX && secondWireX < math.Max(firstWireX1, firstWireX2) && math.Min(secondWireY1, secondWireY2) < firstWireY && firstWireY < math.Max(secondWireY1, secondWireY2)) {
                    intersectY := float64(element.y1)
                    intersectionDist, _ := getManhattanDistance(secondWireX, intersectY)
                    intersectionDists = append(intersectionDists, intersectionDist)
                }
            } else {
                if (element.vertical && !element2.vertical) {
                    firstWireY1 := float64(element.y1)
                    firstWireY2 := float64(element.y2)
                    firstWireX := float64(element.x1)
                    secondWireX1 := float64(element2.x1)
                    secondWireX2 := float64(element2.x2)
                    secondWireY := float64(element2.y1)
                    if(math.Min(firstWireY1, firstWireY2) < secondWireY && secondWireY < math.Max(firstWireY1, firstWireY2) && math.Min(secondWireX1, secondWireX2) < firstWireX && firstWireX < math.Max(secondWireX1, secondWireX2)) {
                        intersectX := float64(element.x1)
                        intersectionDist, _ := getManhattanDistance(intersectX, secondWireY)
                        intersectionDists = append(intersectionDists, intersectionDist)
                    }
                }
            }
        }
    }
    return intersectionDists, nil
}

type lineSegment struct {
    x1 int
    y1 int
    x2 int
    y2 int
    vertical bool
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

    var wires [][]lineSegment

    for _, eachline := range txtlines {
        var wire []lineSegment
        var x,y = 0,0
        stringValues := strings.Split(eachline, ",")
        for i := 0; i<len(stringValues); i++ {
            entry := stringValues[i]
            direction := entry[0:1]
            distance := entry[1:len(entry)]
            intDistance, _ := strconv.Atoi(distance)
            switch direction {
                case "R":
                    endingX := x + intDistance
                    wire = append(wire, lineSegment{x, y, endingX, y, false})
                    x = endingX
                case "L":
                    endingX := x - intDistance
                    wire = append(wire, lineSegment{x, y, endingX, y, false})
                    x = endingX
                case "U":
                    endingY := y + intDistance
                    wire = append(wire, lineSegment{x, y, x, endingY, true})
                    y = endingY
                case "D":
                    endingY := y - intDistance
                    wire = append(wire, lineSegment{x, y, x, endingY, true})
                    y = endingY
            }
        }
        wires = append(wires, wire)
    }
    intersectionDists , _ := getOverlaps(wires[0], wires[1])
    sort.Ints(intersectionDists)
    fmt.Println(intersectionDists)
}
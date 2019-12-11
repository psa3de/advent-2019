package main

import (
    "strings"
    "bufio"
    "os"
    "log"
    "fmt"
    "math"
)

type asteroid struct {
    x int
    y int
}

type asteroidCandidate struct {
    a asteroid
    numAsteroidsSeen int
}

type asteroidEvaluation struct {
    xDiff float64
    yDiff float64
    ratio float64
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

    var asteroids = []asteroid{}
    for x, eachLine := range txtlines {
        stringValues := strings.Split(eachLine, "")
        for y, eachChar := range stringValues {
            if eachChar == "#" {
                asteroids = append(asteroids, asteroid{x,y})
            }
        }
    }

    var asteroidCandidates = []asteroidCandidate{}
    for _, asteroid := range asteroids {
        var ratios = []asteroidEvaluation{}
        for _, otherAsteroid := range asteroids {
            if(otherAsteroid != asteroid) {
                xDiff := float64(otherAsteroid.x - asteroid.x)
                yDiff := float64(otherAsteroid.y - asteroid.y)
                ratio := float64(xDiff) / float64(yDiff)
                foundMatch := false
                for _, entry := range ratios {
                    if ratio == entry.ratio {
                        if(math.Signbit(xDiff) == math.Signbit(entry.xDiff) && math.Signbit(yDiff) == math.Signbit(entry.yDiff)) {
                                foundMatch = true
                        }
                    }
                }
                if !foundMatch {
                    ratios = append(ratios, asteroidEvaluation{xDiff, yDiff, ratio})
                }
            }
        }
        asteroidCandidates = append(asteroidCandidates, asteroidCandidate{asteroid, len(ratios)})
    }

    maxAsteroids := 0
    for _, candidate := range asteroidCandidates {
        if(candidate.numAsteroidsSeen > maxAsteroids) {
            fmt.Println(candidate.numAsteroidsSeen, candidate.a.x, candidate.a.y)
            maxAsteroids = candidate.numAsteroidsSeen
        }
    }

}
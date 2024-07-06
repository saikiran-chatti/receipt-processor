package main

import (
    "math"
    "regexp"
)

func CalculatePoints(receipt Receipt) int {
    points := 0

    points += len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(receipt.Retailer, -1))

    if receipt.Total == math.Floor(receipt.Total) {
        points += 50
    }


    return points
}
package experiment

import (
	"fmt"
	"time"
)

var romans = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func RomanToInteger_Current(roman string) int {
	start := time.Now() // 获取当前时间
	total := 0
	holder := 0
	for holder < len(roman) {
		if holder+1 < len(roman) && romans[string(roman[holder])] < romans[string(roman[holder+1])] {
			total += romans[string(roman[holder+1])] - romans[string(roman[holder])]
			holder += 2
		} else {
			total += romans[string(roman[holder])]
			holder++
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time-consuming to execute RomanToInteger_Current: ", elapsed)
	return total
}

func RomanToInteger_Target(roman string) int{
	start := time.Now()
	total := 0
    romanLen := len(roman)
    for i := range roman {
        if i + 1 < romanLen && romans[string(roman[i])] < romans[string(roman[i + 1])] {
            total -= romans[string(roman[i])]
        } else {
            total += romans[string(roman[i])]
        }
    }
    return total
	fmt.Println("Time-consuming to execute RomanToInteger_Target : ", elapsed)
	return total
}
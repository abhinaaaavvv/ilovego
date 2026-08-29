package main

import "fmt"

type Coord struct {
	lat, long float64
}

func main() {
	m := map[string]int{}
	n := make(map[string]int)

	m["Answer"] = 42
	n["Answer"] = 50

	a := map[string]int{
		"Rollno.": 19,
	}

	b := map[string]Coord{
		"bangalore": {
			12.9716, 77.5946,
		},
		"hyderabad": {
			17.3850, 78.4867,
		},
	}

	fmt.Println(a["Rollno."])
	fmt.Println(b["bangalore"])
	fmt.Println(b["hyderabad"])

}

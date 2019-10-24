// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// // b = beginning
// // e = end
// // Return an array of all possible points
// func genShipMap(b []int, e []int) []int {
// 	result := []int{}
// 	for i := b[0]; i <= e[0]; i++ {
// 		for j := b[1]; j <= e[1]; j++ {
// 			result = append(result, i)
// 			result = append(result, j)
// 		}
// 	}
// 	return result
// }

// // convert to
// func convertToSlice(S string) []string {
// 	// using different vars because of type mismatch
// 	result := strings.Split(S, ",")
// 	temp_S := strings.Join(result, " ")
// 	result = strings.Fields(temp_S)

// 	return result
// }

// func convertToString(I []int) string {
// 	return strings.Trim(strings.Replace(fmt.Sprint(I), " ", " ", -1), "[]")
// }

// // generate point location as []int
// func genPointLocation(S string) []int {
// 	result := []int{}

// 	horisontal_rule, _ := strconv.Atoi(string(S[0]))
// 	// DIRTY DIRTY HACK
// 	// Converting letters to numbers using ascii codes
// 	vertical_rule := []rune(string(S[1]))[0] - 64

// 	result = append(result, horisontal_rule)
// 	result = append(result, int(vertical_rule))
// 	return result
// }

// func Solution(N int, S string, T string) string {
// 	stat := []int{0, 0}

// 	// convert string to slice
// 	// NewS is a converted S to slice
// 	// NewT is a converted T to slice
// 	NewS := convertToSlice(S)
// 	NewT := convertToSlice(T)

// 	// jumping over odd value
// 	for i := 0; i < len(NewS); i = i + 2 {
// 		// counter of sucessful shots
// 		succesfullShots := 0

// 		// ship takes a value of a slice with all available points
// 		ship := genShipMap(
// 			genPointLocation(NewS[i]),
// 			genPointLocation(NewS[i+1]))

// 		shot := []int{}

// 		// jumping over odd value but already using ship length
// 		for j := 0; j < len(ship); j = j + 2 {
// 			for k := 0; k < len(NewT); k++ {
// 				// convert 1A to 11, 2D to 24 and etc.
// 				shot = genPointLocation(NewT[k])

// 				if ship[j] == shot[0] && ship[j+1] == shot[1] {
// 					fmt.Println("BAH", ship[j], ship[j+1])
// 					succesfullShots += 1
// 				}
// 			}
// 		}

// 		// check if ship has been destroyd or just damaged
// 		if succesfullShots*2 == len(ship) {
// 			stat[1] += 1
// 		} else {
// 			stat[0] += 1
// 		}

// 	}
// 	fmt.Println(stat)
// 	return convertToString(stat)
// }

// func main() {
// 	Solution(5, "1A 2A,2D 4D", "2A,3B,2D,4D,3D")
// }

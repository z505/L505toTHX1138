// Convert id L505 to THX1138 using only the originating id for calculation and
// final conversion. i.e. take L505 and convert to THX1138 using only math

package main

import (
	"fmt"
	"strconv"
)

func getInt(char byte) int {
	i, _ := strconv.Atoi(string(char))
	return i
}

// get number of alphabit, i.e. L is 12th number of alphabit
func getAlphabitNum(letter string) int {
	i := (int)(letter[0])
	return i - 64
}

// convert numbers in alphabit (i.e. L is 12th letter) to string
func numToAlphabit(char1,char2,char3 int) string {
	return string(char1+64)+string(char2+64)+string(char3+64)
}

/*

   (L-(5+0+5))(5+0+5)  = 20               (T)
   (12-(5_0_5))(5+0+5) = 20

   (L-(5+0+5))(L-(5+0+5))(L-(5+0+5)) = 8
   12-(5+5)=2*(12-(5+5))*(12-(5+5))  = 8  (H)

	L*(L-5+0+5)
	12*(12-(5+5)) = 24                    (X)
*/

// Convert L505 to THX1138 using only characters and numbers available in
// originating ID and nothing more
func convert(id string) string {
	if id != "L505" {
		return "ERROR"
	}
	_L := getAlphabitNum(id[0:])  // "L" = 12th letter of alphabit
	// convert to 1138 using only L505 as basis
	_1138 := 505*((5+0+5)/5)+((_L*(5+0+5))+(5+0+5))-((5+0+5)/5) // 1138
	// convert to THX using only L505 as basis
	_THX := numToAlphabit(
			(_L-(5+0+5))*(5+0+5),                     // T
			(_L-(5+0+5))*(_L-(5+0+5))*(_L-(5+0+5)),   // H
			_L*(_L-(5+0+5)))                          // X
	result := _THX+strconv.Itoa(_1138)
	return result
}

func main() {
	result := convert("L505")
	fmt.Println(result)
}
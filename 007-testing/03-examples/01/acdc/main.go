// package acdc asks if you're ready to rock
package acdc

// Sum adds an unlimited number of balues of type int
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}

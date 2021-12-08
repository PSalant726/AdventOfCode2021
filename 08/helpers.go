package main

var digitCounts = map[int]int{
	0: 0,
	1: 0,
	2: 0,
	3: 0,
	4: 0,
	5: 0,
	6: 0,
	7: 0,
	8: 0,
	9: 0,
}

func copyMap(origin map[int]int) map[int]int {
	destination := make(map[int]int, len(origin))
	for k, v := range origin {
		destination[k] = v
	}

	return destination
}

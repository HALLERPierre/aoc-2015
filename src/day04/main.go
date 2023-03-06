package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

const Input = "iwrupvqb"

func main() {
	lowestInt5zeroes := findLowestIntForPrefix("00000")
	fmt.Printf("lowest int 5: %v\n", lowestInt5zeroes)

	lowestInt6zeroes := findLowestIntForPrefix("000000")
	fmt.Printf("lowest int 6: %v\n", lowestInt6zeroes)
}

func findLowestIntForPrefix(prefix string) int {
	for i := 0; true; i++ {
		hash := md5.Sum([]byte(Input + strconv.Itoa(i)))
		stringHash := hex.EncodeToString(hash[:])

		if strings.HasPrefix(stringHash, prefix) {
			return i
		}
	}
	return 0
}

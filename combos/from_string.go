package combos

import (
	"fmt"
	"strconv"
)

var letters = [][]string{
	[]string{""},                 // digit: 0
	[]string{" "},                // digit: 1
	[]string{"A", "B", "C"},      // digit: 2
	[]string{"D", "E", "F"},      // digit: 3
	[]string{"G", "H", "I"},      // digit: 4
	[]string{"J", "K", "L"},      // digit: 5
	[]string{"M", "N", "O"},      // digit: 6
	[]string{"P", "Q", "R", "S"}, // digit: 7
	[]string{"T", "U", "V"},      // digit: 8
	[]string{"W", "X", "Y", "Z"}, // digit: 9
}

// FromString returns all combinations for the given digits
func FromString(input string) ([]string, error) {
	var results []string

	for _, chr := range input {
		digit, err := strconv.ParseInt(string(chr), 10, 64)

		if err != nil {
			return nil, fmt.Errorf(
				"invalid character %q in input",
				chr,
			)
		}

		if len(results) == 0 {
			results = letters[digit]
			continue
		}

		prefixes := make([]string, len(results))
		copy(prefixes, results)
		results = nil

		for _, letter := range letters[digit] {
			for _, prefix := range prefixes {
				results = append(results, prefix+letter)
			}
		}
	}

	return results, nil
}

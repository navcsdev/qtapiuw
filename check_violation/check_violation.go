package check_violation

import "fmt"

const ALPHABET_MATRIX_SIZE = 676
const ALPHABET_SIZE = 26

func EmptyTrack(parentNodeMatrix *[ALPHABET_SIZE][ALPHABET_SIZE]int) {
	for i := 0; i < ALPHABET_MATRIX_SIZE; i++ {
		parentNodeMatrix[i/ALPHABET_SIZE][i%ALPHABET_SIZE] = 0
	}
}

func MarkNode(node uint, parentNodes *[ALPHABET_SIZE]int, nodeMatrix *[ALPHABET_SIZE][ALPHABET_SIZE]int, parentNodeMatrix *[ALPHABET_SIZE][ALPHABET_SIZE]int) {
	for i := 0; i < ALPHABET_SIZE; i++ {
		if nodeMatrix[node][i] != 0 && parentNodeMatrix[node][i] == 0 {
			parentNodeMatrix[node][i] = 1
			parentNodes[i] = parentNodes[i] | parentNodes[node]
			MarkNode(uint(i), parentNodes, nodeMatrix, parentNodeMatrix)
		}
	}
}

func IsViolation(relationships [][2]uint) bool {
	var parentNodes [ALPHABET_SIZE]int
	var nodeMatrix [ALPHABET_SIZE][ALPHABET_SIZE]int
	var parentNodeMatrix [ALPHABET_SIZE][ALPHABET_SIZE]int

	for _, relationship := range relationships {
		left := relationship[0] - 'A'
		right := relationship[1] - 'A'
		if nodeMatrix[right][left] == 0 {
			if ((1 << left) & parentNodes[right]) != 0 {
				fmt.Println("Clause Violation", string(relationship[0]), "<", string(relationship[1]))
				return true
			} else {
				nodeMatrix[right][left] = 1
				parentNodes[left] |= (parentNodes[right] | (1 << right))
				EmptyTrack(&parentNodeMatrix)
				MarkNode(left, &parentNodes, &nodeMatrix, &parentNodeMatrix)
			}
		}
	}
	return false
}

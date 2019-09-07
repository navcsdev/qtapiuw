package check_violation

import "testing"

func TestIsViolationTrue(t *testing.T) {
	var relationships [][2]uint

	relationships = append(relationships, [2]uint{'A', 'B'})
	relationships = append(relationships, [2]uint{'B', 'C'})
	relationships = append(relationships, [2]uint{'E', 'B'})
	relationships = append(relationships, [2]uint{'C', 'A'})

	result := IsViolation(relationships)
	if result {
		t.Logf("relationships has violation")
	} else {
		t.Errorf("check violation incorrect")
	}
}

func TestIsViolationTrueMoreRelationships(t *testing.T) {
	var relationships [][2]uint

	relationships = append(relationships, [2]uint{'O', 'A'})
	relationships = append(relationships, [2]uint{'A', 'B'})
	relationships = append(relationships, [2]uint{'B', 'C'})
	relationships = append(relationships, [2]uint{'E', 'B'})
	relationships = append(relationships, [2]uint{'A', 'G'})
	relationships = append(relationships, [2]uint{'G', 'H'})
	relationships = append(relationships, [2]uint{'H', 'J'})
	relationships = append(relationships, [2]uint{'C', 'Z'})
	relationships = append(relationships, [2]uint{'C', 'K'})
	relationships = append(relationships, [2]uint{'K', 'A'})

	result := IsViolation(relationships)
	if result {
		t.Logf("relationships has violation")
	} else {
		t.Errorf("check violation incorrect")
	}
}

func TestIsViolationFalseMoreRelationships(t *testing.T) {
	var relationships [][2]uint

	relationships = append(relationships, [2]uint{'O', 'A'})
	relationships = append(relationships, [2]uint{'A', 'B'})
	relationships = append(relationships, [2]uint{'B', 'C'})
	relationships = append(relationships, [2]uint{'E', 'B'})
	relationships = append(relationships, [2]uint{'A', 'G'})
	relationships = append(relationships, [2]uint{'G', 'H'})
	relationships = append(relationships, [2]uint{'H', 'J'})
	relationships = append(relationships, [2]uint{'C', 'Z'})
	relationships = append(relationships, [2]uint{'C', 'K'})

	result := IsViolation(relationships)
	if result {
		t.Errorf("check violation incorrect")
	} else {
		t.Logf("relationships has violation")
	}
}

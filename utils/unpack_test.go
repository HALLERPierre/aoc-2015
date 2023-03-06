package utils

import "testing"

func TestUnpack(t *testing.T) {
	var s1, s2, s3 string
	toUnpack := []string{"1", "2", "3"}

	Unpack(toUnpack, &s1, &s2, &s3)

	if s1 != "1" || s2 != "2" || s3 != "3" {
		t.Errorf("Unpack failed")
	}
}

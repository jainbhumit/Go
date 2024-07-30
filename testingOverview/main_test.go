package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)

	if total != 10 {
		t.Errorf("The value is %d  and expected is : %d", total, 10)

	}

}

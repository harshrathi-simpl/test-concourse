// main_test.go
package main

import "testing"

// TestAdd is the unit test for the Add function
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 6

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
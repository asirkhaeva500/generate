package main

import (
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(result))
	}
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Expected slice length %d, got %d", size, len(result))
	}
	for i, v := range result {
		if v <= 0 {
			t.Errorf("Element at index %d is not positive: %d", i, v)
		}
	}
}

func TestMaximum_EmptySlice(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", result)
	}
}

func TestMaximum_OneElement(t *testing.T) {
	result := maximum([]int{42})
	if result != 42 {
		t.Errorf("Expected 42 for single element, got %d", result)
	}
}

func TestMaximum_MultipleElements(t *testing.T) {
	result := maximum([]int{1, 2, 3, 4, 5})
	if result != 5 {
		t.Errorf("Expected 5 for max, got %d", result)
	}
}

func TestMaximum_AllSame(t *testing.T) {
	result := maximum([]int{7, 7, 7, 7})
	if result != 7 {
		t.Errorf("Expected 7 for all same, got %d", result)
	}
}

func TestMaxChunks_EmptySlice(t *testing.T) {
	result := maxChunks([]int{})
	if result != 0 {
		t.Errorf("Expected 0 for empty slice, got %d", result)
	}
}

func TestMaxChunks_OneElement(t *testing.T) {
	result := maxChunks([]int{99})
	if result != 99 {
		t.Errorf("Expected 99 for single element, got %d", result)
	}
}

func TestMaxChunks_MultipleElements(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := maxChunks(data)
	if result != 8 {
		t.Errorf("Expected 8 for max, got %d", result)
	}
}

func TestMaxChunks_AllSame(t *testing.T) {
	data := []int{5, 5, 5, 5, 5, 5, 5, 5}
	result := maxChunks(data)
	if result != 5 {
		t.Errorf("Expected 5 for all same, got %d", result)
	}
}

package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name         string
		size         int
		expectedSize int
	}{
		{name: "нулевой размер", size: 0, expectedSize: 0},
		{name: "отрицательный размер", size: -5, expectedSize: 0},
		{name: "один элемент", size: 1, expectedSize: 1},
		{name: "много элементов", size: 1000, expectedSize: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elements := generateRandomElements(tt.size)

			if elements == nil {
				t.Fatal("ожидался пустой слайс, получен nil")
			}
			if len(elements) != tt.expectedSize {
				t.Fatalf("неверный размер слайса: ожидалось %d, получено %d", tt.expectedSize, len(elements))
			}
			for i, value := range elements {
				if value <= 0 {
					t.Fatalf("элемент с индексом %d не положительный: %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{name: "nil-слайс", data: nil, expected: 0},
		{name: "пустой слайс", data: []int{}, expected: 0},
		{name: "один элемент", data: []int{42}, expected: 42},
		{name: "максимум в начале", data: []int{9, 3, 5, 1}, expected: 9},
		{name: "максимум в конце", data: []int{1, 3, 5, 9}, expected: 9},
		{name: "максимум в середине", data: []int{1, 9, 5, 3}, expected: 9},
		{name: "одинаковые элементы", data: []int{7, 7, 7}, expected: 7},
		{name: "отрицательные элементы", data: []int{-5, -1, -10}, expected: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := maximum(tt.data); actual != tt.expected {
				t.Errorf("неверный максимум: ожидалось %d, получено %d", tt.expected, actual)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{name: "nil-слайс", data: nil, expected: 0},
		{name: "пустой слайс", data: []int{}, expected: 0},
		{name: "элементов меньше, чем CHUNKS", data: []int{3, 8, 1}, expected: 8},
		{name: "элементов ровно CHUNKS", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 8},
		{name: "максимум в остатке последней части", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 99}, expected: 99},
		{name: "максимум в первой части", data: []int{99, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, expected: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := maxChunks(tt.data); actual != tt.expected {
				t.Errorf("неверный максимум: ожидалось %d, получено %d", tt.expected, actual)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	sizes := []int{CHUNKS, CHUNKS + 1, 1000, 100_003}

	for _, size := range sizes {
		data := generateRandomElements(size)

		expected := maximum(data)
		if actual := maxChunks(data); actual != expected {
			t.Errorf("размер %d: результаты расходятся, maximum вернул %d, maxChunks вернул %d", size, expected, actual)
		}
	}
}

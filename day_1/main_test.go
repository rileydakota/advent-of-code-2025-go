package main

import "testing"

func TestRotateDial(t *testing.T) {
	tests := []struct {
		name     string
		dial     int
		Rotation Rotation
		expected int
	}{
		{
			name: "rotate left 1 click overflow",
			dial: 0,
			Rotation: Rotation{
				direction: DirLeft,
				clicks:    1,
			},
			expected: 99,
		},
		{
			name: "rotate dial right 1 click",
			dial: 0,
			Rotation: Rotation{
				direction: DirRight,
				clicks:    1,
			},
			expected: 1,
		},
		{
			name: "Rotate multiple out of bounds",
			dial: 0,
			Rotation: Rotation{
				direction: DirRight,
				clicks:    200,
			},
			expected: 99,
		},
		{
			name: "Rotate multiple out of bounds",
			dial: 0,
			Rotation: Rotation{
				direction: DirRight,
				clicks:    100,
			},
			expected: 99,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := rotateDial(test.dial, test.Rotation)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// func TestRotations(t *testing.T) {
// 	rotations := []Rotation{
// 		{
// 			direction: DirLeft,
// 			clicks:    1,
// 		},
// 		{
// 			direction: DirRight,
// 			clicks:    1,
// 		},
// 		{
// 			direction: DirLeft,
// 			clicks:    10,
// 		},
// 		{
// 			direction: DirRight,
// 			clicks:    10,
// 		},
// 	}
// }

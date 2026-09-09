package util

import "testing"

func TestCaseSensitive(t *testing.T) {
	cases := []struct {
		Name     string
		Left     string
		Right    string
		Suppress bool
	}{
		{
			Name:     "empty",
			Left:     "",
			Right:    "",
			Suppress: true,
		},
		{
			Name:     "empty and text",
			Left:     "text",
			Right:    "",
			Suppress: false,
		},
		{
			Name:     "different text",
			Left:     "text",
			Right:    "different text",
			Suppress: false,
		},
		{
			Name:     "same text",
			Left:     "text",
			Right:    "text",
			Suppress: true,
		},
		{
			Name:     "same text different case",
			Left:     "text",
			Right:    "TeXT",
			Suppress: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			if CaseSensitive("test", tc.Left, tc.Right, nil) != tc.Suppress {
				t.Fatalf("Expected CaseSensitive to return %t for '%q' == '%q'", tc.Suppress, tc.Left, tc.Right)
			}
		})
	}
}

func TestFloat32Precision(t *testing.T) {
	cases := []struct {
		Name     string
		Left     string
		Right    string
		Suppress bool
	}{
		{
			Name:     "exact match",
			Left:     "0.001",
			Right:    "0.001",
			Suppress: true,
		},
		{
			Name:     "float32 expansion of 0.001",
			Left:     "0.0010000000474974513",
			Right:    "0.001",
			Suppress: true,
		},
		{
			Name:     "zero is distinct from small values",
			Left:     "0",
			Right:    "0.001",
			Suppress: false,
		},
		{
			Name:     "materially different values",
			Left:     "0.1",
			Right:    "0.01",
			Suppress: false,
		},
		{
			Name:     "zero equals zero",
			Left:     "0",
			Right:    "0",
			Suppress: true,
		},
		{
			Name:     "invalid values are not suppressed",
			Left:     "not-a-number",
			Right:    "0.001",
			Suppress: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			if Float32Precision("test", tc.Left, tc.Right, nil) != tc.Suppress {
				t.Fatalf("Expected Float32Precision to return %t for '%q' == '%q'", tc.Suppress, tc.Left, tc.Right)
			}
		})
	}
}

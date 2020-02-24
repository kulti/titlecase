package titlecase_test

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/require"
)

func TestTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		minor    string
		expected string
	}{
		{"empty", "", "", ""},
		{"wo_str", "", "in of the", ""},
		{"wo_minor", "the quick brown fox", "", "The Quick Brown Fox"},
		{"minor_first", "the quick brown fox", "in of the", "The Quick Brown Fox"},
		{"full", "the quick brown fox in the bag", "in of the", "The Quick Brown Fox in the Bag"},
		{"with_capital", "tHe QUICK brOWn foX IN The bag", "in of the", "The Quick Brown Fox in the Bag"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.expected, titlecase.TitleCase(tc.str, tc.minor))
		})
	}
}

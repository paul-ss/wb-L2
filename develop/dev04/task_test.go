package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var findAnagramsTC = []struct {
	in  []string
	out map[string][]string
}{
	{
		in:  []string{"пятак", "пятка", "тяпка", "пятка", "листок", "слиток", "столик", "жопа", "апож", "1"},
		out: map[string][]string{"пятак": {"пятка", "тяпка"}, "листок": {"слиток", "столик"}, "жопа": {"апож"}},
	},
	{
		in:  []string{},
		out: map[string][]string{},
	},
	{
		in:  []string{"1", "2", "3"},
		out: map[string][]string{},
	},
	{
		in:  []string{"jopa", "jopa", "jopa"},
		out: map[string][]string{},
	},
	{
		in:  []string{"101", "101", "101", "110", "011", "101", "101", "110", "110", "011", "011"},
		out: map[string][]string{"101": {"011", "110"}},
	},
}

func TestFindAnagrams(t *testing.T) {
	for i, tc := range findAnagramsTC {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := FindAnagrams(tc.in)
			require.Equal(t, tc.out, got)
		})
	}
}

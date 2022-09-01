package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var unpackTc = []struct {
	in  string
	out string
}{
	{"a4bc2d5e", "aaaabccddddde"},
	{"abcd", "abcd"},
	{"", ""},
	{`qwe\4\5`, "qwe45"},
	{`qwe\45`, "qwe44444"},
	{`qwe\\5`, `qwe\\\\\`},
	{`qwe\\11`, `qwe\\\\\\\\\\\`},
	{`qwe\111`, `qwe11111111111`},
	{`qwe13`, `qweeeeeeeeeeeee`},
}

var unpackTcFail = []string{
	`\`,
	"45",
	`\q`,
}

func TestUnpack(t *testing.T) {
	for _, tc := range unpackTc {
		t.Run(tc.in, func(t *testing.T) {
			res, err := Unpack(tc.in)
			require.Nil(t, err)
			require.Equal(t, tc.out, res)
		})
	}
}

func TestUnpackFail(t *testing.T) {
	for _, tc := range unpackTcFail {
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.NotNil(t, err)
		})
	}
}

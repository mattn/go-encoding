package encoding

import (
	"io/ioutil"
	"strings"
	"testing"
)

var tests = []struct {
	encoding string
	input    string
	want     string
}{
	{encoding: "latin-1", input: "エラー", want: "\xc3\xa3\xc2\x82\xc2\xa8\xc3\xa3\xc2\x83\xc2\xa9\xc3\xa3\xc2\x83\xc2\xbc"},
	{encoding: "utf-8", input: "エラー", want: "エラー"},
	{encoding: "utf-16be", input: "\x30\xa8\x30\xe9\x30\xfc", want: "エラー"},
	{encoding: "utf-16le", input: "\xa8\x30\xe9\x30\xfc\x30", want: "エラー"},
	{encoding: "utf-16", input: "\x30\xa8\x30\xe9\x30\xfc", want: "エラー"},
	{encoding: "utf-16", input: "\xfe\xff\x30\xa8\x30\xe9\x30\xfc", want: "エラー"},
	{encoding: "utf-16", input: "\xff\xfe\xa8\x30\xe9\x30\xfc\x30", want: "エラー"},
	{encoding: "euc-jp", input: "\xa5\xa8\xa5\xe9\xa1\xbc", want: "エラー"},
}

func TestRunWithEncoding(t *testing.T) {
	for _, test := range tests {
		enc := GetEncoding(test.encoding)
		if enc == nil {
			t.Fatalf("unknown encoding name: %q", test.encoding)
		}
		input := strings.NewReader(test.input)
		got, err := ioutil.ReadAll(enc.NewDecoder().Reader(input))
		if err != nil {
			t.Fatal(err)
		}
		if test.want != string(got) {
			t.Fatalf("got %q (%#x) but want %q for %q", string(got), got, test.want, test.encoding)
		}
	}
}

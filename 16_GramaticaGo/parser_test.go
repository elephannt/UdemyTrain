package marco_test

import (
	"reflect"
	"strings"
	"testing"
	"github.com/elephannt/parsercompo"
	"io/ioutil"
)
//adding
// Ensure the parser can parse strings into Statement ASTs.
//update
//dd
//HOLA
func TestParser_ParseStatement(t *testing.T) {
	f, _ := ioutil.ReadFile("C://Download//Pokemon//gramatic.txt")
	r := string(f[:])

	var tests = []struct {
		s    string
		stmt *marco.SelectStatement
		err  string
	}{
		{
			s: r,
			stmt: &marco.SelectStatement{
				Fields:    []string{"*"},
				TableName: "marco",
			},
		},


		// Errors
		{s: `foo`, err: `found "foo", expected SELECT`},
		{s: `SELECT !`, err: `found "!", expected field`},
		{s: `SELECT field xxx`, err: `found "xxx", expected FROM`},
		{s: `SELECT field FROM *`, err: `found "*", expected table name`},
	}

	for i, tt := range tests {
		stmt, err := marco.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
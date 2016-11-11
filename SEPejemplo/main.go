package main
import (
"bufio"
//"bytes"
"io"
//"strings"
//"io/ioutil"
	"os"
	"strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {

	f, _ := os.Open("C:\\programs\\file.txt")
	scanner := bufio.NewScanner(f)

	// Set the Split method to ScanWords.
	scanner.Split(bufio.ScanWords)

return &Scanner{r: bufio.NewReader(r)}
}

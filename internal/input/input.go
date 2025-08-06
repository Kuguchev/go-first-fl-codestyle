package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader struct {
	reader *bufio.Reader
}

func NewInputReader() *Reader {
	return &Reader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (ir *Reader) ReadStringWithPrompt(prompt string) string {
	fmt.Printf("%s", prompt)
	return ir.read()
}

func (ir *Reader) read() string {
	input, _ := ir.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

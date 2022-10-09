package ch10

import (
	"io"
	"log"
	"os"
	"strings"
)

func ExampleCopy() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read
}

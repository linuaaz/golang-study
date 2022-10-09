package monkey

import (
	"fmt"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
)

func TestPrintlnByGoMonkey(t *testing.T) {
	patches := gomonkey.ApplyFunc(fmt.Println, func(a ...any) (n int, err error) {
		return fmt.Fprintln(os.Stdout, "I have changed the arguments")
	})
	defer patches.Reset()

	fmt.Println("hello world")
}

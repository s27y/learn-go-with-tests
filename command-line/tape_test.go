package poker_test

import (
	"io/ioutil"
	"testing"

	poker "github.com/s27y/learn-go-with-tests/command-line"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()
	tape := poker.NewTape(file)

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

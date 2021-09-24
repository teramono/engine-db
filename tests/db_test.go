package tests

import (
	"testing"

	"github.com/teramono/engine-db/pkg/engine"
)

func TestFs(t *testing.T) {
	fs := engine.NewFS()
	file := fs.Open("test.txt")
	file.Read()
}

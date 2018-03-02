package storage

import (
	"io"
	"testing"
)

func TestIsReader(t *testing.T) {
	var _ io.Reader = NewRecordReaderWriter(512)
}

func TestIsWriter(t *testing.T) {
	var _ io.Writer = NewRecordReaderWriter(512)
}

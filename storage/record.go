package storage

import (
	"fmt"
	"io"
	"sync"
)

type PageSize int

const (
	TwoKiloBytes     PageSize = 2096
	FourKiloBytes             = 4096
	EightKiloBytes            = 8192
	SixteenKiloBytes          = 16384
)

type PageReaderWriter interface {
	io.Reader
	io.Writer
	GetRecordReaderWriter() RecordReaderWriter
}

type RecordReaderWriter interface {
	//	ReadRecord(RecordSchema, Record) error
	//	WriteRecord(RecordSchema, Record) error
}

type pageReaderWriter struct {
	lock   sync.RWMutex
	buffer []byte
}

type recordReaderWriter struct {
	pageReaderWriter *pageReaderWriter
}

func NewPageReaderWriter(size PageSize) PageReaderWriter {
	return &pageReaderWriter{
	//buffer: make([]byte, int.(size)),
	}
}

func (prw *pageReaderWriter) Write(p []byte) (n int, err error) {
	prw.lock.Lock()
	defer prw.lock.Unlock()
	return 0, fmt.Errorf("not implemented")
}

func (prw *pageReaderWriter) Read(p []byte) (n int, err error) {
	prw.lock.RLock()
	defer prw.lock.RUnlock()
	return 0, fmt.Errorf("not implemented")
}

func (prw *pageReaderWriter) GetRecordReaderWriter() RecordReaderWriter {
	return &recordReaderWriter{
	//pageReaderWriter: prw,
	}
}

//func (rrw *recordReaderWriter) WriteRecord(r Record) error {
//	rrw.pageReaderWriter.lock.Lock()
//	defer rrw.pageReaderWriter.lock.Unlock()
//	return fmt.Errorf("not implemented")
//}

//func (rrw *recordReaderWriter) ReadRecord() (Record, error) {
//	rrw.pageReaderWriter.lock.RLock()
//	defer rrw.pageReaderWriter.lock.RUnlock()
//	return nil, fmt.Errorf("not implemented")
//}

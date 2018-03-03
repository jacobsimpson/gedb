package gedb

import (
	"database/sql"
	sqldriver "database/sql/driver"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {
	sql.Register("gedb", &gedbDriver{})
}

type gedbDriver struct{}

type header struct {
	Name         string
	MajorVersion int
	MinorVersion int
	PatchVersion int
	Length       int
}

var currentHeader = header{
	Name:         "gedb",
	MajorVersion: 0,
	MinorVersion: 0,
	PatchVersion: 1,
	Length:       13,
}

func (db *gedbDriver) Open(name string) (sqldriver.Conn, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("unable to open %q", name)
	}
	defer file.Close()

	if _, err := db.readHeader(file); isNewFile(err) {
		// It's a brand new file. Write the gedb header.
		if err := db.writeHeader(file); err != nil {
			return nil, fmt.Errorf("unable to write header in new database %q: %v", name, err)
		}
	} else if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("unable to open database %q: %v", name, err)
	}

	return &gedbConn{}, nil
}

func (db *gedbDriver) readHeader(reader io.Reader) (*header, error) {
	headerBytes := []byte{}
	data := make([]byte, 1)
	for {
		count, err := reader.Read(data)
		if err == io.EOF && count == 0 {
			return nil, &newFileError{}
		} else if err != nil {
			return nil, err
		}
		if data[0] == 10 {
			break
		} else {
			headerBytes = append(headerBytes, data...)
		}
	}
	nameVer := strings.Split(string(headerBytes), " - ")
	if len(nameVer) != 2 {
		return nil, fmt.Errorf("invalid header line, expected '<name> - <version>'")
	}
	versions := strings.Split(nameVer[1], ".")
	if len(versions) != 3 {
		return nil, fmt.Errorf("invalid header line, expected a version like '<major>.<minor>.<patch>'")
	}
	result := header{
		Name:   nameVer[0],
		Length: len(headerBytes),
	}

	if i, err := strconv.ParseInt(versions[0], 10, 64); err != nil {
		return nil, fmt.Errorf("major version should be an integer, found %q", versions[0])
	} else {
		result.MajorVersion = int(i)
	}

	if i, err := strconv.ParseInt(versions[1], 10, 64); err != nil {
		return nil, fmt.Errorf("minor version should be an integer, found %q", versions[1])
	} else {
		result.MinorVersion = int(i)
	}

	if i, err := strconv.ParseInt(versions[2], 10, 64); err != nil {
		return nil, fmt.Errorf("patch version should be an integer, found %q", versions[2])
	} else {
		result.PatchVersion = int(i)
	}

	return &result, nil
}

func (db *gedbDriver) writeHeader(writer io.Writer) error {
	headerLine := fmt.Sprintf("%s - %d.%d.%d\n",
		currentHeader.Name,
		currentHeader.MajorVersion,
		currentHeader.MinorVersion,
		currentHeader.PatchVersion)
	_, err := writer.Write([]byte(headerLine))
	return err
}

type newFileError struct{}

func (e *newFileError) Error() string {
	return "this is a new file"
}

func isNewFile(err error) bool {
	_, ok := err.(*newFileError)
	return ok
}

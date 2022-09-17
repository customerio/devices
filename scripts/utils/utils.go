package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// ReadCsvFromUrl reads a csv file from the given url and returns a 2d array of strings
func ReadCsvFromUrl(url string, delimiter rune, hasHeader, utf16le bool) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ReadCsv(resp.Body, delimiter, hasHeader, utf16le)
}

// ReadCsv reads a csv file from a reader and returns a 2d array of strings
func ReadCsv(reader io.Reader, delimiter rune, hasHeader, utf16le bool) ([][]string, error) {

	var csvReader *csv.Reader

	if utf16le {
		// Content-Type: text/csv; charset=utf-16le needs a decoder to read it in golang :lolsob:
		// this made me spend almost an entire day to figure out
		dec := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
		utf16r := reader
		utf8r := transform.NewReader(utf16r, dec)
		csvReader = csv.NewReader(utf8r)
	} else {
		csvReader = csv.NewReader(reader)
	}

	csvReader.Comma = delimiter
	csvReader.FieldsPerRecord = -1
	csvReader.LazyQuotes = true

	if hasHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, err
		}
	}

	return csvReader.ReadAll()
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func WriteJsonToFile(data interface{}, filename, prefix, indent string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent(prefix, indent)
	return enc.Encode(data)
}

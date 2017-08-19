//Package TFileGo provides API to easy work with text files in C#/Python style
package TFileGo

import (
	"errors"
	"io"
	"os"
	"strconv"
	"fmt"
)

//Constants fot access to file
const (
	F_READ   rune = 'r'
	F_WRITE  rune = 'w'
	F_APPEND rune = 'a'
)
//Constants to move current file position offset
const (
	SEEK_START   = 0
	SEEK_CURRENT = 1
	SEEK_END     = 2
)

//File object
type TFile struct {
	File *os.File
	Name string
	Mode rune
}

//Move current position in file
func (this *TFile) Seek(offset int64, whence int) (int64, error) {
	return this.File.Seek(offset, whence)
}

//Close file handler
func (this *TFile) Close() error {
	return this.File.Close()
}

//Write plain text
//Return how much runes were written and error
func (this *TFile) Write(text string) (n int, err error) {
	n, err = this.File.Write([]byte(text))
	return
}

//Write simple line
//Return how much runes were written and error
func (this *TFile) WriteLine(line string) (n int, err error) {
	n, err = this.File.Write([]byte(line + "\n"))
	return
}

//Write lines array (line by line)
func (this *TFile) WriteLines(lines []string) {
	for _, line := range lines {
		this.WriteLine(line)
	}
}

//Write string using formatting
func (this *TFile) WriteFormat(format string, args ...interface{}) {
	fmt.Fprintf(this.File, format, args...)
}

//Write any type args (Supported: string, int, float, bool, rune)
//Return error
func (this *TFile) WriteArgs(splitter string, args ...interface{}) (err error) {
	length := len(args)
	for i, arg := range args {
		switch arg.(type) {
		case int:
			this.Write(strconv.Itoa(arg.(int)))
			break
		case float64:
			this.Write(strconv.FormatFloat(arg.(float64), 'f', -1, 64))
			break
		case bool:
			this.Write(strconv.FormatBool(arg.(bool)))
			break
		case string:
			this.Write(arg.(string))
			break
		case rune:
			this.Write(string(arg.(rune)))
			break
		default:
			err = errors.New("Type unknown")
			return
		}
		if i != length-1 {
			this.Write(splitter)
		}
	}
	return
}

//Read n symbols, if n == 0 then read all text
//Return string, how much runes were really read and error
func (this *TFile) Read(n int) (string, int, error) {
	if n == 0 {
		result := ""
		n_result := 0
		for {
			buf := make([]byte, 1024)
			readed, err := this.File.Read(buf)
			result += string(buf)
			n_result += readed
			if err != nil && err != io.EOF {
				return result, n_result, err
			}
			if readed == 0 {
				return result, n_result, nil
			}
		}
		return "", 0, nil
	}
	buf := make([]byte, n)
	readed, err := this.File.Read(buf)

	return string(buf), readed, err
}

//ReadLine
//Return string and error
func (this *TFile) ReadLine() (string, error) {
	result := ""
	for {
		buf := make([]byte, 1)
		n, err := this.File.Read(buf)
		result += string(buf)
		if n == 0 || string(buf) == "\n" {
			return result, err
		}
	}
}

//Read all lines and return array
//Return string array
func (this *TFile) ReadLines() []string {
	var lines []string
	for {
		line, err := this.ReadLine()
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
	}
	return lines
}

//Open file with create mode
//Return TFile object and error
func CreateFile(filename string) (*TFile, error) {
	var file TFile
	var err error
	file.File, err = os.Create(filename)
	file.Name = filename
	file.Mode = 'w'
	return &file, err
}

//Open file with append mode
//Return TFile object and error
func AppendFile(filename string) (*TFile, error) {
	var file TFile
	var err error
	file.File, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	file.Name = filename
	file.Mode = 'a'
	return &file, err
}

//Open file with read mode
//Return TFile object and error
func ReadFile(filename string) (*TFile, error) {
	var file TFile
	var err error
	file.File, err = os.Open(filename)
	file.Name = filename
	file.Mode = 'r'
	return &file, err
}

//Open file with read, write or append access
//Return TFile object and error
func OpenFile(filename string, mode rune) (*TFile, error) {
	var err error
	switch mode {
	case 'w':
		return CreateFile(filename)
	case 'r':
		return ReadFile(filename)
	case 'a':
		return AppendFile(filename)
	default:
		err = errors.New("No such mode!")
	}
	return nil, err
}

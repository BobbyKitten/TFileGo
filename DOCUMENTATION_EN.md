# Documentation

## Opening file
There are several ways to do it:

Use wrapper on all others:
OpenFile

```go
func OpenFile(filename string, mode rune) (*TFile, error)
```
- filename is path to file
- mode is how we want to open it. Where are three constant values:
    
```go
const (
F_READ   rune = 'r'
F_WRITE  rune = 'w'
F_APPEND rune = 'a'
)
```
    
Another ways are:
CreateFile (is "F_WRITE"), AppendFile (is "F_APPEND") and ReadFile (is "F_READ")

## Working with file

The main struct to work with files is TFile:

```go
type TFile struct {
  File *os.File
  Name string
  Mode rune
}
```

Mode is how file was opened ('r' - read, 'w' - write, 'a' - append).

Methods:
  * Seek

```go
const (
  SEEK_START   = 0
  SEEK_CURRENT = 1
  SEEK_END     = 2
)
func (this *TFile) Seek(offset int64, whence int) (int64, error)
```

It's a wraper on standart Seek method.
  * Close

```go
func (this *TFile) Close() error
```

  * Write

```go
func (this *TFile) Write(text string) (n int, err error)
```
Write plain text to file.

  * WriteLine

```go
func (this *TFile) WriteLine(line string) (n int, err error)
```
Write your text + '\n'.

  * WriteLines
  
```go
func (this *TFile) WriteLines(lines []string)
```
Write array of strings to file.

  * WriteArgs

```go
func (this *TFile) WriteArgs(splitter string, args ...interface{}) (err error)
```
Can write any argument to file.
Supported types:
  (string, int, float, bool, rune)
This method concatenates all args with splitter.

  * Read

```go
func (this *TFile) Read(n int) (string, int, error)
```
Reads n symbols from file and returns readed text as string, how much symbols was really readed and error if it has been accured.
If n = 0 then method will return all text from file.

  * ReadLine

```go
func (this *TFile) ReadLine() (string, error)
```

  * ReadLines

```go
func (this *TFile) ReadLines() []string
```
Return array of all lines in file.

# Документация

## Подключение модуля

```go
import "github.com/BobbyKitten/TFileGo"
```

## Открытие файла
Для этого существует несколько способов:

Использовать обертку над всеми остальными:
OpenFile

```go
func OpenFile(filename string, mode rune) (*TFile, error)
```
- filename - это путь до самого файла
- mode - это способ открытия файла. В модуле имеются для этого 3 константы:
    
```go
const (
F_READ   rune = 'r' //Чтение
F_WRITE  rune = 'w' //Запись
F_APPEND rune = 'a' //Дозапись
)
```
    
Другими же способами являются:
CreateFile (аналог "F_WRITE"), AppendFile (аналог "F_APPEND") и ReadFile (аналог "F_READ")

## Работа с файлом

Основной структурой данных для работы с файлами является TFile:

```go
type TFile struct {
  File *os.File
  Name string
  Mode rune
}
```

Mode указывает на способ открытия файла ('r' - чтение, 'w' - запись, 'a' - дозапись).

Методы:
  * Seek

```go
const (
  SEEK_START   = 0
  SEEK_CURRENT = 1
  SEEK_END     = 2
)
func (this *TFile) Seek(offset int64, whence int) (int64, error)
```

Просто обертка над стандартным методом Seek.
  * Close

```go
func (this *TFile) Close() error
```

  * Write

```go
func (this *TFile) Write(text string) (n int, err error)
```
Записывает в файл текст таким, какой он есть.

  * WriteLine

```go
func (this *TFile) WriteLine(line string) (n int, err error)
```
Запишет вашу строку + перенос на новую строку.

  * WriteLines
  
```go
func (this *TFile) WriteLines(lines []string)
```
Используется для записи массива строк в файл.

  * WriteArgs

```go
func (this *TFile) WriteArgs(splitter string, args ...interface{}) (err error)
```
Способен записать любые аргументы известного ему типа.
Поддерживаемые типы:
  (string, int, float, bool, rune)
Все аргументы будут объединены с помощью splitter.

  * Read

```go
func (this *TFile) Read(n int) (string, int, error)
```
Прочитает n символов из файла и вернет прочитанную строку, сколько на самом деле было прочитано и ошибку если она на самом деле возникла.
Если n = 0 то метод прочитает весь текст из файла.

  * ReadLine

```go
func (this *TFile) ReadLine() (string, error)
```

  * ReadLines

```go
func (this *TFile) ReadLines() []string
```
Возвращает массив строк из файла.

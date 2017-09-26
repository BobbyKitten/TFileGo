[![Build Status](https://travis-ci.org/BobbyKitten/TFileGo.svg?branch=master)](https://travis-ci.org/BobbyKitten/TFileGo)
# TFileGo
Module for Go language wich provides methods to work with text files

Documentation:
* [English](DOCUMENTATION_EN.md)
* [Russian](DOCUMENTATION_RU.md)

install:

```
go get github.com/BobbyKitten/TFileGo
```

Import:

```go
import "github.com/BobbyKitten/TFileGo"
```

Example:

Read all lines from file:

```go
package main

import (
  "TFileGo"
  "fmt"
)
func main() {
  file, err := TFileGo.OpenFile("text.txt", TFileGo.F_READ)
  if err != nil {
    panic(err)
  }
  lines := file.ReadLines()
  for _, line := range lines {
    fmt.Print(line)
  }
}
```

# zip
golang zip directory and unzip file

### Index
```go
func File(src string, dest string) error
func Dir(srcFile string, destZip string, includeSrc bool) error
func Unzip(zipFile string, destDir string) error
```  

### Usage
```go
package main

import (
	"github.com/fuxingZhang/zip"
	"fmt"
)

func main() {
	src, zipFile, unzipDir := "./test", "./from_test.zip", "./unzip"
	err := Dir(src, zipFile, true)
	if err != nil {
		fmt.Println(err)
	}
	err = Unzip(zipFile, unzipDir)
	if err != nil {
		fmt.Println(err)
	}
}
```

### test
```bash
go test
# Show details
go test -v
# or
go test -run TestFile
go test -run TestDir
go test -run TestUnzip

go test -run ^TestFile$
go test -run ^TestFile$ zip
go test -v -run "TestFile|TestDir"
```
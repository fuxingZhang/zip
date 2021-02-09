# zip
golang zip directory and unzip file


### test
```bash
go test
# Show details
go test -v
# or
go test -run TestFile
go test -run TestDir
go test -run TestDirCarryOriginalHeader
go test -run TestUnzip

go test -run ^TestFile$
go test -run ^TestFile$ zip
go test -run ^TestFileCarryOriginalHeader$ zip
go test -v -run "TestFile|TestFileCarryOriginalHeader"
go test -v -run="TestFile|TestFileCarryOriginalHeader"
```
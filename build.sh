export GOARCH=amd64
export GOOS=linux
go build -o ddlcpad
export GOOS=windows
go build -o ddlcpad.exe

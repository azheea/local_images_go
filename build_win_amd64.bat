@ECHO off
set GOOS=windows
go env -w GOOS=windows
set GOARCH=amd64
go env -w GOARCH=amd64
go build
pause

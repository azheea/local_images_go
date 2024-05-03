@ECHO off
set GOOS=linux
go env -w GOOS=linux
set GOARCH=arm
go env -w GOARCH=arm GOARM=7
go build
pause

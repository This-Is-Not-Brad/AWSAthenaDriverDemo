$env:GOOS = "linux"
go build -o .\bin\getState .\src\main.go

serverless deploy --force
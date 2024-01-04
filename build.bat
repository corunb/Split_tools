::mac
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o  Split_tools_darwin main.go
::m1
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=arm64
go build -ldflags="-s -w" -trimpath -o  Split_tools_m1 main.go
::linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o  Split_tools_linux main.go
::win
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o  Split_tools.exe main.go
set CGO_CFLAGS=-I%VCPKG_ROOT%\installed\x64-windows\include
set CGO_LDFLAGS=-L%VCPKG_ROOT%\installed\x64-windows\lib -L%VCPKG_ROOT%\installed\x64-windows\bin
set PATH=%VCPKG_ROOT%\installed\x64-windows\bin;%PATH%
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -a -v -o main.exe

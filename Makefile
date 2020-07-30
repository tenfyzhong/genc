.PHONY: all clean

all: genc-linux genc-mac genc-win64.exe genc-win32.exe

genc-linux: *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@

genc-mac: *.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $@

genc-win64.exe: *.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $@

genc-win32.exe: *.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o $@

clean:
	$(RM) -f genc-linux genc-mac genc-win64.exe genc-win32.exe

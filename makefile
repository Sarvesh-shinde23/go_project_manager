clean:
	rm -rf *.exe
build:clean

	 go build  -o app.exe

run:build
	./app.exe



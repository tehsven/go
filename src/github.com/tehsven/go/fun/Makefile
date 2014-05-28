all: main

main: foo.o bar.c
	gcc foo.o bar.c -o main

foo.o: foo.go
	go build -gcflags -dynamiclib -dy -o libfoo.dylib foo.go

clean:
	rm -f main *.o
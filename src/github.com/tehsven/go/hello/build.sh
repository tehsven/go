#!/bin/bash
echo "cleaning"
rm libhello.dylib
rm hello*.class

echo "compiling"
gcc -dynamiclib -o libhello.dylib -dy hello.c
javac -cp .:jna-4.1.0.jar hello.java

# go tool cgo -- -dynamiclib -o libhellogo.dylib -dy hello.go

echo "executing"
java -cp .:jna-4.1.0.jar hello
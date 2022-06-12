# Call "C" library from "golang"

Call you C code from golang

## How to run

Execute `make` command, you will see below output

```
% make      
rm main 2> /dev/null || true
cd c-lib; make clean
rm ../libmylib.so 2> /dev/null || true
rm ../libmylib.a 2> /dev/null || true
cd c-lib; make static
gcc -c mylib.c
ar rc libmylib.a mylib.o
ranlib libmylib.a
mv -f libmylib.a ..
rm -f mylib.o
go run main.go
--------------START-----------------
Hello User
--------------END-----------------
# go build -o main main.go && ./main
```
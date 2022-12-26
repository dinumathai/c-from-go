# Call C library from golang

Call you C code from golang

### Prerequisite
1. Mac or Linux OS(not tried in windows)
1. gcc - GNU Compiler Collection
1. ar – create and maintain library archives 
1. ranlib - add or update the table of contents of archive libraries

### How to test ?

Execute `make build` command, you will see below output

```
% make build
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
```

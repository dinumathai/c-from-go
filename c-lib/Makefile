# Places a static or shared (dynamic) library into ../lib

all: clean dynamic

clean:
	rm ../libmylib.so 2> /dev/null || true
	rm ../libmylib.a 2> /dev/null || true

dynamic:
	gcc -fPIC -c mylib.c
	gcc -shared -o libmylib.so mylib.o
	mv -f libmylib.so ..
	rm -f mylib.o

static:
	gcc -c mylib.c
	ar rc libmylib.a mylib.o
	ranlib libmylib.a
	mv -f libmylib.a ..
	rm -f mylib.o


CC=g++


all:
	mkdir -p build
	$(CC) -o build/program src/main.cpp -Iinclude/ -Llib/ -lcom_err -lverto -lgssapi_krb5 -lk5crypto -lkrad -lkrb5 -lkrb5support -lpthread -lsqlite3

clean:
	rm -rf build/

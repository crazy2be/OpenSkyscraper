CC=g++
LDFLAGS=-lsfml-network

all:
	$(CC) $(CFLAGS) -c -o testclient.o testclient.cpp
	$(CC) $(LDFLAGS) -o testclient testclient.o

clean:
	rm -f testclient testclient.o
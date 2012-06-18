#include <SFML/Network.hpp>
#include <iostream>
#include <cstdlib>
#include <cstring>

struct PTest {
	sf::Int32 Field1;
	std::string Field2;
	char Field3[10];
};

int main() {
	sf::IPAddress local = sf::IPAddress::LocalHost;
	sf::SocketTCP client;
	if (client.Connect(1111, sf::IPAddress::LocalHost) != sf::Socket::Done) {
		std::cerr << "Unable to connect to server on localhost:1111, aborting..." << std::endl;
		exit(1);
	}
	
	sf::Packet ver;
	ver << 1;
	client.Send(ver);
	
	sf::Packet test;
	test << "Test message" << 32 << "wooooo!";
	client.Send(test);
	
	PTest ptest;
	ptest.Field1 = 3;
	ptest.Field2 = "Thing one!";
	strcpy(ptest.Field3, "Thing  2!");
	sf::Packet structTest;
	structTest.Append(&ptest, sizeof(ptest));
	client.Send(structTest);
	
	std::cout << "done." << std::endl;
}
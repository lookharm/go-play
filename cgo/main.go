package main

/*
#include <stdio.h>
void hello();

void hello() {
	printf("Hello World\n");
}
*/
import "C"

func main() {
	C.hello()
}

package main

/*
#include <stdio.h>

// C function declaration
void hello() {
    printf("Hello from C\n");
}
*/
import "C"

func main() {
  C.hello() // Call the C function
}

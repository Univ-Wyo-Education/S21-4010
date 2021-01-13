#include <stdio.h>
#include <stdlib.h>

/* Example "functional" C program */

int main ( int argc, char *argv[] ) {
	if ( argc == 3 ) {
		long a = strtol ( argv[1], NULL, 10 );
		long b = strtol ( argv[2], NULL, 10 );
		printf ( "%ld\n", a + b );
	} else {
		printf ( "Error - incorrect number of parameters, should be 2 numbers to add together\n" );
	}
}


package main

/*

#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lenvmap -Wl,-rpath=./

#include <stdio.h>
#include <stdlib.h>

#include "libenvmap.h"

void envmap();

void envmap() {
	graal_isolate_t *isolate = NULL;
	graal_isolatethread_t *thread = NULL;

	if (graal_create_isolate(NULL, &isolate, &thread) != 0) {
		fprintf(stderr, "initialization error\n");
		return;
	}

	printf("Number of entries: %d\n", filter_env(thread, "HOME"));

	graal_tear_down_isolate(thread);
}

*/
import "C"

func main() {
	C.envmap()
}

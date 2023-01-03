#include <stdio.h>
#include <stdlib.h>

#include "libenvmap.h"


void envmap(char *filter);

int main(int argc, char **argv) {
	if (argc != 2) {
		fprintf(stderr, "Usage: %s <filter>\n", argv[0]);
		exit(1);
	}
	envmap(argv[1]);
}

void envmap(char *filter) {
	 graal_isolate_t *isolate = NULL;
	 graal_isolatethread_t *thread = NULL;

	 if (graal_create_isolate(NULL, &isolate, &thread) != 0) {
		 fprintf(stderr, "initialization error\n");
		 return;
	 }

	 printf("Number of entries: %d\n", filter_env(thread, filter));

	 graal_tear_down_isolate(thread);
 }

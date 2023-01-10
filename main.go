package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lmt103parser -Wl,-rpath=./

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "libmt103parser.h"

char* mt103parser(char* incoming) {
	graal_isolate_t* isolate = NULL;
	graal_isolatethread_t* thread = NULL;

	if (graal_create_isolate(NULL, &isolate, &thread) != 0) {
		fprintf(stderr, "initialization error\n");
		return NULL;
	}

	char* result = NULL;
	result = mt103_parser(thread, incoming);

	// Copy the result from the Java / GraalVM 'mt103_parser' function since
	// 'graal_tear_down_isolate' will garbage collect the 'result'.
	// NOTE: This will be free'd by Go
	char* out = (char*)malloc(strlen(result)+1);
	strcpy(out, result);

	graal_tear_down_isolate(thread);

	return out;
}

char* mt103example() {
	graal_isolate_t* isolate = NULL;
	graal_isolatethread_t* thread = NULL;

	if (graal_create_isolate(NULL, &isolate, &thread) != 0) {
		fprintf(stderr, "initialization error\n");
		return NULL;
	}

	char* result = NULL;
	result = mt103_prowide_example(thread);

	// Copy the result from the Java / GraalVM 'mt103_parser' function since
	// 'graal_tear_down_isolate' will garbage collect the 'result'.
	// NOTE: This will be free'd by Go
	char* out = (char*)malloc(strlen(result)+1);
	strcpy(out, result);

	graal_tear_down_isolate(thread);

	return out;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Running prowide example...")
	{
		result := C.mt103example()
		defer C.free(unsafe.Pointer(result))
		fmt.Println("RESULT", C.GoString(result))
	}
	fmt.Println()
	fmt.Println("Running our example...")
	{
		incoming := C.CString(mt103Incoming2)
		defer C.free(unsafe.Pointer(incoming))

		fmt.Println("Calling mt103 parser...")
		result := C.mt103parser(incoming)
		defer C.free(unsafe.Pointer(result))
		fmt.Println("RESULT", C.GoString(result))
	}
}

var mt103Incoming = `{1:F01HSBCCN01AXXX0057000289}{2:O1030000210908HSBCCN013XXX03621145702109080000S}{3:{121:6118c3e2-1bc4-4e54-9cc2-1282561a8be5}}{4:
:20:3MBS5PG66CW1KS27
:23B:CRED
:32A:220721USD1294139
:33B:USD1294139,
:50K:4449238749283
Lotus Manufacturing Ltd.
1881 Baoan Nan Road
Shenzhen
:59:1119238749283
Banksia Minerals Pty. Ltd.
375 Monkey Street
Kalgoorlie
:71A:SHA
-}`

var mt103Incoming2 = `{1:F01BICFOOYYAXXX8683497519}{2:O1031535051028ESPBESMMAXXX54237522470510281535N}{3:{113:ROMF}{108:0510280182794665}{119:STP}}{4:
:20:0061350113089908
:13C:/RNCTIME/1534+0000
:23B:CRED
:23E:SDVA
:32A:061028EUR100000,
:33B:EUR100000,
:50K:/12345678
AGENTES DE BOLSA FOO AGENCIA
AV XXXXX 123 BIS 9 PL
12345 BARCELONA
:52A:/2337
FOOAESMMXXX
:53A:FOOAESMMXXX
:57A:BICFOOYYXXX
:59:/ES0123456789012345671234
FOO AGENTES DE BOLSA ASOC
:71A:OUR
:72:/BNF/TRANSF. BCO. FOO
-}{5:{MAC:88B4F929}{CHK:22EF370A4073}}
`

package xhyve

/*
#include <stdio.h>
#include <uuid/uuid.h>

char uuid_str[37];

extern inline char* uuidgen() {
	uuid_t uuid;

	// generate with random
	uuid_generate_random(uuid);

	// unparse (to string)
	uuid_unparse_upper(uuid, uuid_str);

	return uuid_str;
}
*/
import "C"

func uuidgen() string {
	return C.GoString(C.uuidgen())
}

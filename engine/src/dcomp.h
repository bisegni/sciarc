#ifndef CLIBRARY_H
#define CLIBRARY_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

extern void goCallbackHandler(); 

void ACFunction();
int init();
int submitFastOperation(const char *json_fast_op);
void deinit();


#ifdef __cplusplus
}
#endif

#endif
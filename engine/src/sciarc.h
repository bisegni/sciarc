#ifndef CLIBRARY_H
#define CLIBRARY_H

#ifdef __cplusplus
extern "C" {
#endif
#include <stdint.h>

extern void goCallbackHandler(char * channel_name, void *buff, int32_t buff_len);

void ACFunction();
int init();
int submitFastOperation(char *json_fast_op);
void deinit();
char* getData(const char *channel_name);
void putData(const char *channel_name, const char *channel_value);
void startMonitor(const char *channel_name);
void stopMonitor(const char *channel_name);
#ifdef __cplusplus
}
#endif

#endif
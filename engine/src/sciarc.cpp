#include "sciarc.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/thread/mutex.hpp>
#include <epics/EpicsChannel.h>
#include <map>
#include <stdint.h>
#include <thread>
#include <mutex>

std::map<std::string, std::string> channel_map;
std::mutex channel_map_mutex;

void ACFunction() {
  char name[] = "channle_name";
  int a[5] = {1,2,3,4,5};
  goCallbackHandler(name, a, 5);
}

int init() {
  EpicsChannel::init();
  return 0;
}

int submitFastOperation(char *json_fast_op) {
  int err = 0;
  return err;
}

char* getData(char *channel_name) {
    std::lock_guard guard(channel_map_mutex);
    char* out = (char*)calloc(256, sizeof(char));
    return out;
}

void deinit() {}
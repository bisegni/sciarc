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

namespace pvd = epics::pvData;

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

char* getData(const char *channel_name) {
    std::lock_guard guard(channel_map_mutex);
    std::ostringstream json;
    auto channel = std::make_unique<EpicsChannel>("pva", std::string(channel_name));
    channel->connect();
    auto value = channel->getData();
    if(!value) return nullptr;
    value->dumpValue(json);
    std::string str(json.str());
    auto out = (char*)calloc(str.size(), sizeof(char));
    strcpy(out, str.c_str());
    return out;
}

void startMonitr(const char *channel_name) {

}
void stopMonitr(const char *channel_name) {
  
}
void deinit() {}
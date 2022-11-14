#include "sciarc.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/thread/mutex.hpp>
#include <epics/EpicsChannel.h>

#include <stdint.h>
boost::mutex io_mutex;

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
  boost::mutex::scoped_lock scoped_lock(io_mutex); 
  return err;
}

void deinit() {}
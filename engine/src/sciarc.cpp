#include "sciarc.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/thread/mutex.hpp>
#include <epics/epics.h>

boost::mutex io_mutex;

void ACFunction() {
  goCallbackHandler();
}

int init() {return 0;}

int submitFastOperation(const char *json_fast_op) {
  int err = 0;
  boost::mutex::scoped_lock scoped_lock(io_mutex);

  std::unique_ptr<EpicsChannel> pc = std::make_unique<EpicsChannel>("channel_name");
  
  return err;
}

void deinit() {}
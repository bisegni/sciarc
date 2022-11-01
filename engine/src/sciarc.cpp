#include "sciarc.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/thread/mutex.hpp>

boost::mutex io_mutex;

void ACFunction() {
	printf("ACFunction\n");
  goCallbackHandler();
}

int init() {return 0;}

int submitFastOperation(const char *json_fast_op) {
  int err = 0;
  boost::mutex::scoped_lock scoped_lock(io_mutex); 
  return err;
}

void deinit() {}
#include "sciarc.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/interprocess/sync/scoped_lock.hpp>
#define CHECK_UUID(x) \
boost::mutex::scoped_lock scoped_lock(_map_mutex); \
if(map_submitted_query.find(std::string(x)) == map_submitted_query.end()) return -1; 

void ACFunction() {
	printf("ACFunction\n");
  goCallbackHandler();
}

int init() {return 0;}

int submitFastOperation(const char *json_fast_op) {
  int err = 0;

  return err;
}

void deinit() {}
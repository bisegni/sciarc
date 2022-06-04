#include "dcomp.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>

#define CHECK_UUID(x) \
boost::mutex::scoped_lock scoped_lock(_map_mutex); \
if(map_submitted_query.find(std::string(x)) == map_submitted_query.end()) return -1; 

void ACFunction() {
	printf("ACFunction\n");
  goCallbackHandler();
}

int init() {}

int submitFastOperation(const char *json_fast_op) {
  int err = 0;

  return err;
}

void deinit() {}
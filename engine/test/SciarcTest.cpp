#include <gtest/gtest.h>
#include <sciarc.h>

//simulate the golang handler
void goCallbackHandler() {
    int a = 0;
    a = a+1;
}

TEST(SciarcTest, CallHandler) {
  ACFunction();
}

TEST(SciarcTest, SubmitFastThrow) {
    EXPECT_ANY_THROW(submitFastOperation(nullptr););
}
  

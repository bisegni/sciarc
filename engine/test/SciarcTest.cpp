#include <gtest/gtest.h>
#include <sciarc.h>

//simulate the golang handler
void goCallbackHandler(char *channel_name, void *buff, int32_t buff_len) {
    int a = 0;
    a = a+1;
}

TEST(SciarcTest, CallHandler) {
  ACFunction();
}

TEST(SciarcTest, GetData) {
    auto value = getData("variable:sum");
    EXPECT_NE(value, nullptr);
    EXPECT_TRUE(strlen(value)>0);
    std::free(value);
}
  

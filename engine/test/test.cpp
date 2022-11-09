#include <gtest/gtest.h>
#include <sciarc/sciarc.h>

//simulate the golang handler
void goCallbackHandler() {
    int a = 0;
    a = a+1;
}

TEST(TestSuiteName, TestName) {
  EXPECT_EQ(submitFastOperation(nullptr), 0);
}

int main(int argc, char** argv)
{
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}

#include <gtest/gtest.h>
#include <epics/EpicsChannel.h>

TEST(EpicsTest, ChannelFault) {
    std::unique_ptr<EpicsChannel> pc;
    EXPECT_NO_THROW(pc = std::make_unique<EpicsChannel>("bacd_channel_name"));
    EXPECT_NO_THROW(pc->connect());
}

TEST(EpicsTest, ChannleOK) {
    std::unique_ptr<EpicsChannel> pc;
    epics::pvData::PVStructure::const_shared_pointer val;
    EXPECT_NO_THROW(pc = std::make_unique<EpicsChannel>("ca", "test:data:random"););
    EXPECT_NO_THROW(pc->connect());
    EXPECT_NO_THROW(val = pc->get(););
    EXPECT_NE(val, nullptr);
}
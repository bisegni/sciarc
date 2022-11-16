#include <gtest/gtest.h>
#include <epics/EpicsChannel.h>
#include <iostream>
TEST(EpicsTest, ChannelFault) {
    std::unique_ptr<EpicsChannel> pc;
    EXPECT_NO_THROW(pc = std::make_unique<EpicsChannel>("ca","bacd_channel_name"));
    EXPECT_NO_THROW(pc->connect());
    EXPECT_ANY_THROW(pc->get());
}

TEST(EpicsTest, ChannleOK) {
    std::unique_ptr<EpicsChannel> pc;
    epics::pvData::PVStructure::const_shared_pointer val;
    EXPECT_NO_THROW(pc = std::make_unique<EpicsChannel>("pva", "variable:sum"););
    EXPECT_NO_THROW(pc->connect());
    EXPECT_NO_THROW(val = pc->get(););
    EXPECT_NE(val, nullptr);
}

TEST(EpicsTest, ChannelGetSetGet) {
    std::unique_ptr<EpicsChannel> pc;
    epics::pvData::PVStructure::const_shared_pointer val;
    EXPECT_NO_THROW(pc = std::make_unique<EpicsChannel>("pva", "variable:sum"););
    EXPECT_NO_THROW(pc->connect());
    EXPECT_NO_THROW(val = pc->get(););
    EXPECT_NE(val, nullptr);
    epics::pvData::PVFieldPtrArray fields = val->getPVFields();
    for (auto&& f : fields) {
       std::string name = f->getFieldName();
       std::cout << name << std::endl;
       std::cout << val->getSubField<epics::pvData::PVDouble>("value")->get() << std::endl;
    }
    EXPECT_EQ(val->getSubField<epics::pvData::PVDouble>("value")->get(), 0);
}
#include<epics/EpicsChannel.h>

EpicsChannel::EpicsChannel(const std::string& channel_name):
channel_name(channel_name){}

void EpicsChannel::connect() {
   channel = std::make_unique<pvac::ClientChannel>(provider.connect(channel_name));
}

epics::pvData::PVStructure::const_shared_pointer
EpicsChannel::get() const {
    return channel->get();
}
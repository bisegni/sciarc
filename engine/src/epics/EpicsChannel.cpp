#include<epics/EpicsChannel.h>
#include <pv/caProvider.h>

EpicsChannel::EpicsChannel(
    const std::string& provider_name,
    const std::string& channel_name):
    channel_name(channel_name),
    provider(provider_name){}

void EpicsChannel::init() {
    // "pva" provider automatically in registry
    // add "ca" provider to registry
    epics::pvAccess::ca::CAClientFactory::start();
}

void EpicsChannel::connect() {
   channel = std::make_unique<pvac::ClientChannel>(provider.connect(channel_name));
}

epics::pvData::PVStructure::const_shared_pointer
EpicsChannel::get() const {
    return channel->get();
}
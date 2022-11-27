#ifndef EpicsChannel_H
#define EpicsChannel_H

#include <cadef.h>
#include <string>
#include <memory>
#include <pva/client.h>
#include <pv/configuration.h>
#include <pv/createRequest.h>

enum  MonitorType{
    Fail,
    Cancel,
    Disconnec,
    Data
};

using MonitorEvent = struct {
    MonitorType type;
    const std::string channel_name;
    const std::string message;
    epics::pvData::PVStructure::shared_pointer data;
};

using MonitorEventVec = std::vector<std::shared_ptr<MonitorEvent>>;
using MonitorEventVecShrdPtr = std::shared_ptr<MonitorEventVec>;

class EpicsChannel {
    const std::string channel_name;
    epics::pvData::PVStructure::shared_pointer pvReq = epics::pvData::createRequest("field()");
    epics::pvAccess::Configuration::shared_pointer conf = epics::pvAccess::ConfigurationBuilder()
                                                            .push_env()
                                                            .build();
    pvac::ClientProvider provider;
    std::unique_ptr<pvac::ClientChannel> channel;
    pvac::MonitorSync mon;
public:
explicit EpicsChannel(
    const std::string& provider_name,
    const std::string& channel_name
    );
~EpicsChannel() = default;
static void init();
static void deinit();
void connect();
epics::pvData::PVStructure::const_shared_pointer getData() const;
template<typename T >
void putData(const std::string& name, T new_value) const {
    channel->put().set(name, new_value).exec();
}
void putData(const std::string& name, const epics::pvData::AnyScalar& value) const;
void startMonitor();
MonitorEventVecShrdPtr monitor();
void stopMonitor();
};
#endif
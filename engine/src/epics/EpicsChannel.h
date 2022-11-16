#include <cadef.h>
#include <string>
#include <memory>
#include <pva/client.h>
#include <pv/configuration.h>
#include <pv/createRequest.h>

class EpicsChannel {
    const std::string channel_name;
    epics::pvData::PVStructure::shared_pointer pvReq = epics::pvData::createRequest("field()");
    epics::pvAccess::Configuration::shared_pointer conf = epics::pvAccess::ConfigurationBuilder()
                                                            .push_env()
                                                            .build();
    pvac::ClientProvider provider;
    std::unique_ptr<pvac::ClientChannel> channel;
public:
explicit EpicsChannel(
    const std::string& provider_name,
    const std::string& channel_name);
~EpicsChannel() = default;
static void init();
static void deinit();
void connect();
epics::pvData::PVStructure::const_shared_pointer get() const;
};
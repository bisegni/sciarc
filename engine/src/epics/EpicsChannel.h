#include <cadef.h>
#include <string>
#include <memory>
#include <pva/client.h>
class EpicsChannel {
    const std::string channel_name;
    pvac::ClientProvider provider;
    std::unique_ptr<pvac::ClientChannel> channel;
public:
explicit EpicsChannel(
    const std::string& provider_name,
    const std::string& channel_name);
~EpicsChannel() = default;
static void init();
void connect();
epics::pvData::PVStructure::const_shared_pointer get() const;
};
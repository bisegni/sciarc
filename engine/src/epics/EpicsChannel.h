#include <cadef.h>
#include <string>
#include <memory>
#include <pva/client.h>
class EpicsChannel {
    const std::string channel_name;
    pvac::ClientProvider provider = pvac::ClientProvider("pva");
    std::unique_ptr<pvac::ClientChannel> channel;
public:
explicit EpicsChannel(const std::string& channel_name);
~EpicsChannel() = default;
void connect();
epics::pvData::PVStructure::const_shared_pointer get() const;
};
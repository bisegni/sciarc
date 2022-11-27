#ifndef EpicsChannelMonitor_H
#define EpicsChannelMonitor_H
#include <map>
#include <thread>
#include <mutex>
#include <epics/EpicsChannel.h>

using EpicsChannelMap = std::map<std::string, std::shared_ptr<EpicsChannel>>;
using EpicsChannelMapIterator = std::map<std::string, std::shared_ptr<EpicsChannel>>::iterator;

class EpicsChannelMonitor {
    std::mutex channel_map_mutex;
    std::map<std::string, std::shared_ptr<EpicsChannel>> channel_map;
    std::unique_ptr<std::thread> scheduler_thread;
    std::function<void(const MonitorEventVecShrdPtr& event_data)> data_handler;
    EpicsChannelMapIterator cur_iter = channel_map.end();
    bool run = false;
    void task();
    void processIterator(const std::shared_ptr<EpicsChannel>& epics_channel);
public:
explicit EpicsChannelMonitor() = default;
~EpicsChannelMonitor() = default;
void start();
void stop();
void addChannel(const std::string& channel_name);
void removeChannel(const std::string& channel_name);
void setHandler(std::function<void(const MonitorEventVecShrdPtr& event_data)> new_handler);
};
#endif
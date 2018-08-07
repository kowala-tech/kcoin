#ifndef APP_RPC_PROVIDER_H
#define APP_RPC_PROVIDER_H

#include "provider.h"

class RPCProvider : public Provider
{
  public:
    RPCProvider(std::string host) : Provider(host){};

    void sendRawTransaction(std::string tx);
    int isOracle(kowala::common::address addr);
    uint256 gasEstimation(kowala::core::types::Transaction tx);
    uint256 gasPrice();
    std::string getStorageAt(std::string addr, std::string index);
    std::string blockNumber();
    int getTransactionReceipt();
};

#endif
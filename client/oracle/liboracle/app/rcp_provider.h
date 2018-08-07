#ifndef APP_RPC_PROVIDER_H
#define APP_RPC_PROVIDER_H

#include "provider.h"
#include "kowala.h"

class RPCProvider : public Provider
{
public:
    void sendRawTransaction(std::string);
    int isOracle(kowala::common::address oracle_addr) = 0;
    int gasEstimation() = 0;
    int gasPrice() = 0;
};

#endif
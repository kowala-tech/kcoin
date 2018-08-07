
#ifndef APP_PROVIDER_H
#define APP_PROVIDER_H

#include "kowala.h"
#include "uint256.h"
#include "stubclient.h"
#include <jsonrpccpp/client/connectors/httpclient.h>

using namespace jsonrpc;

class Provider
{
  public:
    Provider(std::string host)
    {
        HttpClient httpclient(host);
        StubClient rpc(httpclient);
        _client = &rpc;
    }
    virtual ~Provider(){};

    virtual void sendRawTransaction(std::string tx) = 0;
    virtual int isOracle(kowala::common::address addr) = 0;
    virtual uint256 gasEstimation(kowala::core::types::Transaction tx) = 0;
    virtual uint256 gasPrice() = 0;
    virtual std::string getStorageAt(std::string addr, std::string index) = 0;
    virtual std::string blockNumber() = 0;
    virtual int getTransactionReceipt() = 0;

  protected:
    StubClient *_client;
};

#endif
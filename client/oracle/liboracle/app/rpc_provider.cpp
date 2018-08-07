#include "rpc_provider.h"
#include "kowala.h"
#include "App.h"
#include "log.h"

// @TODO (rgeraldes) - replace with jsoncpp dep
#include <jsonrpccpp/client.h>
#include "stubclient.h"
#include <jsonrpccpp/client/connectors/httpclient.h>

#include "common.h"

void RPCProvider::sendRawTransaction(std::string tx)
{
    try
    {
        _client->eth_sendRawTransaction(tx);
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
        return;
    }
}

int RPCProvider::isOracle(kowala::common::address addr)
{
    logger->info("chegou");
    int is_oracle = 0;
    bytes data = kowala::accounts::abi::convert(ORACLE_MANAGER_IS_ORACLE, addr);
    logger->info("chegou");
    kowala::core::types::Transaction tx(ArithToUint256(1), ArithToUint256(0), ArithToUint256(0), ArithToUint256(0), kowala::common::to_address(kowala::accounts::ORACLE_MANAGER_ADDR), data);
    logger->info("chegou2");

    Json::Value json;
    json["to"] = toHexPrefixed(tx.to());
    json["data"] = toHexPrefixed(tx.data());

    // @TODO (rgeraldes) - remove
    HttpClient httpclient("http://rpcnode.zygote.kowala.tech:30503/");
    StubClient rpc(httpclient);

    try
    {
        std::string result = _client->eth_call(json, "latest");
        logger->info("result" + result);
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
        return -1;
    }

    return is_oracle;
}

uint256 RPCProvider::gasEstimation(kowala::core::types::Transaction tx)
{
    Json::Value json;
    json["to"] = toHexPrefixed(tx.to());
    json["data"] = toHexPrefixed(tx.data());

    uint256 estimation;
    try
    {
        //estimation = ArithToUint256(_client->eth_estimateGas(json));
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
    }

    return estimation;
}

uint256 RPCProvider::gasPrice()
{
    uint256 price;
    try
    {
        //price = ArithToUint256(_client->eth_gasPrice());
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
    }

    return ArithToUint256(0);
}

std::string RPCProvider::getStorageAt(std::string addr, std::string index)
{
    std::string result;
    try
    {
        // @TODO (rgeraldes)
        //result = Provider::_client->eth_getStorageAt(addr, index, "latest");

        HttpClient httpclient("http://rpcnode.zygote.kowala.tech:30503/");
        StubClient rpc(httpclient);
        logger->info("Address: " + addr + ", Index: " + index);
        result = rpc.eth_getStorageAt(addr, index, "latest");
        logger->info(result);
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
    }

    return result;
}

std::string RPCProvider::blockNumber()
{
    std::string blockNumber;
    try
    {
        blockNumber = _client->eth_blockNumber();
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
    }
}

int RPCProvider::getTransactionReceipt()
{
    /*
    try
    {
        _client->eth_getTransactionReceipt();
    }
    catch (JsonRpcException e)
    {
        logger->error(e.what());
    }
    */
    return 0;
}
#include "coinmarketcap.h"

#include "Enclave_t.h"

#include <string>

#include "https.h"
#include "picojson.h"
#include "arith_uint256.h"
#include "common.h"

static const arith_uint256 decimal_places = exp10<6>();

void CoinMarketCap::get_data(arith_uint256 &ret)
{
    arith_uint256 price(0);
    https::Client client;

    try
    {
        https::Response resp = client.get("api.coinmarketcap.com/v1/ticker/tether");

        picojson::value marketcap_data;
        std::string err = picojson::parse(marketcap_data, resp.getContent());
        if (!err.empty())
        {
            ret = 0;
        }

        if (marketcap_data.is<picojson::array>() &&
            marketcap_data.get<picojson::array>().size() == 1 &&
            marketcap_data.get<picojson::array>()[0].contains("id") &&
            marketcap_data.get<picojson::array>()[0].get("id").is<std::string>() &&
            marketcap_data.get<picojson::array>()[0].get("id").get<std::string>() == "tether" &&
            marketcap_data.get<picojson::array>()[0].contains("price_usd") &&
            marketcap_data.get<picojson::array>()[0].get("price_usd").is<std::string>())
        {
            std::string price_str = marketcap_data.get<picojson::array>()[0].get("price_usd").get<std::string>();
            double price_double = std::stod(price_str) * decimal_places.getdouble();
            price = arith_uint256(std::to_string((unsigned long)price_double)) * kUSD / decimal_places;
        }
    }
    catch (const std::exception &e)
    {
        int value;
        ocall_print_string(&value, e.what());
    }

    ret = price;
}
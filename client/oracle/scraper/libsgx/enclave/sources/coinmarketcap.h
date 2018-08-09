#ifndef ENCLAVE_SOURCES_COINMARKETCAP_H
#define ENCLAVE_SOURCES_COINMARKETCAP_H

#include "scraper.h"
#include "arith_uint256.h"

class CoinMarketCap : public Scraper
{
public:
  void get_data(arith_uint256 &ret);
};

#endif
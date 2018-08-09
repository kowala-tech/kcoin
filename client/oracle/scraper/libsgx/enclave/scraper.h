#ifndef ENCLAVE_SCRAPER_H
#define ENCLAVE_SCRAPER_H

#include "arith_uint256.h"

class Scraper
{
public:
  virtual void get_data(arith_uint256 &ret) = 0;
};

#endif
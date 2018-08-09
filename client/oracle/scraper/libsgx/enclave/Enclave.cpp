#include "Enclave.h"

#include <sgx_tseal.h>
#include <string.h>
#include <vector>
#include <string>
#include <exception>

#include "Enclave_t.h"

#include "scraper.h"
#include "sources/coinmarketcap.h"

#include "arith_uint256.h"
#include "uint256.h"
#include "common.h"
#include "kowala.h"
#include "keccak.h"
#include "ecdsa.h"

// SGX doesn't support thread creation/destruction inside an enclave

long nonce = 0;

uint256 get_currency_price()
{
    int ret;
    std::vector<std::unique_ptr<Scraper>> exchanges;
    exchanges.push_back(make_unique<CoinMarketCap>());

    arith_uint256 sum(0);
    int valid = 0;
    for (auto &scraper : exchanges)
    {
        arith_uint256 price(0);
        scraper->get_data(price);
        if (price != 0)
        {
            valid++;
            sum += price;
        }
        //ocall_log_info(&ret, "coinmarketcap.com");
    }

    // @NOTE - case for valid = 0 - all exchanges fail
    return ArithToUint256(sum / valid);
}

int ecall_currency_value_raw_tx(uint8_t *raw_tx, size_t *raw_tx_len)
{
    int ret;

    uint256 price;
    try
    {
        price = get_currency_price();
    }
    catch (const std::exception &e)
    {
        return -1;
    }

    bytes data = kowala::accounts::abi::convert(ORACLE_MANAGER_SUBMIT_SIG, price);
    kowala::core::types::Transaction tx(ArithToUint256(1), ArithToUint256(500000), ArithToUint256(3000), ArithToUint256(0), kowala::common::to_address(kowala::accounts::ORACLE_MANAGER_ADDR), data);

    ret = ecdsa_sign(tx.protected_sha3().begin(), tx.m_r.begin(), tx.m_s.begin(), &tx.m_v);
    if (ret < 0)
    {
        return -1;
    }

    bytes tx_rlp = tx.asRLP();
    memcpy(raw_tx, tx_rlp.data(), tx_rlp.size());
    *raw_tx_len = tx_rlp.size();

    return 0;
}

int ecall_register_oracle()
{
    return 0;
}

int ecall_deregister_oracle()
{
    return 0;
}
#ifndef _APP_H_
#define _APP_H_

#include <string>
#include <iostream>
#include <memory>

#ifndef FALSE
#define FALSE 0
#endif

#define TOKEN_FILENAME "enclave.token"
#define ENCLAVE_FILENAME "enclave.signed.so"
#define KEY_FILENAME "enclave.key"

const std::string ORACLE_MANAGER_IS_ORACLE = "isOracle(address)";

extern "C"
{
    int initSGX(void);
    int destroySGX(void);
    int assemblePriceTx(uint8_t*, size_t*);
}

#endif /* !_APP_H_ */

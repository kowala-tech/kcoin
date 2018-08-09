#ifndef _APP_H_
#define _APP_H_

#include <string.h>
#include <stdint.h>

#ifndef FALSE
#define FALSE 0
#endif

#define TOKEN_FILENAME "enclave.token"
#define ENCLAVE_FILENAME "enclave.signed.so"
#define KEY_FILENAME "enclave.key"

#ifdef __cplusplus
extern "C"
{
#endif
    int initSGX(void);
    int destroySGX(void);
    int assemblePriceTx(uint8_t *, size_t *);
#ifdef __cplusplus
}
#endif

#endif /* !_APP_H_ */

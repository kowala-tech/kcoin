#include <stdint.h>
#include <cstddef>

#ifndef ENCLAVE_ECDSA_H
#define ENCLAVE_ECDSA_H

int ecdsa_sign(const uint8_t *data, uint8_t *rr, uint8_t *ss, uint8_t *vv);

#endif
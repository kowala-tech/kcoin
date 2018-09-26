#ifndef ENCLAVE_UINT256_H
#define ENCLAVE_UINT256_H

#include <assert.h>
#include <cstring>
#include <stdexcept>
#include <stdint.h>
#include <string>
#include <vector>

#include "blob.h"
#include "common.h"

template <unsigned int BITS>
base_blob<BITS>::base_blob(const std::vector<unsigned char> &vch)
{
    assert(vch.size() == sizeof(data));
    memcpy(data, vch.data(), sizeof(data));
}

template <unsigned int BITS>
void base_blob<BITS>::SetHex(const std::string &str)
{
    SetHex(str.c_str());
}

// Explicit instantiations for base_blob<256>
template base_blob<256>::base_blob(const std::vector<unsigned char> &);
template void base_blob<256>::SetHex(const std::string &);

/** 256-bit opaque blob.
 * @note This type is called uint256 for historical reasons only. It is an
 * opaque blob of 256 bits and has no integer operations. Use arith_uint256 if
 * those are required.
 */
class uint256 : public base_blob<256>
{
  public:
    uint256() {}
    explicit uint256(const std::vector<unsigned char> &vch) : base_blob<256>(vch) {}
    bytes as_bytes() const {
        bytes b(std::begin(data), std::end(data));
        return b;
    }
};

/* uint256 from std::string.
 * This is a separate function because the constructor uint256(const std::string &str) can result
 * in dangerously catching uint256(0) via std::string(const char*).
 */
inline uint256 uint256S(const std::string &str)
{
    uint256 rv;
    rv.SetHex(str);
    return rv;
}

#endif // ENCLAVE_UINT256_H
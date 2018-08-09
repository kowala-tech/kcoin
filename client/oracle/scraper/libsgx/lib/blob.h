#ifndef ENCLAVE_BLOB_H
#define ENCLAVE_BLOB_H

#include <assert.h>
#include <cstring>
#include <stdexcept>
#include <stdint.h>
#include <string>
#include <vector>

/** Template base class for fixed-sized opaque blobs. */
template <unsigned int BITS>
class base_blob
{
  protected:
    static constexpr int WIDTH = BITS / 8;
    uint8_t data[WIDTH];

  public:
    base_blob()
    {
        memset(data, 0, sizeof(data));
    }

    explicit base_blob(const std::vector<unsigned char> &vch);

    bool IsNull() const
    {
        for (int i = 0; i < WIDTH; i++)
            if (data[i] != 0)
                return false;
        return true;
    }

    void SetNull()
    {
        memset(data, 0, sizeof(data));
    }

    inline int Compare(const base_blob &other) const { return memcmp(data, other.data, sizeof(data)); }

    friend inline bool operator==(const base_blob &a, const base_blob &b) { return a.Compare(b) == 0; }
    friend inline bool operator!=(const base_blob &a, const base_blob &b) { return a.Compare(b) != 0; }
    friend inline bool operator<(const base_blob &a, const base_blob &b) { return a.Compare(b) < 0; }

    void SetHex(const std::string &str);

    unsigned char *begin()
    {
        return &data[0];
    }

    unsigned char *end()
    {
        return &data[WIDTH];
    }

    const unsigned char *begin() const
    {
        return &data[0];
    }

    const unsigned char *end() const
    {
        return &data[WIDTH];
    }

    unsigned int size() const
    {
        return sizeof(data);
    }

    uint64_t GetUint64(int pos) const
    {
        const uint8_t *ptr = data + pos * 8;
        return ((uint64_t)ptr[0]) |
               ((uint64_t)ptr[1]) << 8 |
               ((uint64_t)ptr[2]) << 16 |
               ((uint64_t)ptr[3]) << 24 |
               ((uint64_t)ptr[4]) << 32 |
               ((uint64_t)ptr[5]) << 40 |
               ((uint64_t)ptr[6]) << 48 |
               ((uint64_t)ptr[7]) << 56;
    }

    template <typename Stream>
    void Serialize(Stream &s) const
    {
        s.write((char *)data, sizeof(data));
    }

    template <typename Stream>
    void Unserialize(Stream &s)
    {
        s.read((char *)data, sizeof(data));
    }
};

#endif
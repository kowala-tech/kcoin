#ifndef ENCLAVE_KOWALA_H
#define ENCLAVE_KOWALA_H

// @TOOD (rgeraldes) - rely on the blockchain info instead
#define CHAIN_ID 2

#include <string>
#include <map>

#include "keccak.h"
#include "uint256.h"
#include "arith_uint256.h"
#include "common.h"
#include "blob.h"
#include <iterator>

namespace kowala
{

namespace common
{

class address : public base_blob<160>
{
public:
  explicit address(bytes const &_b)
  {
    memcpy(data, _b.data(), std::min<unsigned>(_b.size(), 20));
  }
};

class hash : public base_blob<256>
{
public:
  explicit hash(bytes const &_b)
  {
    memcpy(data, _b.data(), std::min<unsigned>(_b.size(), 32));
  }
};

inline int fromHexChar(char _i) noexcept
{
  if (_i >= '0' && _i <= '9')
    return _i - '0';
  if (_i >= 'a' && _i <= 'f')
    return _i - 'a' + 10;
  if (_i >= 'A' && _i <= 'F')
    return _i - 'A' + 10;
  return -1;
}

inline bytes fromHex(std::string const &_s)
{
  unsigned s = (_s.size() >= 2 && _s[0] == '0' && _s[1] == 'x') ? 2 : 0;
  std::vector<uint8_t> ret;
  ret.reserve((_s.size() - s + 1) / 2);

  if (_s.size() % 2)
  {
    int h = fromHexChar(_s[s++]);
    if (h != -1)
      ret.push_back(h);
    else
      return bytes();
  }
  for (unsigned i = s; i < _s.size(); i += 2)
  {
    int h = fromHexChar(_s[i]);
    int l = fromHexChar(_s[i + 1]);
    if (h != -1 && l != -1)
      ret.push_back((byte)(h * 16 + l));
    else
      return bytes();
  }
  return ret;
}

inline address to_address(std::string const &_s)
{
  //try
  //{
  auto b = fromHex(_s.substr(0, 2) == "0x" ? _s.substr(2) : _s);
  if (b.size() == 20)
    return address(b);
  //}
  //catch (BadHexCharacter &)
  //{
  //}
  //BOOST_THROW_EXCEPTION(InvalidAddress());
}

} // namespace common

namespace accounts
{

const std::string ORACLE_MANAGER_ADDR = "0x4C55B59340FF1398d6aaE362A140D6e93855D4A5";

namespace abi
{
template <class T>
struct ABISerialiser
{
};
template <>
struct ABISerialiser<uint256>
{
  static bytes serialise(uint256 const &_t) { return _t.as_bytes(); }
};

template <>
struct ABISerialiser<kowala::common::address>
{
  static bytes serialise(kowala::common::address const &_t)
  { /*return _t.as_bytes();*/
  }
};

inline bytes abiInAux() { return {}; }
template <class T, class... U>
bytes abiInAux(T const &_t, U const &... _u)
{
  bytes current = ABISerialiser<T>::serialise(_t);
  bytes next = abiInAux(_u...);
  current.insert(current.end(), next.begin(), next.end());
  return current;
}

template <class... T>
bytes convert(std::string _id, T const &... _t)
{
  int ret = 0;
  bytes function_raw(32, 0);
  ret = keccak((const unsigned char *)_id.c_str(), _id.length(), &function_raw[0], 32);
  if (ret)
  {
    // @TODO (rgeraldes) - error
  }

  bytes params_raw = abiInAux(_t...);

  bytes abi;
  abi.reserve(function_raw.size() + params_raw.size());
  abi.insert(abi.end(), function_raw.begin(), function_raw.end());
  abi.insert(abi.end(), params_raw.begin(), params_raw.end());

  return abi;
}
} // namespace abi
} // namespace accounts

namespace rlp
{
using address = kowala::common::address;

// compute how many (non-zero) bytes there are in _i
template <typename T>
static uint8_t byte_length(T _i)
{
  uint8_t i = 0;
  for (; _i != 0; ++i, _i >>= 8)
  {
  }
  return i;
}

inline std::vector<uint8_t> itob(uint64_t num, size_t width = 0)
{
  std::vector<uint8_t> out;

  if (num == 0 && width == 0)
  {
    return out;
  }

  size_t len_len = byte_length<size_t>(num);
  for (long i = len_len - 1; i >= 0; i--)
  {
    out.push_back(static_cast<uint8_t>((num >> (8 * i)) & 0xFF));
  }

  // prepend zero until width
  if (width > out.size())
  {
    out.insert(out.begin(), width - out.size(), 0x0);
  }

  return out;
}

static const byte c_rlpMaxLengthBytes = 8;
static const byte c_rlpDataImmLenStart = 0x80;
static const byte c_rlpListStart = 0xc0;

static const byte c_rlpDataImmLenCount = c_rlpListStart - c_rlpDataImmLenStart - c_rlpMaxLengthBytes;
static const byte c_rlpDataIndLenZero = c_rlpDataImmLenStart + c_rlpDataImmLenCount - 1;
static const byte c_rlpListImmLenCount = 256 - c_rlpListStart - c_rlpMaxLengthBytes;
static const byte c_rlpListIndLenZero = c_rlpListStart + c_rlpListImmLenCount - 1;

inline uint8_t bytesRequired(uint64_t _i) { return byte_length<uint64_t>(_i); }

class RLPStream
{
public:
  bytes const &out() const
  {
    return m_out;
  }

  RLPStream &append(uint8_t _s)
  {
    std::vector<uint8_t> vec(1, 0);
    memcpy(&vec[0], &_s, sizeof(_s));
    return append(vec.begin(), vec.end());
  }

  RLPStream &append(address _s) { return append(_s.begin(), _s.end()); }

  RLPStream &append(uint256 _s)
  {
    return append(_s.begin(), _s.end());
  }
  RLPStream &append(bytes _s) { return append(_s.begin(), _s.end()); }

  template <typename Iter>
  RLPStream &append(Iter begin, Iter end)
  {
    /* NOTE (rgeraldes) - extra code due to uint256 implementation */
    std::vector<uint8_t> vec(begin, end);
    int index;
    for (index = 0; index < vec.size() && vec[index] == 0; ++index)
      ;
    vec.erase(vec.begin(), vec.begin() + index);
    /* END NOTE */

    int i = 1;

    long len = std::distance(vec.begin(), vec.end());
    if (len < 0)
      throw std::invalid_argument("String too long to be encoded.");

    int32_t len_len;
    if (len == 1 && (*vec.begin()) < 0x80)
    {
      m_out.push_back(*vec.begin());
      return *this;
    }

    // longer than 1
    if (len < 56)
    {
      m_out.push_back(0x80 + static_cast<uint8_t>(len));
      m_out.insert(m_out.end(), vec.begin(), vec.end());
    }
    else
    {
      len_len = byte_length<size_t>(len);
      if (len_len > 8)
      {
        throw std::invalid_argument("String too long to be encoded.");
      }

      m_out.push_back(0xb7 + static_cast<uint8_t>(len_len));

      std::vector<uint8_t> b_len = itob(len);
      m_out.insert(m_out.end(), b_len.begin(), b_len.end());
      m_out.insert(m_out.end(), vec.begin(), vec.end());
    }

    return *this;
  }

  template <class T>
  RLPStream &operator<<(T _data)
  {
    return append(_data);
  }

private:
  /*
  void noteAppended(size_t _itemCount = 1);

  /// Push the node-type byte (using @a _base) along with the item count @a _count.
  /// @arg _count is number of characters for strings, data-bytes for ints, or items for lists.
  void pushCount(size_t _count, byte _offset);
  */

  bytes m_out;

  std::vector<std::pair<size_t, size_t>> m_listStack;
}; // namespace rlp
} // namespace rlp

namespace core
{

namespace types
{
using address = kowala::common::address;
using hash = kowala::common::hash;
using RLPStream = kowala::rlp::RLPStream;

// Transaction represents a message-call
class Transaction
{
public:
  Transaction(uint256 nonce, uint256 gas, uint256 gas_price, uint256 value, address to, const bytes &data) : m_nonce(nonce), m_gas(gas), m_gas_price(gas_price), m_value(value), m_to(to), m_data(data) {}

  void streamRLP(RLPStream &_s, bool includeSig) const
  {
    _s.append(m_nonce);
    _s.append(m_gas_price);
    _s.append(m_gas);
    _s.append(m_to);
    _s.append(m_value);
    _s.append(m_data);
    if (includeSig)
    {
      _s.append(m_v);
      _s.append(m_r);
      _s.append(m_s);
    }
  }
  void streamProtectedRLP(RLPStream &_s) const
  {
    _s.append(m_nonce);
    _s.append(m_gas_price);
    _s.append(m_gas);
    _s.append(m_to);
    _s.append(m_value);
    _s.append(m_data);
    _s.append(ArithToUint256(2)); // chain ID
    _s.append(0);
    _s.append(0);
  }

  bytes asRLP() const
  {
    RLPStream s;
    streamRLP(s, true);
    bytes out = s.out();

    int i;
    uint8_t len_len, b;
    size_t len = out.size();
    // list header
    if (len < 56)
    {
      out.insert(out.begin(), static_cast<uint8_t>(0xc0 + len));
    }
    else
    {
      len_len = kowala::rlp::bytesRequired(len);
      if (len_len > 4)
      {
        return out;
      }
      bytes buff;
      buff.push_back(0xf7 + len_len);
      for (i = len_len - 1; i >= 0; i--)
      {
        b = (len >> (8 * i)) & 0xFF;
        buff.push_back(b);
      }
      out.insert(out.begin(), buff.begin(), buff.end());
    }

    return out;
  }

  bytes asProtectedRLP() const
  {
    RLPStream s;
    streamProtectedRLP(s);
    bytes out = s.out();

    int i;
    uint8_t len_len, b;
    size_t len = out.size();
    // list header
    if (len < 56)
    {
      out.insert(out.begin(), static_cast<uint8_t>(0xc0 + len));
    }
    else
    {
      len_len = kowala::rlp::bytesRequired(len);
      if (len_len > 4)
      {
        return out;
      }
      bytes buff;
      buff.push_back(0xf7 + len_len);
      for (i = len_len - 1; i >= 0; i--)
      {
        b = (len >> (8 * i)) & 0xFF;
        buff.push_back(b);
      }
      out.insert(out.begin(), buff.begin(), buff.end());
    }

    return out;
  }

  hash sha3() const
  {
    bytes tx_rlp = this->asRLP();
    bytes h(32, 0);
    if (keccak(tx_rlp.data(), tx_rlp.size(), &h[0], 32) < 0)
    {
      // @TODO (rgeraldes)
    }

    return hash(h);
  }

  hash protected_sha3() const
  {
    bytes tx_rlp = this->asProtectedRLP();
    bytes h(32, 0);
    if (keccak(tx_rlp.data(), tx_rlp.size(), &h[0], 32) < 0)
    {
      // @TODO (rgeraldes)
    }

    return hash(h);
  }

  uint256 nonce() const { return m_nonce; }
  address to() const { return m_to; }
  bytes data() const { return m_data; }

  uint256 m_nonce;
  uint256 m_value;
  address m_to;
  uint256 m_gas_price;
  uint256 m_gas;
  bytes m_data;
  uint256 m_r;
  uint256 m_s;
  uint8_t m_v;
};
} // namespace types
} // namespace core
} // namespace kowala

#endif /* ENCLAVE_KOWALA_H */
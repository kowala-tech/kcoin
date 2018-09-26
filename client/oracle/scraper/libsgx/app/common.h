#ifndef APP_COMMON_H
#define APP_COMMON_H

template <class Iterator>
std::string toHex(Iterator _it, Iterator _end, std::string const &_prefix)
{
    typedef std::iterator_traits<Iterator> traits;
    static_assert(sizeof(typename traits::value_type) == 1, "toHex needs byte-sized element type");

    static char const *hexdigits = "0123456789abcdef";
    size_t off = _prefix.size();
    std::string hex(std::distance(_it, _end) * 2 + off, '0');
    hex.replace(0, off, _prefix);
    for (; _it != _end; _it++)
    {
        hex[off++] = hexdigits[(*_it >> 4) & 0x0f];
        hex[off++] = hexdigits[*_it & 0x0f];
    }
    return hex;
}

/// Convert a series of bytes to the corresponding hex string with 0x prefix.
/// @example toHexPrefixed("A\x69") == "0x4169"
template <class T>
std::string toHexPrefixed(T const &_data)
{
    return toHex(_data.begin(), _data.end(), "0x");
}

#endif
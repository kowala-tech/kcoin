#ifndef ENCLAVE_H
#define ENCLAVE_H

#include <string>

const std::string ORACLE_MANAGER_SUBMIT_SIG = "submitPrice(uint256)";

template <typename T, typename... Args>
std::unique_ptr<T> make_unique(Args &&... args)
{
    return std::unique_ptr<T>(new T(std::forward<Args>(args)...));
}

#endif
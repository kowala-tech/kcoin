#ifndef ENCLAVE_H
#define ENCLAVE_H

#include <string>

const std::string ORACLE_MANAGER_SUBMIT_SIG = "addPrice(uint256)";
const std::string ORACLE_MANAGER_REGISTER_SIG = "registerOracle()";
const std::string ORACLE_MANAGER_DEREGISTER_SIG = "deregisterOracle()";
const std::string ORACLE_MANAGER_RELEASE_SIG = "releaseDeposits()";

template <typename T, typename... Args>
std::unique_ptr<T> make_unique(Args &&... args)
{
    return std::unique_ptr<T>(new T(std::forward<Args>(args)...));
}

#endif
#include "log.h"

#include "Enclave_u.h"

std::shared_ptr<spdlog::logger> logger = spdlog::stdout_color_mt("console");

int ocall_log_error(const char *str)
{
    logger->error(str);
}

int ocall_log_info(const char *str)
{
    logger->info(str);
}

int ocall_print_string(const char *str)
{
    /* Proxy/Bridge will check the length and null-terminate
   * the input string to prevent buffer overflow.
   */
    int ret = printf("%s\n", str);
    fflush(stdout);
    return ret;
}
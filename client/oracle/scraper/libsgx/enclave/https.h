#ifndef ENCLAVE_HTTPS_CLIENT_H
#define ENCLAVE_HTTPS_CLIENT_H

#include <string>

#include "mbedtls/net.h"
#include "mbedtls/ssl.h"
#include "mbedtls/entropy.h"
#include "mbedtls/ctr_drbg.h"
#include "mbedtls/debug.h"

namespace https
{
const std::string METHOD_GET = "GET";

struct URL
{
  std::string host;
  std::string path;
  std::string rawQuery;
};

URL parse(std::string raw_url);

class Request
{
private:
  const std::string method;
  const std::string port = "443";
  const URL url;

public:
  Request(const std::string method, const std::string raw_url) : method(method), url(parse(raw_url)){};

  const std::string raw();

  const std::string &get_host() const
  {
    return url.host;
  }

  const std::string &get_port() const
  {
    return port;
  }
};

class Response
{
private:
  const int status_code;
  const std::string headers;
  const std::string content;

public:
  Response() : status_code(0){};
  Response(const int status_code, const std::string content) : status_code(status_code), headers(headers), content(content){};

  const std::string &getContent() const
  {
    return content;
  }
};

class Client
{
public:
  Client();
  ~Client();

  Response get(std::string raw_url);

private:
  mbedtls_net_context server_fd;
  mbedtls_entropy_context entropy;
  mbedtls_ctr_drbg_context ctr_drbg;
  mbedtls_ssl_context ssl;
  mbedtls_ssl_config conf;
  mbedtls_x509_crt cacert;

  uint32_t flags;

  static const char *pers;

  Response exec(Request &request);
  void close();
};
}

#endif /* ENCLAVE_HTTPS_CLIENT_H */
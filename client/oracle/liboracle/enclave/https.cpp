#include "https.h"

#include "Enclave_t.h"

#include <vector>
#include <string>

#include "Enclave_t.h"
#include "certificate.h"

#include "tinyhttp/http.h"

// include tinyhttp c code
extern "C" {
#include "tinyhttp/http.c"
#include "tinyhttp/header.c"
#include "tinyhttp/chunk.c"
}

#include "mbedtls/net_sockets.h"

struct TinyHTTPResponse
{
    std::vector<char> body;
    int code;
};

static void *realloc_mem(void *opaque, void *ptr, int size)
{
    (void) opaque;
    return realloc(ptr, size);
}

static void body_handler(void *opaque, const char *data, int size)
{
    TinyHTTPResponse *response = (TinyHTTPResponse *) opaque;
    response->body.insert(response->body.end(), data, data + size);
}

static void header_handler(void *opaque, const char *ckey, int nkey, const char *cvalue, int nvalue)
{
    (void) opaque;
}

static void status_handler(void *opaque, int code)
{
    TinyHTTPResponse *response = (TinyHTTPResponse *) opaque;
    response->code = code;
}

static const http_funcs http_handlers = {
    realloc_mem,
    body_handler,
    header_handler,
    status_handler,
};

https::URL https::parse(std::string raw_url)
{
    https::URL url;

    const auto index = raw_url.find_first_of("/");
    if (std::string::npos != index)
    {
        url.host = raw_url.substr(0, index);
        url.rawQuery = raw_url.substr(index + 1);
    }
    else
    {
        url.host = raw_url;
    }

    return url;
}

const std::string https::Request::raw()
{
    std::string request_raw;
    request_raw += this->method + " /v1/ticker/tether/ " + "HTTP/1.1" + "\r\n";
    //std::string("\r\n\r\n");
    request_raw += "Accept: text/html\r\n";
    request_raw += "Host: api.coinmarketcap.com\r\n";
    request_raw += "\r\n";

    return request_raw;
}

https::Client::Client()
{
    int ret = 0;
    // Initialize the RNG and the session data
    mbedtls_net_init_ocall(&server_fd);
    mbedtls_ssl_init(&ssl);
    mbedtls_ssl_config_init(&conf);
    mbedtls_x509_crt_init(&cacert);
    mbedtls_ctr_drbg_init(&ctr_drbg);

    mbedtls_entropy_init(&entropy);
    if ((ret = mbedtls_ctr_drbg_seed(&ctr_drbg, mbedtls_entropy_func, &entropy, (const unsigned char *)pers, strlen(pers))) != 0)
    {
        throw std::runtime_error("mbedtls_ctr_drbg_seed");
    }

    // initialize certificate
    ret = mbedtls_x509_crt_parse(&cacert, (const unsigned char *)mozilla_ca_bundle, sizeof(mozilla_ca_bundle));
    if (ret < 0)
    {
        throw std::runtime_error("mbedtls_x509_crt_parse");
    }
}

https::Client::~Client()
{
    mbedtls_net_free_ocall(&server_fd);
    mbedtls_ssl_free(&ssl);
    mbedtls_ssl_config_free(&conf);
    mbedtls_ctr_drbg_free(&ctr_drbg);
    mbedtls_entropy_free(&entropy);
}

https::Response https::Client::get(std::string raw_url)
{
    https::Request req(https::METHOD_GET, raw_url);
    return exec(req);
}

static void my_debug(void *ctx, int level,
                     const char *file, int line, const char *str){
    // @TODO (rgeraldes)
    //((void)level);
    //fprintf((FILE *)ctx, "%s:%04d: %s", file, line, str);
    //fflush((FILE *)ctx);
};

https::Response https::Client::exec(https::Request &request)
{
    int ret, len;
    unsigned char buffer[4096];

    // start the connection
    if ((ret = mbedtls_net_connect_ocall(&server_fd, request.get_host().c_str(), request.get_port().c_str(), MBEDTLS_NET_PROTO_TCP)) != 0)
    {
        throw std::runtime_error("mbedtls_net_connect");
    }

    if ((ret = mbedtls_net_set_block_ocall(&server_fd) != 0)) 
    {
        throw std::runtime_error("mbedtls_net_connect");
    }

    // configure the TLS layer
    if ((ret = mbedtls_ssl_config_defaults(&conf, MBEDTLS_SSL_IS_CLIENT, MBEDTLS_SSL_TRANSPORT_STREAM, MBEDTLS_SSL_PRESET_DEFAULT)) != 0)
    {
        throw std::runtime_error("mbedtls_ssl_config_defaults");
    }

#if defined(TRACE_TLS_CLIENT)
    mbedtls_ssl_conf_verify(&conf, my_verify, NULL);
#endif

#if defined(MBEDTLS_SSL_MAX_FRAGMENT_LENGTH)
    if ((ret = mbedtls_ssl_conf_max_frag_len(&conf, MBEDTLS_SSL_MAX_FRAG_LEN_NONE)) != 0) {
        throw std::runtime_error("mbedtls_ssl_conf_max_frag_len");
    }
#endif

    mbedtls_ssl_conf_rng(&conf, mbedtls_ctr_drbg_random, &ctr_drbg);
    mbedtls_ssl_conf_dbg(&conf, my_debug, NULL);

    mbedtls_ssl_conf_read_timeout(&conf, 0);
    mbedtls_ssl_conf_ca_chain(&conf, &cacert, NULL);

#if defined(MBEDTLS_SSL_SESSION_TICKETS)
    mbedtls_ssl_conf_session_tickets(&conf, MBEDTLS_SSL_SESSION_TICKETS_ENABLED);
#endif

#if defined(MBEDTLS_SSL_RENEGOTIATION)
    mbedtls_ssl_conf_renegotiation(&conf, MBEDTLS_SSL_RENEGOTIATION_DISABLED);
#endif
    
    if ((ret = mbedtls_ssl_setup(&ssl, &conf)) != 0)
    {
        throw std::runtime_error("mbedtls_ssl_setup");
    }

#if defined(MBEDTLS_X509_CRT_PARSE_C)
    if ((ret = mbedtls_ssl_set_hostname(&ssl, request.get_host().c_str())) != 0) {
        throw std::runtime_error("mbedtls_ssl_set_hostname");
    }
#endif

    mbedtls_ssl_set_bio(&ssl, &server_fd, mbedtls_net_send_ocall, mbedtls_net_recv_ocall, mbedtls_net_recv_timeout_ocall);

    // TLS handshake
    while ((ret = mbedtls_ssl_handshake(&ssl)) != 0)
    {
        if (ret != MBEDTLS_ERR_SSL_WANT_READ && ret != MBEDTLS_ERR_SSL_WANT_WRITE)
        {
#if defined(MBEDTLS_X509_CRT_PARSE_C)
            if ((flags = mbedtls_ssl_get_verify_result(&ssl)) != 0) {
                char temp_buf[1024];
                if (mbedtls_ssl_get_peer_cert(&ssl) != NULL) {
                    mbedtls_x509_crt_info((char *) temp_buf, sizeof(temp_buf) - 1, "|-", mbedtls_ssl_get_peer_cert(&ssl));
                } else {
                    ocall_print_string(&ret, "mbedtls_ssl_get_peer_cert returns NULL");
                }
            } else {
                ocall_print_string(&ret, "X.509 Verifies");
            }
#endif /* MBEDTLS_X509_CRT_PARSE_C */
            if (ret == MBEDTLS_ERR_X509_CERT_VERIFY_FAILED) {
                ocall_print_string(&ret, "Unable to verify the server's certificate.");
            }

            throw std::runtime_error("mbedtls_ssl_handshake failed.");
        }
    }

    if ((ret = mbedtls_ssl_get_record_expansion(&ssl)) >= 0) {
        ocall_print_string(&ret, std::to_string(ret).c_str());
    } else {
        ocall_print_string(&ret, "Record expansion is [unknown (compression)]");
    }

    std::string get_request = request.raw();
    for (int written = 0, frags = 0; written < get_request.size(); written += ret, frags++) {
        while ((ret = mbedtls_ssl_write(&ssl, reinterpret_cast<const unsigned char *>(get_request.c_str()) + written, get_request.size() - written)) <= 0) {
            if (ret != MBEDTLS_ERR_SSL_WANT_READ && ret != MBEDTLS_ERR_SSL_WANT_WRITE) {
                throw std::runtime_error("mbedtls_ssl_write");
            }
        }
    }

    // Read the http response
    // https://github.com/mendsley/tinyhttp/blob/master/example.cpp
    TinyHTTPResponse tiny_resp;
    tiny_resp.code = 0;

    http_roundtripper rt;
    http_init(&rt, http_handlers, &tiny_resp);

    unsigned char *data = buffer;
    bool need_more = true;
    while (need_more) {
        const unsigned char *data = buffer;
        int n_data = mbedtls_ssl_read(&ssl, buffer, sizeof(buffer));

        if (n_data == MBEDTLS_ERR_SSL_WANT_READ || n_data == MBEDTLS_ERR_SSL_WANT_WRITE)
            continue;

        // EOF reached
        if (n_data == 0 && rt.contentlength == -1) {
            break;
        }

        if (n_data < 0) {
            ret = n_data;
            switch (n_data) {
                case MBEDTLS_ERR_SSL_PEER_CLOSE_NOTIFY:
                http_free(&rt);
                throw std::runtime_error("connection was closed gracefully");
                case MBEDTLS_ERR_NET_CONN_RESET:
                http_free(&rt);
                throw std::runtime_error("connected reset");
                default:
                http_free(&rt);
                throw std::runtime_error("mbedtls_ssl_read returned non-sense");
            }
        }
        while (need_more && n_data) {
            int read;
            need_more = (bool) http_data(&rt, (const char *) data, n_data, &read);
            n_data -= read;
            data += read;
        }
    }

    if (http_iserror(&rt))
    {
        http_free(&rt);
        throw std::runtime_error("http roundtrip error code=" + std::to_string(rt.code));
    }
    
    std::string content(tiny_resp.body.begin(), tiny_resp.body.end());
    https::Response resp(tiny_resp.code, content);

    ocall_print_string(&ret, std::to_string(tiny_resp.code).c_str());

    http_free(&rt);

    return resp;
}

void https::Client::close()
{
    int ret;
    do ret = mbedtls_ssl_close_notify(&ssl);
    while (ret == MBEDTLS_ERR_SSL_WANT_WRITE);
    ret = 0;
}

const char *https::Client::pers = "oracle-sgx";
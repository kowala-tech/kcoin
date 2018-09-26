#include "ecdsa.h"

#include <cstring>

#include "Enclave_t.h"
#include <sgx_tseal.h>

#include "mbedtls/bignum.h"
#include "mbedtls/platform.h"
#include "mbedtls/ctr_drbg.h"
#include "mbedtls/ecdsa.h"
#include "mbedtls/entropy.h"
#include "mbedtls/error.h"
#include "mbedtls/config.h"
#include "mbedtls/ecp.h"

#include "mbedtls/sha256.h"
#define SIGN_DEBUG
#undef SIGN_DEBUG

#include "keccak.h"

#define USE_NUM_NONE
#define USE_FIELD_10X26
#define USE_FIELD_INV_BUILTIN
#define USE_SCALAR_8X32
#define USE_SCALAR_INV_BUILTIN
#define NDEBUG
#include "libsecp256k1/src/secp256k1.c"
#include "libsecp256k1/src/modules/recovery/main_impl.h"
//#include "ext.h"

#include "common.h"

#define ECPARAMS MBEDTLS_ECP_DP_SECP256K1

static mbedtls_mpi global_default_secret;

secp256k1_context *ctx = secp256k1_context_create(SECP256K1_CONTEXT_SIGN | SECP256K1_CONTEXT_VERIFY);

/*
    Address:        0xD6e579085c82329C89fca7a9F012bE59028ED53F
    Public key:     04eadf899c53eaa974cda5a813d770a2e7791e686d6f4297f04e527e25d219b25c9e1be2f3fef7b4b35f8c82ce5a7f7ee3669b749e8d04d6913175dbabd6d5a58c
    Private key:    fca939d59ed3b0b69db1faffd2413ae9f6314ae2dc74a9dd2496ab7bdad066f7
*/
#define DEV_KEY "fca939d59ed3b0b69db1faffd2413ae9f6314ae2dc74a9dd2496ab7bdad066f7"

int ecdsa_sign(const uint8_t *data, uint8_t *rr, uint8_t *ss, uint8_t *vv)
{
    //C.secp256k1_context_set_illegal_callback(context, C.callbackFunc(C.secp256k1GoPanicIllegal), nil)
    //C.secp256k1_context_set_error_callback(context, C.callbackFunc(C.secp256k1GoPanicError), nil)

    int ret;
    mbedtls_mpi secret;
    mbedtls_mpi_init(&secret);
    unsigned char secret_buffer[32];

    ret = mbedtls_mpi_read_string(&secret, 16, DEV_KEY);
    if (ret != 0)
    {
        return -1;
        //ret -1;
        //goto exit;
    }

    if (mbedtls_mpi_write_binary(&secret, secret_buffer, sizeof secret_buffer) != 0)
    {
        return -1;
        //ret = -1;
        //goto exit;
    }

    secp256k1_ecdsa_recoverable_signature rsig;
    secp256k1_nonce_function nonce_fn = secp256k1_nonce_function_rfc6979;
    if (secp256k1_ecdsa_sign_recoverable(ctx, &rsig, data, secret_buffer, nonce_fn, NULL) == 0)
    {
        return -1;
    }

    int recid;
    bytes sig(65, 0);
    secp256k1_ecdsa_recoverable_signature_serialize_compact(ctx, sig.data(), &recid, &rsig);

    memcpy(rr, sig.data(), 32);
    memcpy(ss, sig.data() + 32, 32);
    *vv = recid + 35 + 2 /*chain id*/ * 2;

    return 0;
}

int ecall_sgx_select_account(const uint8_t *secret, size_t secret_len)
{
    (void)secret_len;

    uint32_t decrypted_text_length = sgx_get_encrypt_txt_len((const sgx_sealed_data_t *)secret);
    uint8_t y[decrypted_text_length];
    sgx_status_t st;

    st = sgx_unseal_data((const sgx_sealed_data_t *)secret, NULL, 0, y, &decrypted_text_length);
    if (st != SGX_SUCCESS)
    {
        return -1;
    }

    // initialize the global secret key
    mbedtls_mpi_init(&global_default_secret);
    return mbedtls_mpi_read_binary(&global_default_secret, y, sizeof y);
}

int ecdsa_pk_to_addr(const mbedtls_mpi *seckey, unsigned char *addr)
{
    unsigned char pubkey[65];
    unsigned char address[32];
    mbedtls_ecdsa_context ctx;
    size_t buflen = 0;
    int ret;

    mbedtls_ecdsa_init(&ctx);
    mbedtls_ecp_group_load(&ctx.grp, ECPARAMS);

    mbedtls_mpi_copy(&ctx.d, seckey);

    ret = mbedtls_ecp_mul(&ctx.grp, &ctx.Q, &ctx.d, &ctx.grp.G, NULL, NULL);
    if (ret != 0)
    {
        return -1;
    }

    ret = mbedtls_ecp_point_write_binary(&ctx.grp, &ctx.Q, MBEDTLS_ECP_PF_UNCOMPRESSED, &buflen, pubkey, 65);
    if (ret == MBEDTLS_ERR_ECP_BUFFER_TOO_SMALL)
    {
        return -1;
    }
    else if (ret == MBEDTLS_ERR_ECP_BAD_INPUT_DATA)
    {
        return -1;
    }
    if (buflen != 65)
    {
        //LL_CRITICAL("ecp serialization is incorrect olen=%ld", buflen);
    }

    ret = keccak(pubkey + 1, 64, address, 32);
    if (ret != 0)
    {
        return -1;
    }

    memcpy(addr, address + 12, 20);
    return 0;
}

int ecall_sgx_new_account(unsigned char *sealed_key, size_t *sealed_key_len, unsigned char *addr)
{
    int ret = 0;

    // generate a new key
    mbedtls_mpi secret;
    mbedtls_ecp_group grp;
    unsigned char secret_buffer[32];
    mbedtls_mpi_init(&secret);
    mbedtls_ecp_group_init(&grp);
    mbedtls_ecp_group_load(&grp, MBEDTLS_ECP_DP_SECP256K1);
    ret = mbedtls_mpi_read_string(&secret, 16, DEV_KEY);
    if (ret != 0)
    {
        ret - 1;
        goto exit;
    }
    //mbedtls_mpi_fill_random(&secret, grp.nbits / 8, mbedtls_sgx_drbg_random, NULL);
    if (mbedtls_mpi_write_binary(&secret, secret_buffer, sizeof secret_buffer) != 0)
    {
        ret = -1;
        goto exit;
    }
    {
        // sealing the plaintext to ciphertext.
        uint32_t len = sgx_calc_sealed_data_size(0, sizeof(secret_buffer));
        sgx_sealed_data_t *sealed_buffer = (sgx_sealed_data_t *)malloc(len);
        sgx_status_t st = sgx_seal_data(0, NULL, sizeof secret_buffer, secret_buffer, len, sealed_buffer);
        if (st != SGX_SUCCESS)
        {
            ret = -1;
            goto exit;
        }

        // output
        *sealed_key_len = len;
        memcpy(sealed_key, sealed_buffer, len);

        free(sealed_buffer);
    }

    if (ecdsa_pk_to_addr(&secret, addr) < 0)
    {
        ret = -1;
        goto exit;
    }

exit:
    mbedtls_mpi_free(&secret);
    mbedtls_ecp_group_free(&grp);
    return ret;
}

/*
static int derive_mpi(const mbedtls_ecp_group *grp, mbedtls_mpi *x,
                      const unsigned char *buf, size_t blen)
{
    int ret;
    size_t n_size = (grp->nbits + 7) / 8;
    size_t use_size = blen > n_size ? n_size : blen;

    MBEDTLS_MPI_CHK(mbedtls_mpi_read_binary(x, buf, use_size));
    if (use_size * 8 > grp->nbits)
        MBEDTLS_MPI_CHK(mbedtls_mpi_shift_r(x, use_size * 8 - grp->nbits));

    if (mbedtls_mpi_cmp_mpi(x, &grp->N) >= 0)
        MBEDTLS_MPI_CHK(mbedtls_mpi_sub_mpi(x, x, &grp->N));

cleanup:
    return (ret);
}

int mbedtls_ecdsa_sign_with_v(mbedtls_ecp_group *grp, mbedtls_mpi *r, mbedtls_mpi *s, uint8_t *v,
                              const mbedtls_mpi *d, const unsigned char *buf, size_t blen,
                              int (*f_rng)(void *, unsigned char *, size_t), void *p_rng)
{
    int ret, key_tries, sign_tries, blind_tries;
    mbedtls_ecp_point R;
    mbedtls_mpi k, e, t, vv;

    if (grp && grp->N.p == NULL)
        return (MBEDTLS_ERR_ECP_BAD_INPUT_DATA);

    mbedtls_ecp_point_init(&R);
    mbedtls_mpi_init(&k);
    mbedtls_mpi_init(&e);
    mbedtls_mpi_init(&t);
    mbedtls_mpi_init(&vv);

    mbedtls_mpi tmp;
    mbedtls_mpi halfN;
    mbedtls_mpi_init(&tmp);
    mbedtls_mpi_init(&halfN);
    //mbedtls_mpi_read_string(&SECP256K1_N, 16, S_SECP256K1_N);
    //mbedtls_mpi_read_string(&SECP256K1_N_H, 16, S_SECP256K1_N_H);
    mbedtls_mpi_div_int(&halfN, &tmp, &grp->N, 2);

    sign_tries = 0;
    do
    {
        key_tries = 0;
        do
        {
            MBEDTLS_MPI_CHK(mbedtls_ecp_gen_keypair(grp, &k, &R, f_rng, p_rng));
            MBEDTLS_MPI_CHK(mbedtls_mpi_mod_mpi(&vv, &R.Y, &grp->N));
            MBEDTLS_MPI_CHK(mbedtls_mpi_mod_int((mbedtls_mpi_uint *)v, &vv, 2));

            if (mbedtls_mpi_cmp_abs(&R.X, &grp->N) >= 0)
            {
                *v |= 2;
            }
            MBEDTLS_MPI_CHK(mbedtls_mpi_mod_mpi(r, &R.X, &grp->N));
            //*v += 27;

            if (key_tries++ > 10)
            {
                ret = MBEDTLS_ERR_ECP_RANDOM_FAILED;
                goto cleanup;
            }
        } while (mbedtls_mpi_cmp_int(r, 0) == 0);

        MBEDTLS_MPI_CHK(derive_mpi(grp, &e, buf, blen));

        blind_tries = 0;
        do
        {
            size_t n_size = (grp->nbits + 7) / 8;
            MBEDTLS_MPI_CHK(mbedtls_mpi_fill_random(&t, n_size, f_rng, p_rng));
            MBEDTLS_MPI_CHK(mbedtls_mpi_shift_r(&t, 8 * n_size - grp->nbits));

            
            if (++blind_tries > 30)
                return (MBEDTLS_ERR_ECP_RANDOM_FAILED);
        } while (mbedtls_mpi_cmp_int(&t, 1) < 0 ||
                 mbedtls_mpi_cmp_mpi(&t, &grp->N) >= 0);

        MBEDTLS_MPI_CHK(mbedtls_mpi_mul_mpi(s, r, d));
        MBEDTLS_MPI_CHK(mbedtls_mpi_add_mpi(&e, &e, s));
        MBEDTLS_MPI_CHK(mbedtls_mpi_mul_mpi(&e, &e, &t));
        MBEDTLS_MPI_CHK(mbedtls_mpi_mul_mpi(&k, &k, &t));
        MBEDTLS_MPI_CHK(mbedtls_mpi_inv_mod(s, &k, &grp->N));
        MBEDTLS_MPI_CHK(mbedtls_mpi_mul_mpi(s, s, &e));
        MBEDTLS_MPI_CHK(mbedtls_mpi_mod_mpi(s, s, &grp->N));

        if (sign_tries++ > 10)
        {
            ret = MBEDTLS_ERR_ECP_RANDOM_FAILED;
            goto cleanup;
        }
    } while (mbedtls_mpi_cmp_int(s, 0) == 0);

    
    if (mbedtls_mpi_cmp_abs(s, &halfN) == 1)
    {
        mbedtls_mpi_sub_abs(s, &grp->N, s);
        *v ^= 1;
    }

    *v = 35 + 4; //chain_id * 2;
    //*v += 27;

cleanup:
    mbedtls_ecp_point_free(&R);
    mbedtls_mpi_free(&k);
    mbedtls_mpi_free(&e);
    mbedtls_mpi_free(&t);
    mbedtls_mpi_free(&tmp);
    mbedtls_mpi_free(&halfN);
    return (ret);
}

int ecdsa_sign(const uint8_t *data, size_t in_len, uint8_t *rr, uint8_t *ss,
               uint8_t *vv)
{
    int ret;
    mbedtls_ecdsa_context ctx_sign, ctx_verify;
    mbedtls_entropy_context entropy;
    mbedtls_ctr_drbg_context ctr_drbg;
    mbedtls_mpi r, s;

    mbedtls_mpi_init(&r);
    mbedtls_mpi_init(&s);
    mbedtls_ecdsa_init(&ctx_sign);
    mbedtls_ecdsa_init(&ctx_verify);
    mbedtls_ctr_drbg_init(&ctr_drbg);

    mbedtls_ecp_group_load(&ctx_sign.grp, ECPARAMS);

    if (global_default_secret.p == NULL)
    {
        return -1;
    }
    ret = mbedtls_mpi_copy(&ctx_sign.d, &global_default_secret);
    if (ret != 0)
    {
        return -1;
    }
    ret = mbedtls_ecp_mul(&ctx_sign.grp, &ctx_sign.Q, &ctx_sign.d,
                          &ctx_sign.grp.G, NULL, NULL);
    if (ret != 0)
    {
        return -1;
    }

    ret = mbedtls_ecdsa_sign_with_v(&ctx_sign.grp, &r, &s, vv, &ctx_sign.d, data,
                                    in_len, mbedtls_sgx_drbg_random, NULL);
    if (ret != 0)
    {
        goto exit;
    }

    mbedtls_mpi_write_binary(&r, rr, 32);
    mbedtls_mpi_write_binary(&s, ss, 32);

    ret = mbedtls_ecdsa_verify(&ctx_sign.grp, data, in_len, &ctx_sign.Q, &r, &s);
    if (ret != 0)
    {
        goto exit;
    }
    else
    {
    }

exit:
    if (ret != 0)
    {
        char error_buf[100];
        mbedtls_strerror(ret, error_buf, 100);
        // @TODO (rgeraldes) - logs
    }
    mbedtls_ecdsa_free(&ctx_verify);
    mbedtls_ecdsa_free(&ctx_sign);
    mbedtls_ctr_drbg_free(&ctr_drbg);
    mbedtls_entropy_free(&entropy);
    mbedtls_mpi_free(&r);
    mbedtls_mpi_free(&s);
    return (ret);
}

*/
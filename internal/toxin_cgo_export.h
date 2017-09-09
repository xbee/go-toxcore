#ifndef _TOXIN_CGO_EXPORT_H_
#define _TOXIN_CGO_EXPORT_H_

#include "DHT.h"

#ifdef __cplusplus
extern "C" {
#endif

    extern void onGetnodesResponse(IP_Port *ip_port, uint8_t *pubkey, void *ud);

    extern void onFriendIPResponse(void *data, int32_t number, IP_Port ip_port);

#ifdef __cplusplus
}
#endif

#endif

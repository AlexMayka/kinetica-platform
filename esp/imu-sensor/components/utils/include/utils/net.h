#ifndef NET_H
#define NET_H

#ifdef __cplusplus
extern "C" {
#endif

    void get_mac_str(char *out_mac_str);
    void wifi_init_for_esp_now(void);

#ifdef __cplusplus
}
#endif

#endif // NET_H
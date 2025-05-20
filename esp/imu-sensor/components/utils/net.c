#include "utils/net.h"
#include "utils/log.h"
#include "esp_wifi.h"
#include "nvs_flash.h"

void get_mac_str(char *out_mac_str) {
    uint8_t mac[6];
    esp_wifi_get_mac(ESP_IF_WIFI_STA, mac);
    sprintf(out_mac_str, "%02x:%02x:%02x:%02x:%02x:%02x",
            mac[0], mac[1], mac[2], mac[3], mac[4], mac[5]);
}

void wifi_init_for_esp_now(void) {
    LOGI("WIFI", "Start Wi-Fi initialized in STA mode for ESP-NOW");
    ESP_ERROR_CHECK(nvs_flash_init());
    ESP_ERROR_CHECK(esp_netif_init());
    ESP_ERROR_CHECK(esp_event_loop_create_default());

    wifi_init_config_t cfg = WIFI_INIT_CONFIG_DEFAULT();
    ESP_ERROR_CHECK(esp_wifi_init(&cfg));

    ESP_ERROR_CHECK(esp_wifi_set_storage(WIFI_STORAGE_RAM));
    ESP_ERROR_CHECK(esp_wifi_set_mode(WIFI_MODE_STA));
    ESP_ERROR_CHECK(esp_wifi_start());

    LOGI("WIFI", "End Wi-Fi initialized in STA mode for ESP-NOW");
}
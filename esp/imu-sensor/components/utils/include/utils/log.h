#pragma once

#ifdef __cplusplus
extern "C" {
#endif

    #include "esp_log.h"

    #define TAG_MAIN     "MAIN"
    #define TAG_BLE      "BLE"
    #define TAG_LED      "LED"
    #define TAG_ESP_NOW  "ESP_NOW"
    #define LOG_TAG_DEFAULT "APP"

    #define LOGI(tag, fmt, ...) ESP_LOGI(tag, fmt, ##__VA_ARGS__)
    #define LOGW(tag, fmt, ...) ESP_LOGW(tag, fmt, ##__VA_ARGS__)
    #define LOGE(tag, fmt, ...) ESP_LOGE(tag, fmt, ##__VA_ARGS__)
    #define LOGD(tag, fmt, ...) ESP_LOGD(tag, fmt, ##__VA_ARGS__)

    #define LOGI_DEF(fmt, ...) ESP_LOGI(LOG_TAG_DEFAULT, fmt, ##__VA_ARGS__)
    #define LOGE_DEF(fmt, ...) ESP_LOGE(LOG_TAG_DEFAULT, fmt, ##__VA_ARGS__)

    void log_init(void);

#ifdef __cplusplus
}
#endif
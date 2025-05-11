#include "utils/log.h"

void log_init(void) {
    esp_log_level_set("*", ESP_LOG_WARN);
    esp_log_level_set("BLE_INIT", ESP_LOG_INFO);
    esp_log_level_set("ESP-NOW", ESP_LOG_DEBUG);
    esp_log_level_set(LOG_TAG_DEFAULT, ESP_LOG_INFO);
}
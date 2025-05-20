#include "utils/log.h"

void log_init(void) {
    esp_log_level_set("*", ESP_LOG_WARN);
    esp_log_level_set(TAG_MAIN, ESP_LOG_INFO);
    esp_log_level_set(TAG_FMS, ESP_LOG_INFO);
    esp_log_level_set(TAG_LED, ESP_LOG_INFO);
    esp_log_level_set(TAG_ESP_NOW, ESP_LOG_INFO);
    esp_log_level_set(LOG_TAG_DEFAULT, ESP_LOG_INFO);
}
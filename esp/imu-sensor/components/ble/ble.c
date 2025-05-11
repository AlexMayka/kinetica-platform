#include <stdio.h>
#include <string.h>

#include "esp_log.h"
#include "nvs_flash.h"
#include "esp_private/panic_internal.h"
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
#include "host/ble_hs.h"
#include "nimble/nimble_port.h"
#include "services/gap/ble_svc_gap.h"

static const char *TAG_BLE = "BLE";
static esp_err_t ret;

static void on_reset(int reason) {
    ESP_LOGE(TAG_BLE, "BLE reset, reason: %d", reason);
}

void on_sync(const char *device_name) {
}

void ble_init(const char *device_name) {
    esp_log_level_set(TAG_BLE, ESP_LOG_INFO);
    ESP_LOGI(TAG_BLE, "Initializing Nimble");

    // NVS initialization for storing key pairs, IRC, and other data.
    ret = nvs_flash_init();

    // ESP_ERR_NVS_NO_FREE_PAGES     - There are no free pages in NVS, the system cannot work.
    // ESP_ERR_NVS_NEW_VERSION_FOUND - The NFS storage version is different thn expected by the current firmware
    if (ret == ESP_ERR_NVS_NO_FREE_PAGES || ret == ESP_ERR_NVS_NEW_VERSION_FOUND) {
        ESP_ERROR_CHECK(nvs_flash_erase());
        ret = nvs_flash_init();
    }
    ESP_ERROR_CHECK(ret);

    // Launching Nimble
    ESP_ERROR_CHECK(nimble_port_init());

    // GAP startup
    ble_svc_gap_init();
    ble_svc_gap_device_name_set(device_name);

    // Asynchronous startup
    ble_hs_cfg.reset_cb = on_reset; // IF ERROR
    // ble_hs_cfg.sync_cb = on_sync;   // WORK FUNC

    ESP_LOGI(TAG_BLE, "END Initializing Nimble");
}

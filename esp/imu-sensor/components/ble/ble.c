#include <stdio.h>
#include <string.h>

#include "esp_log.h"
#include "nvs_flash.h"
#include "esp_private/panic_internal.h"
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"
#include "host/ble_hs.h"
#include "nimble/nimble_port.h"
#include "nimble/nimble_port_freertos.h"
#include "services/gap/ble_svc_gap.h"
#include "utils/log.h"

static esp_err_t ret;
static char ble_device_name[29];
static bool ble_initialized = false;

void ble_host_task(void *param) {
    nimble_port_run();
    nimble_port_freertos_deinit();
    vTaskDelete(NULL);
}

static void ble_on_reset(int reason) {
    ESP_LOGE(TAG_BLE, "BLE reset, reason: %d", reason);
    panic_abort("");
}

static void ble_on_sync(void) {
    ESP_LOGI(TAG_BLE, "BLE sync");

    struct ble_gap_adv_params adv_params = {
        .conn_mode = BLE_GAP_CONN_MODE_NON,
        .disc_mode = BLE_GAP_DISC_MODE_GEN,
    };

    struct ble_hs_adv_fields fields = {0};
    fields.flags = BLE_HS_ADV_F_DISC_GEN | BLE_HS_ADV_F_BREDR_UNSUP;

    fields.name = (uint8_t *)ble_device_name;
    fields.name_len = strlen(ble_device_name);
    fields.name_is_complete = 1;

    ESP_LOGI(TAG_BLE, "Advertising name: %s (len=%d)", ble_device_name, fields.name_len);

    int rc = ble_gap_adv_set_fields(&fields);
    if (rc != 0) {
        ESP_LOGE(TAG_BLE, "ble_gap_adv_set_fields failed: %d", rc);
        return;
    }

    rc = ble_gap_adv_start(0, NULL, BLE_HS_FOREVER, &adv_params, NULL, NULL);
    if (rc != 0) {
        ESP_LOGE(TAG_BLE, "ble_gap_adv_start failed: %d", rc);
        return;
    }

    ESP_LOGI(TAG_BLE, "BLE advertising started");
}

void ble_advertise_start(const char *device_name) {
    LOGI(TAG_BLE, "Initializing Nimble");

    if (ble_initialized) {
        LOGI(TAG_BLE, "Nimble already initialized");
        return;
    }

    size_t name_len = strlen(device_name);
    if (name_len >= sizeof(ble_device_name)) {
        LOGW(TAG_BLE, "Device name too long (%d >= %d)", name_len, sizeof(ble_device_name));
        return;
    }

    strncpy(ble_device_name, device_name, sizeof(ble_device_name));
    ble_device_name[sizeof(ble_device_name) - 1] = '\0';  // гарантируем \0

    // NVS инициализация
    ret = nvs_flash_init();
    if (ret == ESP_ERR_NVS_NO_FREE_PAGES || ret == ESP_ERR_NVS_NEW_VERSION_FOUND) {
        ESP_ERROR_CHECK(nvs_flash_erase());
        ret = nvs_flash_init();
    }
    ESP_ERROR_CHECK(ret);

    // Инициализация BLE стека
    ESP_ERROR_CHECK(nimble_port_init());

    // Установка GAP имени
    ble_svc_gap_init();
    ESP_ERROR_CHECK(ble_svc_gap_device_name_set(ble_device_name));

    // Установка callback'ов
    ble_hs_cfg.reset_cb = ble_on_reset;
    ble_hs_cfg.sync_cb = ble_on_sync;

    ble_initialized = true;

    nimble_port_freertos_init(ble_host_task);
    ESP_LOGI(TAG_BLE, "END Initializing Nimble");
}
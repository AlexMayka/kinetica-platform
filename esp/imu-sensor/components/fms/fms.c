#include "fms.h"

#include "freertos/FreeRTOS.h"
#include "freertos/task.h"

#include "ble.h"
#include "esp_wifi.h"
#include "utils/log.h"
#include "utils/net.h"

static sensor_state_t fms_sensor_state;

void state_start_fms(void) {
    LOGI(TAG_FMS, "Start STATE_START");

    wifi_init_for_esp_now();

    fms_sensor_state = STATE_BLE_APP;
    LOGI(TAG_FMS, "End STATE_START");
}

void state_ble_app(void) {
    LOGI(TAG_FMS, "Start STATE_BLE_APP");

    char mac_str[18];
    get_mac_str(mac_str);

    char full_name[32];
    snprintf(full_name, sizeof(full_name), "ESP_%s", mac_str);

    LOGI(TAG_FMS, "Device name: %s", full_name);
    ble_advertise_start(full_name);

    fms_sensor_state = STATE_SEND_DATA;

    LOGI(TAG_FMS, "End STATE_BLE_APP");
}

void state_get_command(void) {
    LOGI(TAG_FMS, "Start STATE_GET_COMMAND");
    fms_sensor_state = STATE_SEND_DATA;
    LOGI(TAG_FMS, "End STATE_GET_COMMAND");
}

void state_send_data(void) {
    LOGI(TAG_FMS, "Start STATE_SEND_DATA");
    fms_sensor_state = STATE_SLEEP;
    LOGI(TAG_FMS, "End STATE_SEND_DATA");
}

void state_sleep(void) {
    LOGI(TAG_FMS, "Start STATE_SLEEP");
    // fms_sensor_state = STATE_STOP;
    LOGI(TAG_FMS, "End STATE_SLEEP");
}

typedef void (*state_func_t)(void);
static state_func_t state_handlers[] = {
    [STATE_START]       = state_start_fms,
    [STATE_BLE_APP]     = state_ble_app,
    [STATE_GET_COMMAND] = state_get_command,
    [STATE_SEND_DATA]   = state_send_data,
    [STATE_SLEEP]       = state_sleep,
    [STATE_STOP]        = NULL,
};


void loop(void *args) {
    while (true) {
        if (fms_sensor_state == STATE_STOP) {
            LOGI(TAG_FMS, "Stopping fms");
            vTaskDelete(NULL);
            return;
        }
        if (state_handlers[fms_sensor_state]) {
            state_handlers[fms_sensor_state]();
        }
        vTaskDelay(pdMS_TO_TICKS(50));
    }
}

void start_fms(void) {
    LOGI(TAG_FMS, "Starting fms");
    fms_sensor_state = STATE_START;
    xTaskCreate(loop, "fms_task", 4096, NULL, 5, NULL);
}
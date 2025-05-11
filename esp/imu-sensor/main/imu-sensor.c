#include <stdio.h>
#include "freertos/FreeRTOS.h"

#include "ble.h"
#include "utils/log.h"
#include "utils/net.h"


void app_main(void) {
    LOGI(TAG_MAIN, "Starting work");

    char mac_str[18];
    get_mac_str(mac_str);

    printf("MAC: %s\n", mac_str);
    // ble_init(mac_str);
}
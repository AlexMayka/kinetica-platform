#include <stdio.h>
#include "freertos/FreeRTOS.h"

#include "fms.h"
#include "utils/log.h"


void app_main(void) {
    LOGI(TAG_MAIN, "Starting work");
    start_fms();
    LOGI(TAG_MAIN, "END work");
}
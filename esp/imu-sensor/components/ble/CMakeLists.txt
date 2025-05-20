idf_component_register(SRCS "ble.c"
        INCLUDE_DIRS "."
        REQUIRES bt nvs_flash utils)
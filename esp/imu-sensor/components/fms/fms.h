#ifndef FMS_H
#define FMS_H

#ifdef __cplusplus
extern "C" {
#endif

    typedef enum {
        STATE_START,
        STATE_BLE_APP,
        STATE_GET_COMMAND,
        STATE_COLLECT,
        STATE_SEND_DATA,
        STATE_SLEEP,
        STATE_STOP,
    } sensor_state_t;

    void start_fms(void);

#ifdef __cplusplus
}
#endif

#endif //FMS_H

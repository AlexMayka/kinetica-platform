<div align="center">

# 🦴 Kinetica Platform

**Full-stack IoT motion capture: ESP32 sensors + Go server**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![ESP-IDF](https://img.shields.io/badge/ESP--IDF-E7352C?style=for-the-badge&logo=espressif&logoColor=white)](https://docs.espressif.com/projects/esp-idf/)
[![BLE](https://img.shields.io/badge/Bluetooth_LE-0082FC?style=for-the-badge&logo=bluetooth&logoColor=white)](https://www.bluetooth.com)

<br>

<table>
<tr>
<td align="center"><h3>📡 BLE + WiFi</h3><sub>dual transport</sub></td>
<td align="center"><h3>🎯 6/9-axis</h3><sub>IMU support</sub></td>
<td align="center"><h3>⚡ Real-time</h3><sub>data streaming</sub></td>
<td align="center"><h3>📋 Documented</h3><sub>architecture & specs</sub></td>
</tr>
</table>

</div>

---

## 💡 What It Does

A complete **motion capture pipeline** from physical sensors to data server:

1. **ESP32 firmware** reads IMU data (accelerometer + gyroscope + magnetometer)
2. **Kinetica Protocol** encodes readings into compact binary packets
3. **BLE/WiFi transport** delivers packets to the server in real-time
4. **Go server** discovers devices, receives streams, and processes data

Built for **biomechanics research, wearable sensors, and IoT prototyping**.

---

## 🏗️ System Architecture

```
┌─────────────────┐         ┌──────────────────┐
│   ESP32 Sensor   │  BLE/  │    Go Server      │
│                  │  WiFi  │                    │
│  IMU → FMS → TX ├────────►│ Discovery → RX    │
│  (C, ESP-IDF)   │        │ (Go, cross-plat)  │
└─────────────────┘         └──────────────────┘
         │                           │
         └──── Kinetica Protocol ────┘
              (binary, CRC, TLV)
```

### ESP32 Firmware (`esp/`)

| Component | Description |
|-----------|-------------|
| **IMU driver** | Reads accelerometer + gyroscope sensors |
| **BLE GATT** | Bluetooth Low Energy server for data transmission |
| **WiFi TX** | Alternative high-throughput channel |
| **State machine** | Init → calibration → streaming lifecycle |

### Go Server (`go-server/`)

| Component | Description |
|-----------|-------------|
| **BLE scanner** | Auto-discovers BLE sensors |
| **WiFi discovery** | Network device detection (macOS) |
| **Config** | JSON settings with validation + tests |
| **CLI** | Interactive console with spinner & banners |

---

## 📁 Project Structure

```
esp/imu-sensor/
├── main/                   → Firmware entry point
└── components/
    ├── ble/                → BLE GATT server (C)
    ├── fms/                → Finite state machine (C)
    └── utils/              → Logging, networking

go-server/
├── cmd/                    → Server entry point
├── config/                 → Settings + tests
├── internal/
│   ├── device/ble/         → BLE device management
│   ├── device/wifi/        → WiFi device discovery
│   ├── console/            → CLI interface
│   └── system/             → System requirements
└── utils/                  → Error types, OS helpers

docs/                       → Architecture diagrams, specs
```

---

## 🔗 Related Projects

| Project | Description |
|---------|-------------|
| **[kinetica-protocol](https://github.com/AlexMayka/kinetica-protocol)** | Binary protocol for sensor communication (TCP/UDP/Serial/BLE) |

---

## 🛠 Tech Stack

| Layer | Technology |
|-------|------------|
| Firmware | C, ESP-IDF, ESP32 |
| Server | Go |
| Transport | BLE (GATT), WiFi |
| Protocol | Kinetica Protocol (custom binary) |
| Sensors | 6-axis / 9-axis IMU |

## License

MIT

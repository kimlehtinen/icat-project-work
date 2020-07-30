#include <Kim3z_MySignals_Wifi.h>
#include <MySignals.h>
#include "Wire.h"
#include "SPI.h"

const char* WIFI_SSID = "";
const char* WIFI_PASSWORD = "";
const char* API_URL = "";
const char* API_PATH = "/api/v1/iot/blood-pressure/store";
const char* API_PORT = "3050";

Kim3z_MySignals_Wifi wifiModule;

void setup() {
  // Connect to wifi
  wifiModule.connectToWifi(WIFI_SSID, WIFI_PASSWORD);

  // Send request
  wifiModule.postUrlEncoded(API_URL, API_PATH, API_PORT, "diastolic=5&systolic=34&pulse_per_min=67");
}

void loop() {
  delay(5000);
}


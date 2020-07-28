///////////////////////// Author: Kim Lehtinen <kim.lehtinen@student.uwasa.fi> ///////////////////////////
/*
MIT License

Copyright (c) 2020 Kim Lehtinen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

#include <Kim3z_MySignals_Wifi.h>
#include <MySignals.h>
#include "Wire.h"
#include "SPI.h"

Kim3z_MySignals_Wifi::Kim3z_MySignals_Wifi()
{
  //
}

// Sends http requests to a webserver
void Kim3z_MySignals_Wifi::post(char* url, char* path, char* port)
{
  if (sendATcommand("AT+CIPMUX=1", "OK", 6000)) {
        MySignals.println("AT+CIPMUX=1: OK");
    } else {
        MySignals.println("AT+CIPMUX=1: ERROR");
    }

    int cipstart_initial_len = 23;
    char cipstart[cipstart_initial_len + strlen(url) + strlen(port)];
    sprintf(cipstart, "AT+CIPSTART=4,\"TCP\",\"%s\",%s", url, port);
    if (sendATcommand(cipstart, "OK", 6000)) {
        MySignals.println("TCP: OK");
    } else {
        MySignals.println("TCP: ERROR");
    }
    
    const char* data = "foo=bartestingjees";

    // http method
    int method_initial_len = 16;
    char method[method_initial_len + strlen(path)];
    sprintf(method, "POST %s HTTP/1.1\r\n", path);

    // http host
    int host_initial_len = 9;
    char host[host_initial_len + strlen(url)];
    sprintf(host, "Host: %s\r\n", url);

    // http content-type
    const char* contentType = "Content-Type: application/x-www-form-urlencoded\r\n";

    // http content-length
    int content_initial_len = 18;
    char content_length[content_initial_len + strlen(data)];
    sprintf(content_length, "Content-Length: %i\r\n\r\n", strlen(data));

    // total length of http request string
    int cmd_len = strlen(method) + strlen(host) + strlen(contentType) + strlen(content_length) + strlen(data);

    // join all parts of http request to create final http string
    char cmd[cmd_len];
    sprintf(cmd, "%s%s%s%s%s", method, host, contentType, content_length, data);

    char cipsend[64];
    sprintf(cipsend, "AT+CIPSEND=4,%i", cmd_len);

    if (sendATcommand(cipsend, "OK", 6000)) {
        MySignals.println("AT+CIPSEND=4: OK");
    } else {
        MySignals.println("AT+CIPSEND=4: ERROR");
    }

    if (sendATcommand(cmd, "OK", 6000)) {
        MySignals.println("CMD: OK");
    } else {
        MySignals.println("CMD: ERROR");
    }

    const char* cipclose = "AT+CIPCLOSE=4";
    if (sendATcommand(cipclose, "OK", 6000)) {
        MySignals.println("CMD CIPCLOSE: OK");
    } else {
        MySignals.println("CMD CIPCLOSE: ERROR");
    }
}
/////////////////////////////////////////////////////////////////////////////////////////////////////


///////////////////////////////////// CODE BORROWED /////////////////////////////////////////////////
// License for connectToWifi and sendATcommand, credit to original authors.
// NOTE: Original source code is modified by Kim Lehtinen <kim.lehtinen@student.uwasa.fi> for this project.
// Original source code can be found here: https://www.cooking-hacks.com/mysignals-hw-v1-ehealth-medical-biometric-iot-platform-arduino-tutorial/index.html
/*
    Copyright (C) 2017 Libelium Comunicaciones Distribuidas S.L.
   http://www.libelium.com
    By using it you accept the MySignals Terms and Conditions.
    You can find them at: http://libelium.com/legal
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
    Version:           2.0
    Design:            David Gascon
    Implementation:    Luis Martin / Victor Boria
*/
void Kim3z_MySignals_Wifi::connectToWifi(char* wifiSSID, char* wifiPassword)
{
  Serial.begin(115200);

  MySignals.begin();

    // Enable WiFi ESP8266 Power -> bit1:1
    bitSet(MySignals.expanderState, EXP_ESP8266_POWER);
    MySignals.expanderWrite(MySignals.expanderState);

    MySignals.initSensorUART();

    MySignals.enableSensorUART(WIFI_ESP8266);
    delay(1000);

    // Checks if the WiFi module is started
    int8_t answer = sendATcommand("AT", "OK", 6000);
    if (answer == 0) {
        MySignals.println("Error");
        // waits for an answer from the module
        while (answer == 0) {
            // Send AT every two seconds and wait for the answer
            answer = sendATcommand("AT", "OK", 6000);
        }
    }
    else if (answer == 1) {
        MySignals.println("WiFi succesfully working!");

        if (sendATcommand("AT+CWMODE=1", "OK", 6000)) {
            MySignals.println("CWMODE OK");
        } else {
            MySignals.println("CWMODE Error");
        }

        int cwjap_initial_len = 14;
        char cwjap[cwjap_initial_len + strlen(wifiSSID) + strlen(wifiPassword)];
        sprintf(cwjap, "AT+CWJAP=\"%s\",\"%s\"", wifiSSID, wifiPassword);

        if (sendATcommand(cwjap, "OK", 20000)) {
        MySignals.println("Connected!");
        } else {
            MySignals.println("Error");
        }
    }
}

int8_t Kim3z_MySignals_Wifi::sendATcommand(char* ATcommand, char* expected_answer1, unsigned int timeout) {

  uint8_t x = 0,  answer = 0;
  char response[500];
  unsigned long previous;

  memset(response, '\0', sizeof(response));    // Initialize the string

  delay(100);

  while ( Serial.available() > 0) Serial.read();   // Clean the input buffer

  delay(1000);
  Serial.println(ATcommand);    // Send the AT command

  x = 0;
  previous = millis();

  // this loop waits for the answer
  do {

    if (Serial.available() != 0) {
      response[x] = Serial.read();
      x++;
      // check if the desired answer is in the response of the module
      if (strstr(response, expected_answer1) != NULL) {
        answer = 1;
        MySignals.println(response);

      }
    }
    // Waits for the asnwer with time out
  }
  while ((answer == 0) && ((millis() - previous) < timeout));

  return answer;
}
///////////////////////////////////// END CODE BORROWED /////////////////////////////////////////////////

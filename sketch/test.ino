#include <temp.pb.h>


#include <DHT.h>
#include <DHT_U.h>

#include <pb_common.h>
#include <pb.h>
#include <pb_encode.h>
#include <pb_decode.h>


#include <ESP8266WiFi.h>

#define DHTPIN 5
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);



#define DEVICEID 100

const char* ssid     = "Huang";
const char* password = "25616500";
const char* addr     = "192.168.68.106";
const uint16_t port  = 10101;

WiFiClient client;
// WiFiServer server(80);
// // Variable to store thee HTTP req
// String header;

// // Auxiliar var to stroe the current output
// const int  output5 = 5;
// const int  output4 = 4;

// const long timeoutTime = 2000;

// setup WIFI and sensor
void setup() {
  Serial.begin(115200);
  delay(10);
  
  pinMode(LED_BUILTIN, OUTPUT);
  // pinMode(output5, OUTPUT);
  // pinMode(output4, OUTPUT);
  // digitalWrite(output5,LOW);
  // digitalWrite(output4,LOW);

  Serial.println();
  Serial.print("Setting up WIFI for SSID ");
  Serial.println(ssid);

  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    Serial.println("WIFI connection failed, reconnecting...");
    delay(500);
  }

  Serial.println("");
  Serial.print("WiFi connected, ");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  Serial.println("Starting DHT11 sensor...");
  dht.begin();

  // Serial.println("Starting ESP server...");
  // server.begin();
}


void loop() {
  digitalWrite(LED_BUILTIN, LOW);
  Serial.print("connecting to ");
  Serial.println(addr);

  if (!client.connect(addr, port)) {
    Serial.println("connection failed");
    Serial.println("wait 5 sec to reconnect...");
    delay(5000);
    return;
  }

  Serial.println("reading humidity/temp...");

  float hum = dht.readHumidity();
  float tmp = dht.readTemperature();


  if (isnan(hum)||isnan(tmp))
  {
    Serial.println("Failed to read sensor data");
    return;
  }
  
  float hiCel = dht.computeHeatIndex(tmp, hum, false);
  
  pb_TempEvent temp = pb_TempEvent_init_zero;
  temp.deviceId = 1;
  temp.eventId = 100;
  temp.humidity = hum;
  temp.tempCel = tmp;
  temp.heatIdxCel = hiCel;
  
  sendTemp(temp);
  digitalWrite(LED_BUILTIN, HIGH);
  
  delay(5000);
}

void sendTemp(pb_TempEvent e) {
  uint8_t buffer[128];
  pb_ostream_t stream = pb_ostream_from_buffer(buffer, sizeof(buffer));



  if (!pb_encode(&stream, pb_TempEvent_fields, &e)){
    Serial.println("failed to encode temp proto");
    Serial.println(PB_GET_ERROR(&stream));
    return;
  }

  Serial.print("sending temp...");
  Serial.println(e.tempCel);
  client.write(buffer, stream.bytes_written);
}

// void receiveData()uint8_t{
//   // Allocate space for the decoded message 
//   MyMessage msg ={}
//   uint8_t buffer[128];
//   pb_istream_t stream;

//   client.read(buffer,128)
//   stream = pb_istream_from_buffer(buffer,count);
//   pb_decode(&stream,MyMessage_fields, &msg)

//   pb_istream_t stream =  pb_istream_from_buffer(b,)
  
  
// }


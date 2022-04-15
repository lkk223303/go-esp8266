#include <temp.pb.h>

#include <pb_common.h>
#include <pb.h>
#include <pb_encode.h>
#include <pb_decode.h>

#include <ESP8266WiFi.h>

#define DEVICEID 100
// 設定無線基地台ssid 和密碼
const char *ssid = "Huang";
const char *password = "25616500";
const char* addr     = "<server-ip-addr>";
const uint16_t port  = 10101;

WiFiClient client;

// 設定 web server port number 80
WiFiServer server(80);

// 儲存 HTTP request 的變數
String header;

String output5State = "off";
String output4State = "off";

// 指定輸出的gpio pins
const int output5 = 5;
const int output4 = 4;

// the setup function runs once when you press reset or power the board
void setup()
{
  Serial.begin(9600);
  // initialize digital pin LED_BUILTIN as an output.
  pinMode(output5, OUTPUT);
  pinMode(output4, OUTPUT);

  //   設定為低電位
  digitalWrite(output5, LOW);
  digitalWrite(output4, LOW);

  // SSID password connection
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }

  // use COM Port 列出取得 IP address
  Serial.println("");
  Serial.println("WiFi connected.");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
  server.begin();
}

// the loop function runs over and over again forever
void loop()
{
  digitalWrite(output4, HIGH);
  WiFiClient client = server.available();   // 等待 clients 連線

  Serial.print("connecting to ");
  Serial.println(addr);

  if (!client.connect(addr, port)) {
    Serial.println("connection failed");
    Serial.println("wait 5 sec to reconnect...");
    delay(5000);
    return;
  }

  pb_TempEvent temp = pb_TempEvent_init_zero;
  temp.deviceId = 12;
  temp.eventId = 100;
  temp.humidity = 0.66;
  temp.tempCel = 27;
  temp.heatIdxCel = 1;

  digitalWrite(output4, LOW);
  sendTemp(temp);
  delay(3000);


  // if (client) {                             // 假使新的用戶端連線
  //   Serial.println("New Client.");          // 從序列 Port印出訊息內容
  //   String currentLine = "";                // 清空這行的內容 
  //   while (client.connected()) {            // 當 client繼續連線持續執行迴圈
  //     if (client.available()) {             // 假使從 client 有讀到字元
  //       char c = client.read();             // 讀取這個字元
  //       Serial.write(c);                    // 印出這個字元在串列視窗
  //       header += c;
  //       if (c == '\n') {                    // 假使是換行符號

  //         // 假使目前的一行是空白且有一個新行，就結束 client HTTP 的要求
  //         if (currentLine.length() == 0) {
  //           // HTTP 表頭開始時，會有回應碼 response code (如： HTTP/1.1 200 OK)
  //           client.println("HTTP/1.1 200 OK");
  //           client.println("Content-type:text/html");
  //           client.println("Connection: close");
  //           client.println();
            
  //           // 將 GPIOs 開或關
  //           if (header.indexOf("GET /5/on") >= 0) {
  //             Serial.println("GPIO 5 on");
  //             output5State = "on";
  //             digitalWrite(output5, HIGH);

  //           } else if (header.indexOf("GET /5/off") >= 0) 
  //             {
  //                Serial.println("GPIO 5 off");
  //                output5State = "off";
  //                digitalWrite(output5, LOW);
  //              } else if (header.indexOf("GET /4/on") >= 0) 
  //                {
  //                  Serial.println("GPIO 4 on");
  //                  output4State = "on";
  //                  digitalWrite(output4, HIGH);
  //                } else if (header.indexOf("GET /4/off") >= 0) 
  //                  {
  //                    Serial.println("GPIO 4 off");
  //                    output4State = "off";
  //                    digitalWrite(output4, LOW);
  //                  }
            
  //           // 顯示 HTML 網頁
  //           client.println("<html>");
  //           client.println("<head>");
  //           client.println("<link rel=\"icon\" href=\"data:,\">");

  //           // 設定 on/off 按鈕的CSS
  //           client.println("<style>html { font-family: Helvetica; display: inline-block; margin: 0px auto; text-align: center;}");
  //           client.println(".button { background-color: #195B6A; border: none; color: white; padding: 16px 40px;");
  //           client.println("text-decoration: none; font-size: 30px; margin: 2px; cursor: pointer;}");
  //           client.println(".button2 {background-color: #77878A;}</style></head>");
            
  //           // 網頁表頭
  //           client.println("<body><h1>ESP8266 Web Server</html>");
            
  //           // 顯示現在GPIO 5 按鈕的狀態是 ON/OFF  
  //           client.println("<p>GPIO 5 - State " + output5State + "</p>");

  //           // 按鈕假使狀態是 off, 就要顯示 ON        
  //           if (output5State=="off") {
  //             client.println("<p><a href=\"/5/on\"><button class=\"button\">ON</button></a></p>");
  //           } else {
  //             client.println("<p><a href=\"/5/off\"><button class=\"button button2\">OFF</button></a></p>");
  //           } 
               
  //           // 顯示現在GPIO 4 按鈕的狀態是 ON/OFF  
  //           client.println("<p>GPIO 4 - State " + output4State + "</p>");

  //           // 按鈕假使狀態是 off, 就要顯示 ON      
  //           if (output4State=="off") {
  //             client.println("<p><a href=\"/4/on\"><button class=\"button\">ON</button/a></p>");
  //           } else {
  //             client.println("<p><a href=\"/4/off\"><button class=\"button button2\">OFF</button/a></p>");
  //           }
  //           client.println("</body></html>");                           
            
  //           // 使用空白行結束 HTTP回應
  //           client.println();
           
  //           break;
  //         } else {   // 假使有新的一行, 清除目前這一行
  //           currentLine = "";
  //         }
  //       } else if (c != '\r') {  // 讀取到的不是換行符號 
  //         currentLine += c;      // 增加一個字元在本行最後
  //       }
  //     }
  //   }
  //   // 清除表頭變數
  //   header = "";
  //   // 關閉連線 connection
  //   client.stop();
  //   Serial.println("Client disconnected.");
  //   Serial.println("");
  // }           // wait for a second
}

void sendTemp(pb_TempEvent e){
  uint8_t buffer[128];
  pb_ostream_t stream = pb_ostream_from_buffer(buffer,sizeof(buffer));

   if (!pb_encode(&stream, pb_TempEvent_fields, &e)){
    Serial.println("failed to encode temp proto");
    Serial.println(PB_GET_ERROR(&stream));
    return;
  }
  Serial.print("sending temp...");
  Serial.println(e.tempCel);
  client.write(buffer, stream.bytes_written);
}
# Analytics Service

The analytical service is designed to collect and store information about actions performed by users. Other services can send messages about user actions, for example, authorization, changing settings, adding an avatar. The service provides high performance and reliability in processing requests.

# Startup and Configuration

    Install the Go programming language
    Clone this repository: git clone https://github.com/yourusername/your-repo.git

    Move it to the project folder:
cd your-repo

    Set the settings:
go get github.com/gorilla/mux

    Configure the configuration service by editing the config.google file.
    Start the server:
    go run main.go

# Sending Requests

You can use tools like cURL or Postman to send requests to your analytics server.
Example of a cURL Request:
curl -X POST 'http://localhost:8080/analytics' \
--header 'X-Tantum-UserAgent: DeviceID=G1752G75-7C56-4G49-BGFA5ACBGC963471;DeviceType=iOS;OsVersion=15.5;AppVersion=4.3 (725)' \
--header 'X-Tantum-Authorization: 2daba111-1e48-4ba1-8753-2daba1119a09' \
--header 'Content-Type: application/json' \
--data-raw '{
"module" : "settings",
"type" : "alert",
"event" : "click",
"name" : "подтверждение выхода",
"data" : {"action" : "cancel"}
}'
#  Specify the port, database accesses, and logging level using command line flags
    Example: 
    go run main.go -port=9090 -db-credentials=my_db_creds -log-level=debug

# Test results and performance

You can perform performance testing using ApacheBench (ab) to evaluate the performance of your server:
ab -n 1000 -c 10 http://localhost:8080/analytics

Example.
Concurrency Level:      10
Time taken for tests:   0.041 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Requests per second:    24457.05 [#/sec] (mean)
Time per request:       0.409 [ms] (mean)
Transfer rate:          2173.43 [Kbytes/sec] received

Operation System: Ubuntu 20.04 LTS
Version Go: 1.16.5
Library: github.com/gorilla/mux v1.8.0

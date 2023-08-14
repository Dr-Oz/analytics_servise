# Аналитический Сервис

Аналитический сервис предназначен для сбора и хранения информации о действиях, совершенных пользователями. Другие сервисы могут отправлять сообщения о действиях пользователей, например, авторизация, изменение настроек, добавление аватарки. Сервис обеспечивает высокую производительность и надежность в обработке запросов.

# Запуск и Конфигурирование

    Установите язык программирования Go
    Клонируйте данный репозиторий: git clone https://github.com/yourusername/your-repo.git

    Перейдите в папку проекта:
    cd your-repo

    Установите зависимости:
    go get github.com/gorilla/mux

    Настройте конфигурацию сервиса, отредактировав файл config.go.
    Запустите сервер:
    go run main.go

# Отправка Запросов

  Вы можете использовать инструменты, такие как cURL или Postman, для отправки запросов на ваш аналитический сервер.
  Пример cURL Запроса:
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

# Результаты тестирования и производительность

Вы можете выполнить тестирование производительности с помощью ApacheBench (ab), чтобы оценить производительность вашего сервера:
ab -n 1000 -c 10 http://localhost:8080/analytics

ПРИМЕР РЕЗУЛЬТАТОВ.
Concurrency Level:      10
Time taken for tests:   0.041 seconds
Complete requests:      1000
Failed requests:        0
Non-2xx responses:      1000
Requests per second:    24457.05 [#/sec] (mean)
Time per request:       0.409 [ms] (mean)
Transfer rate:          2173.43 [Kbytes/sec] received

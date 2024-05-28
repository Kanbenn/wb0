
Краткие пояснения к коду по заданию:

* `Данные статичны, исходя из этого подумайте насчет модели хранения в кэше и в PostgreSQL.`
Задействовал тип данных JSONB в PostgreSQL и []byte в кэше [postgres.go](https://github.com/Kanbenn/mywbgonats/blob/main/internal/storage/postgres.go#L17)

* `Подумайте как избежать проблем, связанных с тем, что в канал могут закинуть что-угодно`
json.Unmarshal поля order_id и по нему if - корректные данные пришли в канал или мусор [app.go](https://github.com/Kanbenn/mywbgonats/blob/main/internal/app/app.go#L28).

* `Чтобы проверить работает ли подписка онлайн, сделайте себе отдельный скрипт, для публикации данных в канал`
отдельный скрипт [publisher](https://github.com/Kanbenn/mywbgonats/blob/main/cmd/publisher/publisher.go#L11), с возможностью отправки отдельных json-файлов через флаг командной строки -j. 

* `Подумайте как не терять данные в случае ошибок или проблем с сервисом`
опция DurableName в подписке на канал, которая позволяет получить пропущенные сообщения при пере-подключении к серверу [subscriber.go](https://github.com/Kanbenn/mywbgonats/blob/main/internal/subscriber/subscriber.go#L33)

* `Nats-streaming разверните локально (не путать с Nats)`
Nats-streaming и PostgreSQL локально через [docker-compose](https://github.com/Kanbenn/mywbgonats/blob/main/docker-compose.yaml)

* `Покройте сервис автотестами`
  юнит-тесты для кэша [cache_test.go](https://github.com/Kanbenn/mywbgonats/blob/main/internal/storage/cache_test.go)

* `Устройте вашему сервису стресс-тест: Воспользуйтесь утилитами WRK и Vegeta `
```
vegeta.exe attack -rate 1000 -duration=30s  -targets vegeta-targets.txt | tee .\results.bin | vegeta.exe report
Requests      [total, rate, throughput]         29999, 1000.43, 1000.40
Duration      [total, attack, wait]             29.987s, 29.986s, 1.086ms
Latencies     [min, mean, 50, 90, 95, 99, max]  505.2µs, 4.873ms, 4.237ms, 9.135ms, 10.587ms, 15.407ms, 58.766ms
Bytes In      [total, mean]                     22388977, 746.32
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:29999
Error Set:
```

### Флаги командной строки
для указания недефолтного адреса веб-сервера, nats'a и postgres [config.go](https://github.com/Kanbenn/mywbgonats/blob/main/internal/config/config.go)

### Кэш с поддержкой безопасного асинхронного доступа
через мьютексы и closure-функции [withRLock](https://github.com/Kanbenn/mywbgonats/blob/main/internal/storage/cache.go#L61)

### Локальный запуск Nats и PostgreSQL в Docker
```bash
docker-compose up --build 
```

![01-edu-system-blue](./ui/static/img/Wb_L0_100x100.png)

# WB L0 Project

<details>
<summary>Требования к реализации проекта </summary>

## Задание:

<br>

Необходимо разработать демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе.

Что нужно сделать:

1. **Развернуть локально PostgreSQL**
    - Создать свою БД
    - Настроить своего пользователя
    - Создать таблицы для хранения полученных данных
2. **Разработать сервис**
    - Реализовать подключение и подписку на канал в nats-streaming
    - Полученные данные записывать в БД
    - Реализовать кэширование полученных данных в сервисе (сохранять in memory)
    - В случае падения сервиса необходимо восстанавливать кэш из БД
    - Запустить http-сервер и выдавать данные по id из кэша
3. **Разработать простейший интерфейс отображения полученных данных по id заказа**

### Советы

1. Данные статичны, исходя из этого подумайте насчет модели хранения в кэше и в PostgreSQL. Модель в файле model.json 
2. Подумайте как избежать проблем, связанных с тем, что в канал могут закинуть что-угодно 
3. Чтобы проверить работает ли подписка онлайн, сделайте себе отдельный скрипт, для публикации данных в канал
4. Подумайте как не терять данные в случае ошибок или проблем с сервисом
5. Nats-streaming разверните локально (не путать с Nats)

<details>
 <summary>Модель данных в формате JSON</summary>

 <br>

```json
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```
<br>

</details>

</details>



<details>
<summary>Архитектура и Дизайн</summary>

<br>

### Структура базы данных  

<br>

![ERD](./ui/static/img/wb_l0_db.png)

<br>
<br>

### Архитектура сервиса

<br>

![diagramm](./ui/static/img/wb_l0_diagramm.png)


</details>

## Запуск

`make up`

### Запуск скрипта

`make sender`
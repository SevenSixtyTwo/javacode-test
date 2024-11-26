# Тестовое задание

## Описание

Напишите приложение, которое реализует REST API для управления кошельками. Приложение должно принимать запросы для выполнения операций пополнения и снятия средств, а также предоставлять возможность получения баланса кошелька.

## Эндпоинты

### 1. Пополнение или снятие средств

**Метод:** `POST`  
**URL:** `api/v1/wallet`

**Депозит:**
```json
{
  "valletId": "UUID",
  "operationType": "DEPOSIT",
  "amount": 1000
}
```
**Вывод средств:**
```json
{
  "valletId": "UUID",
  "operationType": "WITHDRAW",
  "amount": 1000
}
```

### 2. Баланс кошелька

**Метод:** `GET`  
**URL:** `api/v1/wallets/{WALLET_UUID}`

**Response:**
```json
{
    "valletType": "UUID",
    "operationType": "BALANCE",
    "amount": 1000
}
```

## Дополнительные условия

* обратить особое внимание на проблемы при работе в конкурентной среде (1000 RPS по одному кошельку)
* ни один запрос не должен быть не обработан (50X error)
* приложение должно запускаться в docker контейнере, база данных тоже, вся система должна подниматься с помощью docker-compose
* необходимо покрыть приложение тестами
* переменные окружения должны ситываться из файла config.env

## Запуск

### Запуск docker compose
```bash
docker-compose -f ./docker/docker-compose.yml up -d
```

### Добавление сервера в pgAdmin

* **Name:** bank
* **Host name/address:** postgres_container
* **Port:** 5432
* **Maintenance database:** bank
* **User:** pguser 
* **Password:** 1212 
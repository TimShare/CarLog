# CarLog

CarLog — HTTP API личного журнала автомобиля: машины, обслуживание, заправки, неисправности и напоминания.

## Запуск

```bash
docker-compose up --build -d
```

Проверка API:

```bash
curl http://localhost:8080/ping
curl http://localhost:8080/health
```

Остановка:

```bash
docker-compose down
```

## Создание автомобиля

```bash
curl -X POST http://localhost:8080/cars \
  -H 'Content-Type: application/json' \
  -d '{
    "make": "Toyota",
    "model": "Camry XV70",
    "vin": "JTNB11HK5K3000001",
    "year": 2019,
    "current_mileage": 145000
  }'
```

Полный контракт API находится в [`api/openapi.yaml`](api/openapi.yaml).

## Интеграционные тесты

```bash
./scripts/test-integration.sh
```

Тесты запускаются с отдельными API и PostgreSQL. Перед запуском применяются миграции и загружается `tests/integration/fixtures/database.sql`; после завершения всё тестовое окружение удаляется.

Обновление канонических ответов:

```bash
./scripts/test-integration.sh pytest --canonize
```

После обновления необходимо проверить изменения в `tests/integration/cars/canondata`.

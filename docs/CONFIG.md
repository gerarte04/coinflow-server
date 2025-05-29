# Конфигурация Coinflow Server

## api-gateway

**Файл конфигурации**

```yaml
http:                           # хост и порт API Gateway
  host: api-gateway
  port: 8080

storage_service:                # хост и порт storage-service
  host: storage-service
  port: 5051

collection_service:             # хост и порт collection-service
  host: collection-service
  port: 50051

auth_service:                   # хост и порт auth-service
  host: auth-service
  port: 5053

security:
  access_expiration_time: 30m   # время жизни access токена
```

**Переменные окружения**

| Переменная | Значение |
| --- | --- |
| JWT_PUBLIC_KEY | Публичный ключ ed25519, закодированный в base64 |

## auth-service

**Файл конфигурации**

```yaml
auth_service:                       # хост и порт auth-service
  host: auth-service
  port: 5053

postgres:                           # данные PostgreSQL для подключения
  host: postgres
  port: 5432
  db: coinflow_pg_db
  user: admin
  password: adminpass

redis:                              # данные Redis для подключения
  host: redis
  port: 6379
  user: admin
  user_password: adminpass
  db_number: 0

jwt:
  access_expiration_time: 30m       # время жизни access токена
  refresh_expiration_time: 168h     # время жизни refresh токена
```

**Переменные окружения**

| Переменная | Значение |
| --- | --- |
| JWT_PUBLIC_KEY | Публичный ключ ed25519, закодированный в base64 |
| JWT_PRIVATE_KEY | Приватный ключ ed25519, закодированный в base64 |

## classification-service

**Переменные окружения**

| Переменная | Значение |
| --- | --- |
| CLASSIFICATION_MODEL_NAME | Названия языковой модели с Huggingface |
| GRPC_CLASSIFICATION_SERVICE_HOST | Хост classification-service |
| GRPC_CLASSIFICATION_SERVICE_POST | Порт classification-service |

## collection-service

**Файл конфигурации**

```yaml
collection_service:                 # хост и порт collection-service
  host: collection-service
  port: 50051

classification_service:             # хост и порт classification-service
  host: classification-service
  port: 50052

postgres:                           # данные PostgreSQL для подключения
  host: postgres
  port: 5432
  db: coinflow_pg_db
  user: admin
  password: adminpass

service:
  do_translate: true
```

| Поле | Значение |
| --- | --- |
| do_translate | **boolean** <br> Указывает, должен ли производиться перевод описания операции на английский язык перед определением категории таковой при помощи API переводчика, указанном в переменной TRANSLATE_API_ADDRESS. |

**Переменные окружения**

| Переменная | Значение |
| --- | --- |
| TRANSLATE_API_ADDRESS | Адрес сервиса для перевода текста на английский для последующей передачи в classification-service. Указание необходимо только в случае, если задано ```do_translate: true```. Поддерживается API Яндекс Переводчика версии v2. |
| TRANSLATE_API_KEY | Api-key для аутентификации (для обращения к API используется сервисный аккаунт Yandex Cloud).

## storage-service

**Файл конфигурации**

```yaml
storage_service:                # хост и порт storage-service
  host: storage-service
  port: 5051

collection_service:             # хост и порт collection-service
  host: collection-service
  port: 50051

postgres:                       # данные PostgreSQL для подключения
  host: postgres
  port: 5432
  db: coinflow_pg_db
  user: admin
  password: adminpass

service:
  category_chan_buffer: 128
```

| Поле | Значение |
| --- | --- |
| category_chan_buffer | **integer** <br> Указывает буфер канала, из которого читаются операции для последующей передачи в collection-service. Увеличение приводит к большей пропускной способности. |
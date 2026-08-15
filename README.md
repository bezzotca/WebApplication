# WebProject

Веб-приложение на Go, организованное по принципам чистой архитектуры.

## Структура

```
.
├── cmd/app/main.go                  # Точка входа: сборка зависимостей (DI)
├── config/                          # Загрузка конфигурации из окружения
├── internal/
│   ├── domain/                      # Сущности и бизнес-правила (ядро приложения)
│   │   └── entity/
│   ├── usecase/                     # Бизнес-логика + интерфейсы репозиториев (порты)
│   ├── infrastructure/postgres/     # Реализация репозиториев поверх Postgres (адаптеры)
│   └── delivery/http/               # Транспортный слой: handlers, router, middleware
│       ├── handler/
│       └── middleware/
├── pkg/
│   └── database/postgres/           # ИЗОЛИРОВАННОЕ ЯДРО подключения к PostgreSQL
│       ├── config.go                 # Конфиг подключения (DSN)
│       └── postgres.go               # Connect(): пул соединений, ping, таймауты
├── migrations/                      # SQL-миграции
├── deployments/docker/              # docker-compose для локального Postgres
└── api/                             # Спецификации API (OpenAPI и т.п.)
```

## Правило зависимостей

```
delivery ──► usecase ──► domain
                ▲
infrastructure ─┘
                ▲
         pkg/database/postgres
```

- `domain` ничего не импортирует из других слоёв — чистые сущности и бизнес-ошибки.
- `usecase` знает только интерфейсы (`UserRepository`), но не знает про Postgres, SQL, HTTP.
- `infrastructure/postgres` — единственное место, где встречается SQL/pgx-типы; реализует интерфейсы `usecase`.
- `delivery/http` вызывает usecase, ничего не знает о БД напрямую.
- `pkg/database/postgres` — самостоятельное, переиспользуемое ядро: только открытие/проверка/закрытие пула соединений. Не содержит бизнес-логики и не зависит от остальных слоёв — поэтому вынесено в `pkg`, а не в `internal`.

## Запуск

```bash
cp .env.example .env
docker compose -f deployments/docker/docker-compose.yml up -d
go run ./cmd/app
```

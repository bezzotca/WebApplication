# WebApplication

Минимальный веб-сервер на Go с PostgreSQL.

## Структура

```
.
├── main.go             # точка входа: подключение к БД, роуты, запуск сервера
├── handlers.go         # HTTP-хендлеры
├── db.go               # пул соединений и SQL-запросы
├── schema.sql          # схема БД
├── docker-compose.yml  # локальный Postgres
├── .env.example        # список переменных окружения
└── go.mod
```

Всё в одном пакете `main`. Разбиение на файлы — только для читаемости: внутри пакета
файлы видят друг друга без импортов, поэтому их можно свободно делить и объединять.

## Запуск

```bash
docker compose up -d
```

```bash
go mod tidy && go run .
```

## Проверка

```bash
curl localhost:8080/health
```

```bash
curl -X POST localhost:8080/users -d '{"email":"a@b.c","name":"Alice"}'
```

```bash
curl localhost:8080/users/1
```

## Эндпоинты

| Метод | Путь         | Описание             |
|-------|--------------|----------------------|
| GET   | `/health`    | проверка живости     |
| POST  | `/users`     | создать пользователя |
| GET   | `/users/{id}`| получить по id       |

# Todo API

Простое REST API для управления задачами с JWT-авторизацией.

## Стек

- Go
- PostgreSQL
- JWT (golang-jwt/jwt)
- chi (роутер)

## Запуск проекта

### 1. Поднять PostgreSQL через Docker

```bash
docker run --name my-postgres -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres:18
```

Создайте базу данных:

```bash
docker exec -it my-postgres psql -U postgres -c "CREATE DATABASE yourdbname;"
```

### 2. Настроить переменные окружения

Скопируйте `.env.example` в `.env` и заполните своими значениями:

```bash
cp .env.example .env
```

Понадобится два значения:
- `JWT_SECRET` — любая случайная строка, используется для подписи токенов
- `DB_URL` — строка подключения к вашей Postgres-базе

### 3. Запустить проект

```bash
go run main.go
```

Сервер поднимется на `http://localhost:8080`.

## Эндпоинты

| Метод  | Путь         | Описание           | Требует авторизации |
|--------|--------------|---------------------|----------------------|
| POST   | /register    | Регистрация          | Нет                  |
| POST   | /login       | Логин, выдаёт JWT     | Нет                  |
| POST   | /task        | Создать задачу        | Да                   |
| GET    | /tasks       | Список задач          | Да                   |
| GET    | /task/{id}   | Задача по id           | Да                   |
| PUT    | /task/{id}   | Обновить задачу        | Да                   |
| DELETE | /task/{id}   | Удалить задачу         | Да                   |

## Авторизация

Все защищённые эндпоинты требуют заголовок:

```
Authorization: Bearer <ваш_токен>
```

Токен выдаётся при успешном логине (`POST /login`) и действителен 24 часа.

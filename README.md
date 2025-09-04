# Gnomiki 🧙♂️

Простое веб-приложение для управления списком гномиков.

## Запуск

```bash
# Установка зависимостей
go mod tidy

# Запуск сервера
go run cmd/server/main.go
```

Откройте http://localhost:8080

## Структура

- `cmd/server/` - точка входа
- `internal/api/` - HTTP handlers
- `internal/db/` - работа с SQLite
- `internal/models/` - модели данных
- `static/` - фронтенд файлы

## API

- `GET /api/gnomiki` - получить всех гномиков
- `POST /api/gnomiki` - добавить гномика
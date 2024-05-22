# Распределенный вычислитель арифметических выражений (агент)

Финальное задание по второму спринту Яндекс Лицея (GoLang)

## Требуется установка агента

Подробнее тут: https://github.com/zhukovrost/orchestrator.git

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/zhukovrost/agentYL.git
    ```

2. Перейдите в директорию проекта:
    ```sh
    cd agentYL
    ```

3. Установите зависимости:
    ```sh
    go mod tidy
    ```

## Запуск

Для запуска сервера выполните:

```sh
go run cmd/agent/main.go
```

## Инструкция по использованию



## Структура проекта

```
agentYL/
├── cmd/
│   └── agent/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── config/
│   │   └── config.go
│   └── service/
│       └── service.go
├── pkg/
│   ├── utils/
│   │   └── utils.go
│   └── middleware/
│       └── middleware.go
├── .gitignore
├── go.mod
└── README.md
```

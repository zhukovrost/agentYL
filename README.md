# Распределенный вычислитель арифметических выражений (агент)

Финальное задание по второму спринту Яндекс Лицея (GoLang)

Этот сервис получает задачи от *оркестратора* и возвращает ему результат.

![SVG Image](service.svg)

## Требуется установка оркестратора

Подробнее тут: https://github.com/zhukovrost/orchestratorYL.git

## Установка

1. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/zhukovrost/agentYL.git
    ```

2. Перейдите в директорию проекта:
    ```sh
    cd agentYL
    ```
   
3. При необходимости установите окружение (Linux):

   ```sh 
   export PATH=$PATH:/usr/local/go/bin
   ```

4. Установите зависимости:
    ```sh
    go mod tidy
    ```

## Запуск

Для запуска сервера выполните:

```sh
go run cmd/agent/main.go
```

Также можно запустить программу с флагом -debug, чтобы видеть сообщения логгера на уровне debug.
Задать мощность можно с помощью переменной окружения COMPUTING_POWER.

```sh
COMPUTING_POWER=10 go run cmd/agent/main.go -debug
```

## Инструкция по использованию

Единственное, что необходимо, это просто запустить агента. 
Он будет отправлять запросы на получение задач оркестратору и выполнять их. 
Результат сам будет отправляться обратно к оркестратору.

## Структура проекта
- docker-compose.yml - файл конфигурации для Docker Compose.
- Dockerfile - файл с инструкциями по сборке образа Docker.
- run.sh - скрипт для автоматического запуска команд Docker Compose.

## Задание
В этом проекте используются Docker и Docker Compose для сборки и развертывания приложения. Схема и примеры будут инициализированны на этапе развертывания.
Вы можете запустить проект, используя предоставленный скрипт или вручную выполняя команды Docker Compose.

## Сбор и развертывание приложения
### Вручную через команды Docker Compose
- Должны быть установлены Docker и Docker Compose.
- Выполните команду для остановки и удаления контейнеров, томов и сетей, связанных с проектом:
```
docker-compose down -v
```
- Выполните команду для сборки и запуска приложения:
```
docker-compose up --build
```
### Через скрипт
- Сделайте скрипт исполняемым
```
chmod +x run.sh
```
- Запустить скрипт
```
./run.sh
```

## Примеры запросов после сборки
После деплоя оно будет доступно по адресу `http://localhost:8080`
Пример:
```
http://localhost:8080/api/ping
```
```
http://localhost:8080/api/tenders
```


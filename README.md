# Avito Parser

Схема работы приложения находится в файле appScheme.pdf

## Запуск проекта

```shell
$ git clone https://github.com/vpoletaev11/avitoParser
$ cd avitoParser
```

Перед запуском контейнера Вам необходимо настоить системные переменные в файле docker-compose.yml.

Укажите почту, с которой должна происходить рассылка сообщений в переменную `SENDER_MAIL`.

Пароль от этой почты в переменную `MAIL_PASSWORD`.

```shell
$ docker-compose up
```

## Отправка подписки на изменение цены в сервис

```shell
$ curl -d "email=example@mail.com&link=https://www.avito.ru/link/to/ad" -X POST localhost:8080/
```
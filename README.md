## Exchange App API

## Описание проекта

REST-сервис Exchanges предназначен для расчета всех вариантов размена указанной суммы денег с использованием доступных номиналов банкнот.

### Требования к оформлению

- ✅ конфигурация (хост, порт, уровень логирования) 
- ✅ graceful shutdown                                 
- ✅ unit тесты алгоритма                                 
- ✅ оформлен в общедоступном git репозитории       

### Что сделал дополнительно

- ✅ Генерация документации с помощью Swagger.
- ✅ Развёртывание в Docker контейнере.
- ✅ Makefile.
- ✅ Настроил CI/CD.

### Установка:

- #### Склонируйте репозиторий:
```bash
git clone https://github.com/prerec/media-mine.git
```

- #### Перейдите в директорию проекта
```bash
cd media-mine
```

- #### Установите зависимости:
```bash
go mod download
```

- #### Запустите тесты командой
```bash
make test
```

Результат выполнения можно наблюдать в консоли. Также в корне проекта будут сгенерированы файлы `cover.txt` и 
`index.html`. Откройте последний в браузере для наглядного представления покрытия тестами тестируемого пакета.

### Запуск:

#### 1) Локально:

- Запустите `main.go` файл:
```bash
make run
```

#### 2) В Docker контейнере:
Для корректной работы из контейнера - параметр `host` файла конфигурации должен быть равен `0.0.0.0`. 
При смене порта в файле конфигурации нужно будет также учесть это в Makefile, поэтому придётся внести изменения
в команду `docker run -d -p 8080:8080 exchanger`, установив новые значения.

1) Самостоятельная сборка:

- Соберите Docker образ:
```bash
make docker_build
```
- Запустите Docker контейнер:
```bash
make docker_run
```
2) Также можно воспользоваться образом из dockerhub, который собирается при каждом обновлении в ветке main
во время процессов CI/CD:

- Скачайте Docker образ с dockerhub:
```bash
make docker_pull
```
- Запустите Docker контейнер:
```bash
make dockerhub_run
```
### Использование:

Отправьте POST-запрос на адрес http://localhost:8080/api/exchange/ с JSON-телом, содержащим сумму и номиналы банкнот:

```json
{
  "amount": 400,
  "banknotes": [
    5000,
    2000,
    1000,
    500,
    200,
    100,
    50
  ]
}
```

### Документация:

Документация к API доступна по адресу http://localhost:8080/swagger/index.html


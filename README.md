# demo-downloader
Демо многопоточного скачивания файлов по HTTP.

### Структура проекта

* __cmd/__ - точка входа в приложение
* __config/__ - конфигурация через аргументы
* __internal/downloader__ - логика многопоточного скачивания
* __internal/storage__ - логика хранения и записи
* __internal/target__ - логика оценки файла

### Запуск

```bash
go run cmd/main.go -url https://speedtest.selectel.ru/100MB -output /tmp/my_file -chunk 1048576
```
### Параметры:

* __-url__ - ссылка на файл (обязательно)
* __-output__ - имя сохраняемого файла (обязательно)
* __-threads__ - размер одного чанка в байтах

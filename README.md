# `Среда для разработки Go Web-сервиса`

Базовая заготовка для веб-сервера. 
Содержит контейнеры: MySql,Go,Adminer.

Команды для запуска:

1) `docker-compose exec up -d` - запуск контейнеров
2) `watchexec --restart --exts "js,css,html,go" --watch . "docker-compose restart app"` - рестарт контенера при изменении файлов "js,css,html,go"

Перед запуском утилиты `watchec` её необходимо установить от сюда https://github.com/watchexec/watchexec/releases/tag/1.14.0 . На Ubuntu 18.04 я поставил файл https://github.com/watchexec/watchexec/releases/download/1.14.0/watchexec-1.14.0-x86_64-unknown-linux-gnu.deb
  

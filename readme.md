# Github hook listener

Слушает delete хук и выполняет команду

## Запуск:
```bash
github-hook-listener /path/to/config.json
```

## Пример конфига:

```json
{
  "Port": ":8000",
  "Path": "/deleteHook",
  "GithubSecret": "GithubSecretKey",
  "ExecCommand" : "./remove.sh"
}
```
### Параметры конфига:
* Port - порт для прослушивания
* Path - отностиельный путь
* GithubSecret - секретный ключ от репозитория, см https://github.com/vendor/repo/settings/hooks/new
* ExecCommand - команда для выполнения

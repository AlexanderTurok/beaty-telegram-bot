# beaty-telegram-bot
Telegram Bot for holding a Beaty Contest.

[Go Telegram SDK](https://github.com/go-telegram-bot-api/telegram-bot-api) for using Telegram API.
Postgres for storing users data. Redis for Caching.

# To start:
1. Create .env file
  - API_KEY = Telegram API key 
  - DB_PASSWORD = Password to your SQL database 
  - REDIS_PASSWORD = Redis Password 
2. Edit Config.yml in Configs folder
3. Execute command: go run cmd/main.go

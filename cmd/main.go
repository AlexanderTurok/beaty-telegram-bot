package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/telegram"
	"github.com/go-redis/redis/v9"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error while loading env: %s", err.Error())
	}

	botApi, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		log.Fatalf("error while getting api key: %s", err.Error())
	}

	botApi.Debug = true

	postgres, err := sql.Open("postgres", fmt.Sprintf("dbname=%s sslmode=disable", os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("error while openig postgres: %s", err.Error())
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	bot := telegram.NewBot(botApi, redis, postgres, ctx)

	bot.Start()
}

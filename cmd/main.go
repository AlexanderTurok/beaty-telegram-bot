package main

import (
	"context"
	"os"

	"github.com/AlexanderTurok/telegram-beaty-bot/internal/bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error while loading env: %s", err)
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	redis := repository.NewRedis(ctx, repository.Config{
		Host:     viper.GetString("redis.address"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	botApi, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		logrus.Fatalf("error while getting api key: %s", err)
	}

	botApi.Debug = false

	repository := repository.NewRepository(ctx, db, redis)
	service := service.NewService(repository)
	bot := bot.NewBot(botApi, service)

	bot.Start()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

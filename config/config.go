package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./files/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error read config file: %v", err)
	}

	cfg.App.Port = viper.GetString("app.port")
	cfg.App.Url = viper.GetString("app.url")

	cfg.Database.Host = viper.GetString("database.host")
	cfg.Database.Port = viper.GetString("database.port")
	cfg.Database.User = viper.GetString("database.user")
	cfg.Database.Password = viper.GetString("database.password")
	cfg.Database.Name = viper.GetString("database.name")

	cfg.Storage.UploadBaseDir = viper.GetString("storage.upload_base_dir")

	cfg.GRPC.OrderURL = viper.GetString("grpc.order_url")
	cfg.GRPC.URL = viper.GetString("grpc.url")

	cfg.Kafka.Broker = viper.GetString("kafka.broker")
	cfg.Kafka.KafkaTopics = viper.GetStringMapString("kafka.topics")

	return cfg

}

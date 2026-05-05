package initialize

import (
	"fmt"
	"os"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("production")
	viper.SetConfigType("yaml")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	env := os.Getenv("ENV_CONFIG")
	if env == "dev" {
		viper.SetConfigName("dev")
	}

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&global.Config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		return
	}
}

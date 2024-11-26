package infrastructure

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
)

func InitializeConfiguration() {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("error reading config file, %s\n", err)
		os.Exit(1)
	}

	fmt.Println("=== Load Conf List ===")
	infrastructureconfiguration.Port = viper.GetString("PORT")
	infrastructureconfiguration.ServiceName = viper.GetString("SERVICE_NAME")
	infrastructureconfiguration.Environment = viper.GetString("ENVIRONMENT")
	infrastructureconfiguration.RedisAddr = viper.GetString("REDIS_ADDR")
	infrastructureconfiguration.RedisDB = viper.GetInt("REDIS_DB")
	infrastructureconfiguration.RedisPass = viper.GetString("REDIS_PASS")
	infrastructureconfiguration.HashSalt = viper.GetString("HASH_SALT")
	infrastructureconfiguration.PostgresqlConn = viper.GetString("POSTGRESQL_CONN")
	infrastructureconfiguration.LoggingTool = viper.GetString("LOGGING_TOOL")
	infrastructureconfiguration.GelfAddr = viper.GetString("GELF_ADDR")
	infrastructureconfiguration.JaegerEndpoint = viper.GetString("JAEGER_ENDPOINT")
	infrastructureconfiguration.TracingTool = viper.GetString("TRACING_TOOL")
	infrastructureconfiguration.AspectoKey = viper.GetString("ASPECTO_KEY")
	infrastructureconfiguration.DatabaseName = viper.GetString("DATABASE_NAME")
	fmt.Println("=== End Load Conf List ===")

}

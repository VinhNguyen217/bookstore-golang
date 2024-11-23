package conf

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/samber/do"
)

type Config struct {
	ApiService struct {
		Port int64 `envconfig:"API_PORT"`
	}
	MYSQL struct {
		Host            string `envconfig:"MYSQL_HOST"`
		Port            int64  `envconfig:"MYSQL_PORT"`
		User            string `envconfig:"MYSQL_USER"`
		Password        string `envconfig:"MYSQL_PASSWORD"`
		DBName          string `envconfig:"MYSQL_DBNAME"`
		MigrationFolder string `envconfig:"MYSQL_MIGRATION_FOLDER"`
	}

	JWT struct {
		PublicKeyFilePath  string `envconfig:"JWT_PUBLIC_KEY_FILE_PATH"`
		PrivateKeyFilePath string `envconfig:"JWT_PRIVATE_KEY_FILE_PATH"`
	}

	Casbin struct {
		RBACModelPath  string `envconfig:"RBAC_MODEL_PATH"`
		RBACPolicyPath string `envconfig:"RBAC_POLICY_PATH"`
	}
}

func NewConfig(di *do.Injector) (*Config, error) {
	cf := &Config{}
	_ = godotenv.Load(".env")
	err := envconfig.Process("", cf)
	return cf, err
}

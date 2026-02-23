package config

type Config struct {
	App      AppConfig      `yaml:"app" validate:"required"`
	Database DatabaseConfig `yaml:"database" validate:"required"`
	Storage  StorageConfig  `yaml:"storage" validate:"required"`
	GRPC     GRPCConfig     `yaml:"grpc" validate:"required"`
	Kafka    KafkaConfig    `yaml:"kafka" validate:"required"`
	Secret   SecretConfig   `yaml:"secret" validate:"required"`
}

type SecretConfig struct {
	JWTSecret string `yaml:"jwt_secret" validate:"required"`
}

type AppConfig struct {
	Port string `yaml:"port" validate:"required"`
	Url  string `yaml:"url" validate:"required"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Name     string `yaml:"name" validate:"required"`
	Port     string `yaml:"port" validate:"required"`
}

type StorageConfig struct {
	UploadBaseDir string `yaml:"upload_base_dir" validate:"required"`
}

type GRPCConfig struct {
	OrderURL string `yaml:"order_url" validate:"required"`
	URL      string `yaml:"url" validate:"required"`
}

type KafkaConfig struct {
	Broker      string            `yaml:"broker" validate:"required"`
	KafkaTopics map[string]string `yaml:"topics" validate:"required"`
}

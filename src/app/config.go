package app

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	port      int
	version   string
	env       string
	jwtSecret string
	db        DB
}

func NewConfig() (*Config, error) {
	cfg := Config{
		port:    8080,
		version: "1.0.0",
	}

	flag.IntVar(&cfg.port, "port", 4000, "Porta do servidor da API")
	flag.StringVar(&cfg.env, "env", "desenvolvimento", "Ambiente (desenvolvimento|homologação|produção)")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Número máximo de conexões abertas no PostgreSQL")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "Número máximo de conexões inativas no PostgreSQL")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "Tempo máximo de conexão inativa no PostgreSQL")

	env, err := loadEnv("../.env")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("falha ao ler arquivo .env: %v", err))
	}

	key := "DATABASE_URL"
	if value, exists := getEnvValue(env, key); exists {
		cfg.db.dsn = value
	} else {
		return nil, errors.New(fmt.Sprintf("Chave %s não foi encontrada no arquivo .env", key))
	}

	key = "JWT_SECRET"
	if value, exists := getEnvValue(env, key); exists {
		cfg.jwtSecret = value
	} else {
		return nil, errors.New(fmt.Sprintf("Chave %s não foi encontrada no arquivo .env", key))
	}

	return &cfg, nil
}

func getEnvValue(env map[string]string, key string) (string, bool) {
	value, exists := env[key]
	return value, exists
}

func loadEnv(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	env := map[string]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return env, nil
}

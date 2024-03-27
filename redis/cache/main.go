package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Inicializando o cliente Redis
	client, err := InitRedis()
	if err != nil {
		fmt.Println("Erro ao inicializar o cliente Redis:", err)
		return
	}
	defer client.Close()

	// Exemplo de uso do Redis para armazenar e recuperar dados
	ctx := context.Background()

	// Definindo uma chave e seu valor
	key := "hello"
	value := "world"

	// Armazenando a chave e o valor no Redis com um tempo de expiração de 1 minuto
	err = client.Set(ctx, key, value, 1*time.Minute).Err()
	if err != nil {
		fmt.Println("Erro ao definir a chave no Redis:", err)
		return
	}

	// Recuperando o valor associado à chave
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Erro ao recuperar o valor do Redis:", err)
		return
	}
	fmt.Println("Valor recuperado do Redis:", val)

	// Aguardando alguns segundos para que a chave expire
	time.Sleep(2 * time.Second)

	// Tentando recuperar o valor novamente após a expiração
	val, err = client.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("A chave expirou no Redis")
	} else if err != nil {
		fmt.Println("Erro ao recuperar o valor do Redis:", err)
		return
	} else {
		fmt.Println("Valor recuperado do Redis ja expirado:", val)
	}
}

// Função para inicializar o cliente Redis
func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Endereço do servidor Redis
		Password: "",               // Senha (opcional)
		DB:       0,                // Número do banco de dados (padrão 0)
	})
	return client, nil
}

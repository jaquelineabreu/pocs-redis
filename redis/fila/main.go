package main

import (
	"context"
	"fmt"

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

	// Exemplo de uso do Redis como fila
	ctx := context.Background()

	// Nome da fila no Redis
	queueName := "minha_fila"

	// Produtor: adicionar itens à fila
	for i := 0; i < 5; i++ {
		item := fmt.Sprintf("item_%d", i)
		err := client.LPush(ctx, queueName, item).Err()
		if err != nil {
			fmt.Println("Erro ao adicionar item à fila:", err)
		} else {
			fmt.Println("Item adicionado à fila:", item)
		}
	}

	// Consumidor: remover itens da fila
	for {
		item, err := client.BRPop(ctx, 0, queueName).Result()
		if err != nil {
			fmt.Println("Erro ao remover item da fila:", err)
			break
		}
		fmt.Println("Item removido da fila:", item[1])
	}

	// Fechar a conexão com o Redis
	client.Close()
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

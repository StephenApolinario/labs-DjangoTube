package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// Conectando ao RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Erro ao abrir o canal:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	// Declarando a exchange do tipo "topic"
	err = ch.ExchangeDeclare(
		"my-ex", //Nome da exchange
		"topic", // Tipo da exchange
		true,    // Durável
		false,   // Auto-delete
		false,   // Não é uma exchange interna
		false,   // No-wait
		nil,     // Sem argumentos adicionais
	)

	if err != nil {
		log.Fatal("Error declaring the exchange: ", err)
	}

	// Definindo a mensagem
	if len(os.Args) < 2 {
		log.Fatal("Você precisa fornecer uma mensagem para enviar.")
	}
	message := os.Args[1]

	// Definindo a rounting key (Opcioanl)
	routingKey := "" // Valor padrão
	if len(os.Args) > 2 {
		routingKey = os.Args[2] // Usa a routing key fornecida, se houver
	}

	err = ch.Publish(
		"my-ex",    // Nome da exchange
		routingKey, // Routing key (Optional)
		false,      // Não é obrigatório encontrar fila
		false,      // Não é obrigatório entregar imediatamente
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatal("Erro ao publicar a mensagem: ", err)
	}

	fmt.Printf("Mensagem enviada: %s com a routing key %s\n", message, routingKey)
}

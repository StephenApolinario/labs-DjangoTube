package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Conectando ao RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Erro ao abrir ao conectar ao RabbitMQ:", err)
	}
	defer conn.Close()

	// Abrindo um canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Erro ao abrir o canal:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logs", // Nome da fila
		true,   // Durável
		false,  // Delete quando não estiver em uso
		false,  // Exclusiva
		false,  // No-wait
		nil,    // Argumentos adicionais
	)
	if err != nil {
		log.Fatal("Erro ao declarar a fila:", err)
	}

	// Ordem com Auto-ACK False
	// 1. Le a fila de pagamentos
	// 2. Processa a mensagem
	// 3. Acknowledge a mensagem

	// Ordem com Auto-ACK True
	// 1. Acknowledge a mensagem
	// 2. Le a fila de pagamentos
	// 3. Processa a mensagem

	// Se o Auto-ACK estiver ativado, a mensagem é automaticamente reconhecida (acknowledged) assim que é lida da fila, antes mesmo de ser processada.
	// Problemas que podem ocorrer se o sistema cair com Auto-ACK ativado:
	// 1. Perda de mensagens: Se o sistema cair após a mensagem ser lida, mas antes de ser processada, a mensagem será perdida, pois já foi reconhecida.
	// 2. Falta de confiabilidade: Mensagens importantes podem não ser processadas, levando a inconsistências ou falhas no sistema.
	// 3. Dificuldade de recuperação: Sem saber quais mensagens foram processadas, pode ser difícil ou impossível recuperar o estado correto do sistema após uma falha.

	// Melhor caso:
	// Auto ACK sempre Falso
	// Processa a mensagem
	// ACK Manual

	// Já fez muitas empresas perderem dinheiro quando cai o sistema

	// Consumindo mensagens da fila
	msgs, err := ch.Consume(
		q.Name,  // Nome da fila
		"goapp", // Identificador do consumidor
		true,    // Auto-ack CUIDADO: Reconhece automaticamente que a mensagem foi recebida (Se a mensagem não for processada, ela será perdida porque apaga automaticamente)
		false,   // Exclusiva
		false,   // No-local
		false,   // No-wait
		nil,     // Argumentos adicionais
	)
	if err != nil {
		log.Fatal("Erro ao consumir a fila:", err)
	}

	forever := make(chan bool)

	// Loop para processar as mensagens recebidas
	go func() {
		for msg := range msgs {
			log.Printf("Recebida uma mensagem: %s", msg.Body)
		}
	}()

	log.Println("Aguardando mensagens...")
	<-forever
}

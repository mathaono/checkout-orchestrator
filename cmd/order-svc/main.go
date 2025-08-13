package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Configuração de logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Canal para graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// TODO: Inicializar componentes
	// - Conexão com MySQL
	// - Conexão com Kafka
	// - Servidor HTTP

	log.Println("Serviço de pedidos iniciado...")

	// Aguarda sinal de shutdown
	<-quit
	log.Println("Encerrando serviço de pedidos...")
}

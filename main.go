package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "rabbitmq-connection-checker",
	Short: "RabbitMQ connection checker",
	RunE:  run,
}

var host string
var port int
var user string
var pass string
var tls bool

func init() {
	cmd.Flags().StringVar(&host, "host", "127.0.0.1", "Host")
	cmd.Flags().IntVar(&port, "port", 5672, "Port")
	cmd.Flags().StringVarP(&user, "user", "u", "", "Username")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "Password")
	cmd.Flags().BoolVar(&tls, "tls", false, "Password")
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := user
	if pass != "" {
		auth = fmt.Sprintf("%s:%s", user, pass)
	}

	log.Println("Checking connection to", addr)

	protocol := "amqp"
	if tls {
		protocol = "amqps"
	}

	conn, err := amqp.Dial(fmt.Sprintf("%s://%s@%s/", protocol, auth, addr))
	failOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Println("Connection OK")

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

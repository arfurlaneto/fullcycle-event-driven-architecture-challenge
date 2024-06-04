package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/database"
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/usecase/get_balance"
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/usecase/update_balance"
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/web"
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/web/webserver"
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDb := database.NewBalanceDB(db)
	getBlance := get_balance.NewGetBalanceUseCase(balanceDb)
	updateBalance := update_balance.NewUpdateBalanceUseCase(balanceDb)

	webserver := webserver.NewWebServer(":8080")
	clientHandler := web.NewWebBalanceHandler(*getBlance)
	webserver.AddHandler("/balances/{account_id}", clientHandler.GetBalance)

	go startKafkaConsumer(updateBalance)

	fmt.Println("Server is running")
	webserver.Start()
}

type BalanceUpdatedKafkaEvent struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom        string  `json:"account_id_from"`
		AccountIDTo          string  `json:"account_id_to"`
		BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
		BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
	} `json:"Payload"`
}

func startKafkaConsumer(updateBalance *update_balance.UpdateBalanceUseCase) {
	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})
	msgChan := make(chan *ckafka.Message)
	go kafkaConsumer.Consume(msgChan)

	for {
		msg := <-msgChan
		event := BalanceUpdatedKafkaEvent{}
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Updating %s balance to %f\n", event.Payload.AccountIDFrom, event.Payload.BalanceAccountIDFrom)
			updateBalance.Execute(update_balance.UpdateBalanceInputDTO{AccountId: event.Payload.AccountIDFrom, Balance: event.Payload.BalanceAccountIDFrom})
			fmt.Printf("Updating %s balance to %f\n", event.Payload.AccountIDTo, event.Payload.BalanceAccountIDTo)
			updateBalance.Execute(update_balance.UpdateBalanceInputDTO{AccountId: event.Payload.AccountIDTo, Balance: event.Payload.BalanceAccountIDTo})
		}
	}
}

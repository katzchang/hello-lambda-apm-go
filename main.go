package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	newrelic "github.com/newrelic/go-agent"
	"github.com/newrelic/go-agent/_integrations/nrlambda"
	insights "github.com/newrelic/go-insights/client"
	"math/rand"
	"os"
)

func handler(ctx context.Context) {
	if txn := newrelic.FromContext(ctx); nil != txn {
		_ = txn.AddAttribute("userLevel", "gold")
		_ = txn.Application().RecordCustomEvent("MyEvent", map[string]interface{}{
			"zip": "zap",
		})
	}
	fmt.Println("hello world")
}

func main() {
	//lambda.Start(handler)
	cfg := nrlambda.NewConfig()
	app, err := newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println("error creating app (invalid config):", err)
	}
	insightInsertKey := os.Getenv("NEW_RELIC_INSIGHTS_INSERT_KEY")
	client := insights.NewInsertClient(insightInsertKey, os.Getenv("NEW_RELIC_ACCOUNT_ID"))
	h := NRHandler{client: client}
	nrlambda.Start(h.Xx, app)
}

type TestType struct {
	EventType    string `json:"eventType"`
	AwesomeScore int    `json:"AwesomeScore"`
}

type NRHandler struct {
	client *insights.InsertClient
}

func (h *NRHandler) Xx(ctx context.Context) {
	score := rand.Intn(100)
	testData := TestType{
		EventType:    "HelloEvent",
		AwesomeScore: score,
	}

	if postErr := h.client.PostEvent(testData); postErr != nil {
		log.Errorf("Error: %v\n", postErr)
	}
}

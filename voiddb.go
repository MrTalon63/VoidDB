package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
	"voiddb/endpoint"
	"voiddb/memory"
)

func Save() {

	if !memory.StateChangedSaveRequired {
		return
	}

	start := time.Now().UnixNano()
	jsonString, _ := json.Marshal(memory.Data)
	err := ioutil.WriteFile("save.json", jsonString, os.ModePerm)

	if err != nil {
		fmt.Print(jsonString)
		panic(err)
	}

	end := time.Now().UnixNano()
	elapsed := (start - end) * int64(time.Nanosecond/time.Millisecond)

	fmt.Printf("Saved %d entries in %d ms\n", len(memory.Data), elapsed)
	memory.StateChangedSaveRequired = false
}

func Load() {

	start := time.Now().UnixNano()
	jsonString, err := ioutil.ReadFile("save.json")
	if err != nil {
		panic(err)
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(jsonString, &jsonMap)
	if err != nil {
		panic(err)
	}

	memory.Data = jsonMap
	end := time.Now().UnixNano()
	elapsed := (start - end) * int64(time.Nanosecond/time.Millisecond)

	fmt.Printf("Loaded %d entries in %d ms\n", len(memory.Data), elapsed)
}

func main() {

	// load data
	Load()

	// fiber setup
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	// routing
	app.Get("/", endpoint.Index)
	app.Get("/stats", endpoint.Stats)
	app.Get("/dump", endpoint.Dump)
	app.Post("/db/{key}", endpoint.Set)
	app.Get("/db/{key}", endpoint.Get)
	app.Delete("/db/{key}", endpoint.Del)

	// auto-save
	c := cron.New()
	_, _ = c.AddFunc("@every 5s", Save)
	c.Start()

	// exit-save
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		Save()
		os.Exit(0)
	}()

	// bind
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

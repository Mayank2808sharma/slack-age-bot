package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

// Bot represents the Slack bot
type Bot struct {
	client *slacker.Slacker
}

// NewBot initializes a new instance of Bot
func NewBot(token, appToken string) *Bot {
	return &Bot{
		client: slacker.NewClient(token, appToken),
	}
}

// Start begins listening for Slack events
func (b *Bot) Start(ctx context.Context) error {
	go b.printCommandEvents()
	b.registerCommands()
	return b.client.Listen(ctx)
}

// printCommandEvents logs command events
func (b *Bot) printCommandEvents() {
	for event := range b.client.CommandEvents() {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// registerCommands sets up the commands for the bot
func (b *Bot) registerCommands() {
	b.client.Command("my dob is <dob>", &slacker.CommandDefinition{
		Description: "Calculate age based on date of birth",
		Handler:     dobCommandHandler,
	})

	b.client.Command("<dob>", &slacker.CommandDefinition{
		Description: "Calculate age based on date of birth",
		Handler:     simpleDobCommandHandler,
	})
}

// dobCommandHandler handles the "my dob is <dob>" command
func dobCommandHandler(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	dob := request.Param("dob")
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		response.Reply("Invalid date of birth format. Please provide a valid date in YYYY-MM-DD format.")
		return
	}
	age := calculateAge(dateOfBirth)
	response.Reply(fmt.Sprintf("Your age is %d.", age))
}

// simpleDobCommandHandler handles the "<dob>" command
func simpleDobCommandHandler(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	dobStr := request.Param("dob")
	fmt.Printf("Received DOB: %s\n", dobStr) // Debugging log

	// Extracting the date from the parameter
	parts := strings.Fields(dobStr)
	if len(parts) < 2 {
		response.Reply("Please provide a valid date in YYYY-MM-DD format.")
		return
	}
	dateStr := parts[1]

	dob, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		response.Reply("Invalid date format. Please provide the date in YYYY-MM-DD format.")
		return
	}
	age := calculateAge(dob)
	response.Reply(fmt.Sprintf("Your age is %d.", age))
}

// calculateAge calculates age from the date of birth
func calculateAge(dateOfBirth time.Time) int {
	now := time.Now()
	age := now.Year() - dateOfBirth.Year()
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}
	return age
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	if botToken == "" || appToken == "" {
		log.Fatal("Slack tokens are required")
	}

	bot := NewBot(botToken, appToken)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := bot.Start(ctx); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	aai "github.com/AssemblyAI/assemblyai-go-sdk"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("ASSEMBLYAI_API_KEY")
	client := aai.NewClient(apiKey)

	ctx := context.Background()

	files, err := os.ReadDir("./audio")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f, err := os.Open("./audio/" + file.Name())
		if err != nil {
			fmt.Println("Something bad happened:", err)
			os.Exit(1)
		}

		fileName := file.Name()
		fileName = fileName[:len(fileName)-4]

		transcript, err := client.Transcripts.TranscribeFromReader(ctx, f, &aai.TranscriptOptionalParams{
			SpeakerLabels: aai.Bool(true),
		})

		if err != nil {
			fmt.Println("Something bad happened:", err)
			os.Exit(1)
		}

		report := CreateReport(fileName, *transcript.Text)
		report.Generate()
	}
}

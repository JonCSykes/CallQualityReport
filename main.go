package main

import (
	"context"
	"fmt"
	"log"
	"os"

	aai "github.com/AssemblyAI/assemblyai-go-sdk"
)

func main() {
	apiKey := "5d8d209eb0e94fe68ce1de241d6fc4f7"
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

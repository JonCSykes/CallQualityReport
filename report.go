package main

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/invopop/jsonschema"
	"github.com/openai/openai-go"
)

type Report struct {
	Name         string
	Transcript   string
	CallAnalysis *CallAnalysis `json:"call_analysis"`
}

type CallAnalysis struct {
	Overview                    string      `json:"overview"`
	Strengths                   []string    `json:"strengths"`
	Weaknesses                  []string    `json:"weaknesses"`
	OpportunitiesForImprovement []string    `json:"opportunities_for_improvement"`
	PerformanceRating           Performance `json:"performance_rating"`
	Conclusion                  string      `json:"conclusion"`
}

type Performance struct {
	EmpathyAndEmotionalIntelligence float64 `json:"empathy_and_emotional_intelligence"`
	Professionalism                 float64 `json:"professionalism"`
	ProblemSolvingSkills float64 `json:"problem_solving
_skills"`
	CommunicationClarity    float64 `json:"communication_clarity"`
	CustomerCentricApproach float64 `json:"customer_centric_approach"`
	OverallRating           float64 `json:"overall_rating"`
}

func GenerateSchema[T any]() interface{} {
	// Structured Outputs uses a subset of JSON schema
	// These flags are necessary to comply with the subset
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

// Generate the JSON schema at initialization time
var CallAnalysisSchema = GenerateSchema[CallAnalysis]()

func CreateReport(name, transcript string) *Report {
	envVars := map[string]string{
		"OPENAI_API_KEY":    os.Getenv("OPENAI_API_KEY"),
		"OPENAI_PROJECT_ID": os.Getenv("OPENAI_PROJECT_ID"),
	}

	for key, value := range envVars {
		err := os.Setenv(key, value)
		if err != nil {
			log.Fatalf("Failed to set environment variable %s: %v", key, err)
		}
	}

	// Verifying environment variables
	for key := range envVars {
		_, exists := os.LookupEnv(key)
		if !exists {
			log.Fatalf("Environment variable %s is not set", key)
		}
	}

	callAnalysis := Analyze(transcript)

	return &Report{
		Name:         name,
		Transcript:   transcript,
		CallAnalysis: callAnalysis,
	}
}

func Analyze(transcript string) *CallAnalysis {
	client := openai.NewClient()
	ctx := context.Background()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        openai.F("callAnalysis"),
		Description: openai.F("The analysis of a customer service call"),
		Schema:      openai.F(CallAnalysisSchema),
		Strict:      openai.Bool(true),
	}

	reportingGuidelines, err := os.ReadFile("report-guidelines.md")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Query the Chat Completions API
	chat, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(string(reportingGuidelines)),
			openai.UserMessage(transcript),
		}),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](
			openai.ResponseFormatJSONSchemaParam{
				Type:       openai.F(openai.ResponseFormatJSONSchemaTypeJSONSchema),
				JSONSchema: openai.F(schemaParam),
			},
		),
		// Only certain models can perform structured outputs
		Model: openai.F(openai.ChatModelGPT4o2024_08_06),
	})

	if err != nil {
		panic(err.Error())
	}

	callAnalysis := CallAnalysis{}
	err = json.Unmarshal([]byte(chat.Choices[0].Message.Content), &callAnalysis)
	if err != nil {
		panic(err.Error())
	}

	return &callAnalysis
}

func (report *Report) Generate() {
	markdownTemplate, err := os.ReadFile("report-template.md")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a new template and parse the markdown into it
	tmpl, err := template.New("report").Parse(string(markdownTemplate))
	if err != nil {
		log.Fatalf("Error creating template: %v", err)
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, report.CallAnalysis); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	reportFile, err := os.Create("reports/" + report.Name + "_" + time.Now().Format("20060102150405") + ".md")
	if err != nil {
		log.Fatalf("Error creating report file: %v", err)
	}

	if _, err := reportFile.Write(buffer.Bytes()); err != nil {
		log.Fatalf("Error writing to report file: %v", err)
	}

	if err := reportFile.Close(); err != nil {
		log.Fatalf("Error closing report file: %v", err)
	}

	log.Printf("Report generated successfully: %s", report.Name)
}

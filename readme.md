# Call Center Analysis Project

This project analyzes call center audio files, generates transcripts, and creates detailed performance reports for customer service representatives.  

## Prerequisites


- Git (https://git-scm.com/downloads)
- Go 1.20 or later (https://golang.org/dl/)
- An AssemblyAI API key (https://assemblyai.com)
- An OpenAI API key, Project ID, and configured Assistant (https://platform.openai.com)


## Installation

### Clone the repository:  

```
git clone https://github.com/yourusername/call-center-analysis.git
cd call-center-analysis
```

### Install dependencies:  
`go mod tidy`

## Configuration

###Set up environment variables:
```
export ASSEMBLYAI_API_KEY=your_assemblyai_api_key
export OPENAI_API_KEY=your_openai_api_key
export OPENAI_PROJECT_ID=your_openai_project_id
```
### Ensure you have audio files in the audio directory:  

`mkdir -p audio`

# Add your audio files to the audio directory


## Running the Project

### Run the main application:  
`go run main.go`

The application will process each audio file in the audio directory, generate transcripts, analyze the calls, and create markdown reports in the reports directory.  

## Project Structure
- `main.go`: Entry point of the application. Reads audio files, generates transcripts, and creates reports.
- `report.go`: Contains the logic for creating and generating reports.
- `reports/`: Directory where the generated reports are saved.
- `audio/`: Directory where the audio files to be processed are stored.
- `report-template.md`: Template for generating the markdown reports.
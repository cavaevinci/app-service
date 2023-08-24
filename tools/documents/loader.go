package documents

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/tmc/langchaingo/tools"
)

const DefualtRootPath = "./data/documents/"

var _ tools.Tool = &DocumentsTool{} 
var ErrFileLoad = "Error while loading requested file"

type DocumentsTool struct {
	RootPath string
}

func NewLoader(options ...DocuemntsToolOption) (*DocumentsTool, error) {
	documents := &DocumentsTool{
		RootPath: DefualtRootPath,
	}

	for _, option := range options {
		option(documents)
	}
	
	return documents, nil
}

func (tool *DocumentsTool) Call(ctx context.Context, input string) (string, error) {
	fmt.Println("Retrieving documents with input...")
	fmt.Println(input)

	var toolInput struct {
		FileName string
		Query    string
	}

	re := regexp.MustCompile(`(?s)\{.*\}`)
	jsonString := re.FindString(input)

	err := json.Unmarshal([]byte(jsonString), &toolInput)
	if err != nil {
		return "", err
	}
	
	path := tool.RootPath + toolInput.FileName

	fileContents, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("%v: %s", ErrFileLoad, err), nil
	}

	return string(fileContents), nil
} 

func (agent *DocumentsTool) Name() string {
	return "Library Agent"	
}

func (agent *DocumentsTool) Description() string {
	return `
		Library Agent is an agent specialized in fetching and reading documents from your internal knowledge library. It helps you access relevant information to provide accurate and informative responses.
		
		**Usage Example**:
		1. Identify the appropriate document from the library list based on the user's query.
		2. Use the "FileName" parameter to request the document in the JSON format:
		{
			"FileName": "filename.txt"
		}		
		3. Receive the content of the file as output.
		4. Integrate the content into your response in a coherent and natural way.
		
		**Available Documents**:
		- Name: Astro Synapse Impressum
		  - FileName: astrosynapse_impressum.txt
		  - Description: General information about Astro Synapse. Use to answer any information related to Astro Synapse, such as the company's mission, vision, website, etc.
		- Name: Astro Synapse Executive Summary
		  - FileName: astrosynapse_es.txt
		  - Description: Summary of the Astro Synapse Executive Summary. Use to answer questions about Astro Synapse's business plans, investment asks and cycles, broader details about finance and operations.
		- Name: Astro Synapse Team
		  - FileName: astrosynapse_team.txt
		  - Description: Information about the Astro Synapse team. Contains short bios and relevant links. Use this to retrieve information about the Astro Synapse team.
		- Name: ASAI Architecture
		  - FileName: asai_architecture.txt
		  - Description: ASAI Architecture describes the code and technologies used to build and run ASAI. Use this document to understand how ASAI works and to answer any technical questions about ASAI.
		- Name: Creator Impressum
		  - FileName: creator_impressum.txt
		  - Description: Information about the creator of ASAI. Use to answer any questions about the creator of ASAI.
		- Name: Welcome Script
		  - FileName: welcome_script.txt
		  - Description: Welcome script for new users. Use this document when you are prompted to welcome or onboard a new user to our system.
		
		**Error Handling**:
		If you encounter an error or cannot fetch the requested document, inform the user and offer alternative assistance or information.
	`
}
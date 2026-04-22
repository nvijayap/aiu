/*----------------------------------------------------------------------------------

Yes, the Agent-to-Agent (A2A) protocol is considered one of the latest and
most significant approaches to agentic AI, specifically focusing on
interoperability and collaboration between AI agents.

Introduced by Google in April 2025 and currently stewarded by the Linux Foundation,
A2A is designed to allow AI agents from different providers or frameworks to
communicate and work together to solve complex, multi-step tasks.

Here is an analysis of A2A in the context of the latest agentic AI trends:

. What Makes A2A a "Latest" Approach?

  . Interoperability Focus:

    While previous approaches (like LangChain or crewAI) focused on
	creating agentic workflows within their own ecosystem, A2A acts
	as a "universal translator" or messaging tier that lets agents
	from different vendors (e.g., Salesforce, SAP, Google)
	interact directly.

  . Move to Production:

	A2A addresses the shift from pilot projects to enterprise-level,
	production-ready, multi-agent systems that cross application boundaries.

  . Protocol-Based Collaboration:

    A2A is an open-source, vendor-neutral protocol that enables agents to
	discover each other, negotiate tasks, and hand off work securely using
	web standards like JSON-RPC and HTTP.

A2A vs. Other "Latest" Approaches

A2A is often confused with or used alongside
other modern agentic developments:

A2A (Agent-to-Agent): Focuses on collaboration between agents
(e.g., a "recruiter" agent talking to a "scheduler" agent).

MCP (Model Context Protocol): Introduced by Anthropic,
this acts as the "USB-C for AI," focusing on how agents
connect to tools, data sources, and APIs
(e.g., an agent accessing a database).

A2A and MCP together: These are complementary, not competing.

A2A enables the "talk" (coordination),
while MCP enables the "do" (tool interaction).

ACP (Agent Communication Protocol):

IBM's, now largely merged or aligned with, A2A/open-source standards.

Key Features and Limitations

Strengths: Strong industry backing (60+ partners, including LangChain, Salesforce, and Microsoft),
ability to handle long-running tasks, and native support for security (e.g., Authentication/Authorization).

Current State: It is very new (introduced April 2025), so while it has significant momentum,
it is still developing, and widespread production examples are emerging.

Challenges: As A2A connects more agents, it creates complex, un-ordered, and
potentially "tangled" webs of interconnections, creating a need for robust
governance and monitoring.

In summary, A2A is not just a tool but a standardized communication protocol
that is currently leading the push toward a more collaborative and
interoperable "Internet of Agents" in 2025-2026.

The Agent-to-Agent (A2A) protocol is widely considered one of the latest and
most representative approaches to enabling multi-agent collaboration.

Introduced by Google in April 2025, it serves as an open-standard communication
layer that allows AI agents from different vendors and frameworks to "talk" to
each other, discover capabilities, and delegate tasks.

A2A is not a replacement for agent frameworks like LangChain or CrewAI; rather,
it is a messaging protocol that sits above them to ensure interoperability.

Key Characteristics of A2A

Interoperability:

Enables a "planning agent" from one vendor to delegate a subtask to a "specialized agent"
from another, breaking down vendor silos.

Discovery via "Agent Cards":

Each agent publishes a standardized JSON document describing its identity,
endpoints, authentication requirements, and specific skills.

Long-Running Tasks:

Specifically designed to manage tasks that may take hours or days,
providing real-time status updates and feedback mechanisms.

Web Standard Foundation: Built on familiar technologies like HTTP/HTTPS, JSON-RPC,
and Server-Sent Events (SSE) to ease enterprise adoption.

A2A vs. MCP (The "Talk" vs. "Do" Distinction)

A2A is frequently paired with Anthropic's Model Context Protocol (MCP),
which was released in late 2024.

They address different layers of the agentic stack:

MCP (Context/Execution): Acts like a "universal adapter" (USB-C for AI)
to connect agents to data sources, local files, and external tools.

A2A (Coordination/Communication): Acts as the communication bus for agents
to collaborate with each other.

Latest Trends and Industry Backing

As of early 2026, the A2A protocol has moved from its origins at Google to
become a project under the Linux Foundation. It has gained broad backing from
over 150 organizations, including Microsoft, Salesforce, SAP, and ServiceNow.

While A2A is a leading approach, some industry experts view it as a "bridge"
to the "Internet of Agents", though it faces ongoing challenges regarding complex
security implementations and resource efficiency in edge computing environments.

----------------------------------------------------------------------------------*/

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/ollama/ollama/api"
)

// main
func main() {

	args := os.Args
	l := len(args)
	if l != 2 {
		fmt.Printf("\n\tNeed city name as arg\n\n")
		return
	}
	city := os.Args[1]

	fmt.Println()
	log.Println("Setting up the client from environment")
	client, err := api.ClientFromEnvironment()
	if err != nil {
		fmt.Println("main: err1:", err)
	}

	log.Println("Defining the tool for Ollama")
	toolPropertiesMap := api.NewToolPropertiesMap()
	toolPropertiesMap.Set(
		"location", 
		api.ToolProperty{
			// Type: "",
			Description: "Name of city like 'San Jose'",
		},
	)

	weatherTool := api.Tool{
		Type: "function",
		Function: api.ToolFunction{
			Name:        "get_weather",
			Description: "Get the current weather for a city",
			Parameters: api.ToolFunctionParameters {
				Type: "string",
				Properties: toolPropertiesMap,
			},
		},
	}

	log.Println("Initial request with model that supports tools")
	req := &api.ChatRequest{
		Model:    "llama3.1",
		Messages: []api.Message{{Role: "user", Content: fmt.Sprintf("Weather in %s?", city)}},
		Tools:    []api.Tool{weatherTool},
	}

	log.Println("Processing tool call and appending result")
	ctx := context.Background()
	content, singleLine := "", ""
	re := regexp.MustCompile(`[\n\f\t]+`)
	client.Chat(ctx, req, func(resp api.ChatResponse) error {
		if len(resp.Message.ToolCalls) > 0 {
			result := getWeather(city)
			req.Messages = append(req.Messages, resp.Message)
			req.Messages = append(req.Messages, api.Message{Role: "tool", Content: result})
			client.Chat(ctx, req, func(finalResp api.ChatResponse) error {
				content += finalResp.Message.Content
				return nil
			})
		}
		singleLine = re.ReplaceAllString(content, " ")
		return nil
	})

	singleLine = strings.ReplaceAll(singleLine, ". ", ".\n")
	fmt.Printf("%s\n\n", singleLine)
}

// get weather
func getWeather(location string) string {
	// jisnukrsna.world comes under the umbrella of mehersys.com (my company)
	endPoint := "https://jisnukrsna.world:7878/gw"
	formData := url.Values{
        "city": {location},
    }
	resp, err := http.PostForm(endPoint, formData)
	if err != nil {
		fmt.Println("getWeather: err1:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getWeather: err2:", err)
	}
	jsonString := string(body)
	// fmt.Println("jsonString:", jsonString)
	jsonString = strings.ReplaceAll(jsonString, "{", fmt.Sprintf(`{"location":"%s",`, location))
	// fmt.Println("jsonString:", jsonString)
	body = []byte(jsonString)
	var formattedJSON bytes.Buffer
	err = json.Indent(&formattedJSON, body, "", "    ")
	if err != nil {
		fmt.Println("getWeather: err3:", err)
	}
	fmt.Printf("\nformattedJSON from weather service:\n\n%s\n\n", &formattedJSON)
	return formattedJSON.String()
}

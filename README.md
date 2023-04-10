# Lightweight AGI: Building General AI with Golang and Flexible Language Models

Lightweight AGI is a project aimed at creating a simple and effective Artificial General Intelligence (AGI) agent using Golang and versatile Large Language Models (LLMs) compatible with the LLMClient interface. The agent is designed to handle a wide range of objectives by refining objectives, executing tasks, evaluating results, and prioritizing further tasks. With its adaptable architecture, Lightweight AGI can be applied to various domains, including but not limited to gaming, problem-solving, and knowledge acquisition.

## Features

- **Objective Refinement**: A unique and innovative approach to refining the main objective into smaller, more manageable objectives, enabling the AGI agent to tackle complex problems effectively.
- Task Creation: Generates tasks and milestones based on the refined objectives.
- Execution Agent: Executes tasks using the OpenAI GPT-4 model (or GPT-3.5 model).
- Evaluation Agent: Evaluates the results of tasks and their effectiveness.
- Prioritization Agent: Prioritizes tasks based on their relevance and importance.
- Task Context Agent: Stores the context of tasks for future reference.
- **In-Memory Vector Store**: Supports an efficient in-memory vector store to save embedding vectors, which can be queried for similarity, enabling faster access and improved performance.

## Note: Currently Supported LLM Provider

While Lightweight AGI is designed to be flexible and work with various Large Language Models (LLMs) that satisfy the LLMClient interface, please be aware that, as of now, the project only supports OpenAI's GPT-4 or GPT-3.5 as the LLM provider.

Future updates may include support for additional LLM providers. Stay tuned for further developments and enhancements to the Lightweight AGI project.


## Installation
1. Install Go and set up your Go workspace.
2. Clone the repository:

```bash
git clone https://github.com/zawakin/lightweight-agi.git
```

3. Navigate to the repository:
```bash
cd lightweight-agi
```

4. Install the required packages:
```bash
go mod download
```

5. Create a .env file with your OpenAI API key:
```makefile
OPENAI_API_KEY=your_openai_api_key_here
```

## Usage

Run the main program:

```bash
go run main.go
```

The AGI agent will start learning how to play chess by executing tasks, evaluating results, and refining its objectives.

## Sequence Diagram

![Sequence Diagram](./img/lightweight-agi-sequence-diagram.svg)

## Flowchart

![Flowchart](./img/lightweight-agi-flowchart.svg)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT

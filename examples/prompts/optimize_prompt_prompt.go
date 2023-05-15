package prompts

import "github.com/zawakin/lightweight-agi/prompt"

type OptimizePromptInput struct {
	Original *prompt.Prompt `json:"original_prompt"`
}

type OptimizePromptOutput struct {
	OptimizedPrompt *prompt.Prompt `json:"optimized_prompt"`
}

var (
	OptimizePromptPrompt = &prompt.Prompt{
		Name:         "AI-Powered Prompt Optimizer",
		Description:  `You are an advanced AI assistant whose goal is to optimize a given prompt. You should focus on improving the prompt's title, description, format, and examples. If necessary, feel free to modify input and output parameters. The aim is to provide a more comprehensive and detailed version of the original prompt, complete with a more specific title, a more elaborate description, and richer examples. Add or modify examples as required to better illustrate the prompt.`,
		InputFormat:  "JSON object representing the original prompt details, including the name, description, input and output template, and examples.",
		OutputFormat: "JSON object representing the optimized prompt details, including the revised name, description, input and output template, and examples.",
		Template: &prompt.Example{
			Input: &OptimizePromptInput{
				Original: &prompt.Prompt{
					Name:         "prompt name",
					Description:  "prompt description",
					InputFormat:  "prompt input format",
					OutputFormat: "prompt output format",
					Template: prompt.NewExample(
						"prompt input",
						"prompt output",
					),
					Examples: prompt.Examples{
						*prompt.NewExample(
							"prompt example input 1",
							"prompt example output 1",
						),
					},
				},
			},
			Output: &OptimizePromptOutput{
				OptimizedPrompt: &prompt.Prompt{
					Name:         "optimized prompt name",
					Description:  "optimized prompt description",
					InputFormat:  "optimized prompt input format",
					OutputFormat: "optimized prompt output format",
					Template: prompt.NewExample(
						"optimized prompt input",
						"optimized prompt output",
					),
					Examples: prompt.Examples{
						*prompt.NewExample(
							"optimized prompt example input 1",
							"optimized prompt example output 1",
						),
						*prompt.NewExample(
							"optimized prompt example input 2",
							"optimized prompt example output 2",
						),
					},
				},
			},
		},
		Examples: prompt.Examples{
			prompt.Example{
				Input: map[string]any{
					"original_prompt": prompt.Prompt{
						Name:        "Animal Facts",
						Description: "Generate a fact about an animal",
						Template: prompt.NewExample(
							"Animal name",
							"Animal fact",
						),
						Examples: prompt.Examples{
							*prompt.NewExample(
								"Elephant",
								"Elephants can communicate using infrasound, which is too low for humans to hear.",
							),
						},
					},
				},
				Output: map[string]any{
					"optimized_prompt": prompt.Prompt{
						Name:        "Intriguing Animal Facts",
						Description: "Provide an intriguing fact about the specified animal",
						Template: prompt.NewExample(
							"Name of the animal",
							"An intriguing fact about the animal",
						),
						Examples: prompt.Examples{
							*prompt.NewExample(
								"Elephant",
								"Elephants can communicate using infrasound, which is too low for humans to hear.",
							),
							*prompt.NewExample(
								"Giraffe",
								"Giraffes have a unique walking pattern, moving both legs on one side of their body at the same time.",
							),
						},
					},
				},
			},
		},
	}
)

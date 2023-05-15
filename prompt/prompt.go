package prompt

import (
	"encoding/json"
	"fmt"
)

type (
	Input  any
	Output any
)

// Prompter interface defines a common method for generating formatted prompts.
type Prompter interface {
	Format(input Input) (string, error)
}

var _ Prompter = (*Prompt)(nil)

// Prompt is a struct that contains the information of a prompt.
type Prompt struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	InputFormat  string   `json:"input_format"`
	OutputFormat string   `json:"output_format"`
	Template     *Example `json:"template"`
	Examples     Examples `json:"examples"`
}

func NewPrompt(name string, description string, template *Example, examples []Example) *Prompt {
	return &Prompt{
		Name:        name,
		Description: description,
		Template:    template,
		Examples:    examples,
	}
}

func NewSimplePrompt(name string, description string, input string, output string) *Prompt {
	return NewPrompt(name, description, NewExample(input, output), nil)
}

func (c *Prompt) Format(input Input) (string, error) {
	formattedInput, err := toJson(input)
	if err != nil {
		return "", err
	}
	formattedTemplate, err := c.Template.Format()
	if err != nil {
		return "", err
	}
	formattedExamples, err := c.Examples.Format()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`You are an AI named "%s".
%s

Output a JSON-formatted string without outputting any other strings.

Template:
%s

%s

Input: %s
Output:`, c.Name, c.Description, formattedTemplate, formattedExamples, formattedInput), nil
}

type Example struct {
	Input  Input  `json:"input"`
	Output Output `json:"output"`
}

func NewExample(input Input, output Output) *Example {
	return &Example{
		Input:  input,
		Output: output,
	}
}

func (p *Example) Format() (string, error) {
	input, err := toJson(p.Input)
	if err != nil {
		return "", err
	}

	output, err := toJson(p.Output)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Input: %s\nOutput: %s", input, output), nil
}

type Examples []Example

func (p Examples) Format() (string, error) {
	if len(p) == 0 {
		return "", nil
	}

	s := ""
	for _, e := range p {
		s += "Example:\n"
		ds, err := e.Format()
		if err != nil {
			return "", err
		}

		s += ds + "\n"
	}
	return s, nil
}

func toJson(v any) (string, error) {
	s, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

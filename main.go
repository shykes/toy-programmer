package main

import (
	"dagger/toy-programmer/internal/dagger"
	"os"
)

type ToyProgrammer struct{}

// Write a Go program
func (m *ToyProgrammer) GoProgram(assignment string) *dagger.Container {
	// Define env var
	envBytes := []byte("OPENAI_KEY=foobar\nOPENAI_MODEL=blabla")
	// Write to .env
	err := os.WriteFile(".env", envBytes, 0644)
	if err != nil {
									panic(err)
	}
	// Create a new workspace, using third-party module
	before := dag.ToyWorkspace()
	// Run the agent loop in the workspace
	after := dag.Llm().
		WithToyWorkspace(before).
		WithPromptVar("assignment", assignment).
		WithPromptFile(dag.CurrentModule().Source().File("prompt.txt")).
		Loop().
		ToyWorkspace()
	// Return the modified workspace's container
	return after.Container()
}

package main

import (
	"dagger/toy-programmer/internal/dagger"
)

type ToyProgrammer struct{}

// Write a Go program
func (m *ToyProgrammer) GoProgram(assignment string) *dagger.Container {
	// Create a new workspace, using third-party module
	before := dag.ToyWorkspace()
	// Run the agent loop in the workspace
	after := dag.Llm().
		WithToyWorkspace(before).
		WithPromptVar("assignment", assignment).
		WithPromptFile(dag.CurrentModule().Source().File("prompt.txt")).
		ToyWorkspace()
	// Return the modified workspace's container
	return after.Container()
}

package main

import (
	"dagger/toy-programmer/internal/dagger"
)

type ToyProgrammer struct{}

// Write a Go program
func (m *ToyProgrammer) GoProgram(
	// The programming assignment
	// Example: "write me a curl clone"
	assignment string,
	// Run LLM-powered QA on the result
	// +optional
	qa bool) *dagger.Container {
	ws := dag.ToyWorkspace().Write("assignment.txt", assignment)

	env := dag.Env().
		WithToyWorkspaceInput("workspace", ws, "The workspace for the assignment").
		WithToyWorkspaceOutput("workspace", "The workspace with the results of the assignment")

	result := dag.LLM().
		WithEnv(env).
		WithPrompt("You are an expert go programmer. You have access to a workspace").
		WithPrompt("Complete the assignment written at assignment.txt").
		WithPrompt("Don't stop until the code builds").
		Env().
		Output("workspace").
		AsToyWorkspace().
		Container()
	if qa {
		qaEnv := dag.Env().
			WithContainerInput("results", result, "The container with the results of the assignment").
			WithContainerOutput("results", "The container with the QA results")

		return dag.LLM().
			WithEnv(qaEnv).
			WithPrompt("You are an expert QA engineer. You have access to a container").
			WithPrompt("There is a go program in the current directory. Build and run it. Understand what it does. Write your findings in QA.md").
			WithPrompt("Include a table of each command you ran, and the result").
			WithPrompt("Be careful not to wipe the state of the container with a new image. Focus on withExec, file, directory").
			Env().
			Output("results").
			AsContainer()
	}

	return result
}

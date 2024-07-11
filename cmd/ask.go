package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

type askOptions struct {
	modelName string
	path      string
	question  string
}

var askOpts askOptions

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question about the specified image",
	RunE:  runAsk,
}

func init() {
	rootCmd.AddCommand(askCmd)

	flags := askCmd.Flags()
	flags.StringVarP(&askOpts.modelName, "model", "m", "llava:13b", "Model to use")
	flags.StringVarP(&askOpts.path, "file", "f", "", "Path to image file")
	flags.StringVarP(&askOpts.question, "question", "q", "", "Question to ask")

	askCmd.MarkFlagRequired("file")
	askCmd.MarkFlagRequired("question")
}

func runAsk(cmd *cobra.Command, args []string) error {
	imgData, err := os.ReadFile(askOpts.path)
	if err != nil {
		return err
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	req := &api.GenerateRequest{
		Model:  askOpts.modelName,
		Prompt: askOpts.question,
		Images: []api.ImageData{imgData},
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		// In streaming mode, responses are partial so we call fmt.Print (and not
		// Println) in order to avoid spurious newlines being introduced. The
		// model will insert its own newlines if it wants.
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		return err
	}
	fmt.Println()

	return nil
}

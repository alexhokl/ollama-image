package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

type summarizeOptions struct {
	modelName string
	path      string
}

var summarizeOpts summarizeOptions

// summarizeCmd represents the summarize command
var summarizeCmd = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize an image",
	RunE:  runSummarize,
}

func init() {
	rootCmd.AddCommand(summarizeCmd)

	flags := summarizeCmd.Flags()
	flags.StringVarP(&summarizeOpts.modelName, "model", "m", "llava:13b", "Model to use")
	flags.StringVarP(&summarizeOpts.path, "file", "f", "", "Path to image file")

	summarizeCmd.MarkFlagRequired("file")
}

func runSummarize(cmd *cobra.Command, args []string) error {
	imgData, err := os.ReadFile(summarizeOpts.path)
	if err != nil {
		return err
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	req := &api.GenerateRequest{
		Model:  summarizeOpts.modelName,
		Prompt: "describe this image",
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

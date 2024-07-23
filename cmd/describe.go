package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

type describeOptions struct {
	modelName string
	path      string
}

var describeOpts describeOptions

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe an image",
	RunE:  rundescribe,
}

func init() {
	rootCmd.AddCommand(describeCmd)

	flags := describeCmd.Flags()
	flags.StringVarP(&describeOpts.modelName, "model", "m", "llava:13b", "Model to use")
	flags.StringVarP(&describeOpts.path, "file", "f", "", "Path to image file")

	describeCmd.MarkFlagRequired("file")
}

func rundescribe(cmd *cobra.Command, args []string) error {
	imgData, err := os.ReadFile(describeOpts.path)
	if err != nil {
		return err
	}

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return err
	}

	req := &api.GenerateRequest{
		Model:  describeOpts.modelName,
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

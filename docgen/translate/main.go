package main

import (
	"context"
	"fmt"
	"github.com/avast/retry-go"
	"github.com/sashabaranov/go-openai"
	"os"
	"path/filepath"
	"strings"
)
import _ "github.com/joho/godotenv/autoload"

func main() {
	token := os.Getenv("OPENAI_TOKEN")
	client := openai.NewClient(token)
	ctx := context.Background()
	languages := [][2]string{{
		"zh", "Chinese",
	}}
	for _, language := range languages {
		dir := language[0]
		lang := language[1]
		filepath.Walk("./docs/en", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				return nil
			}
			fmt.Printf("Working on %s\n", path)
			content, err := os.ReadFile(path)
			if err != nil {
				panic(err)
			}
			lines := strings.Split(string(content), "\n")
			chunkSize := 150
			numChunks := (len(lines) + chunkSize - 1) / chunkSize
			chunks := make([][]string, numChunks)
			// Split the lines into chunks
			for i, line := range lines {
				chunkIndex := i / chunkSize
				chunks[chunkIndex] = append(chunks[chunkIndex], line)
			}
			outPath := filepath.Join("./docs", dir, path[len("docs/en/"):])
			if _, err := os.Stat(outPath); err == nil {
				fmt.Printf("Skipping %s\n", outPath)
				return nil
			}
			results := make([]string, len(chunks))
			for i, chunk := range chunks {
				content := strings.Join(chunk, "\n")
				messages := []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleSystem,
						Content: "You will be helping me to translate technical documentation from English into " + lang +
							". Make your best effort to use the most natural and ideological form, including changing the sequence of sentences, use different words that are more natural." +
							"You should also use your best judgement to identify terminologies that should keep English form." +
							"The document is using markdown format and may contain some special characters that works on Gitbook. Do not translate those special characters or change the format." +
							"The document may also contain CLI usage examples, in which case, you should only translate the usage text, not the command or the arguments." +
							"For example, if you see code block (wrapped in ```), you should not translate the command name or the argument usage but only the comments or the description text.",
					}, {
						Role:    openai.ChatMessageRoleUser,
						Content: content,
					},
				}

				var response openai.ChatCompletionResponse
				err = retry.Do(func() error {
					var err error
					response, err = client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
						Model:    openai.GPT3Dot5Turbo,
						Messages: messages,
					})
					if err != nil {
						return err
					}
					return nil
				}, retry.RetryIf(func(err error) bool {
					return strings.Contains(err.Error(), "429")
				}))
				if err != nil {
					panic(err)
				}
				results[i] = response.Choices[0].Message.Content
			}
			fmt.Printf("Writing to %s\n", outPath)
			err = os.MkdirAll(filepath.Dir(outPath), 0755)
			if err != nil {
				panic(err)
			}
			err = os.WriteFile(outPath, []byte(strings.Join(results, "\n")), 0644)
			if err != nil {
				panic(err)
			}
			return nil
		})
	}
}

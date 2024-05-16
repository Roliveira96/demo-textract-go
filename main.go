package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"os"
)

func main() {
	ctx := context.TODO()
	cfg, err := loadAws(ctx)
	if err != nil {
		panic(err)
	}
	tex := textract.NewFromConfig(cfg)

	file, err := os.ReadFile("./teste.pdf")
	if err != nil {
		panic(err)
	}

	output, err := tex.DetectDocumentText(ctx, &textract.DetectDocumentTextInput{
		Document: &types.Document{
			Bytes: file,
		},
	})

	if err != nil {
		panic(err)
	}

	generateJsonOutput(err, output)

}

func generateJsonOutput(err error, output *textract.DetectDocumentTextOutput) {
	jsonData, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}

func loadAws(ctx context.Context) (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return aws.Config{}, fmt.Errorf("configuration error: %w", err)
	}
	return cfg, nil
}

package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	region = "us-west-2"
)

type Config struct {
	BucketName string
	TableName  string
}

func getDynamodbClient() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)}, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return dynamodb.New(sess)
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func getTerraformExecutablePath() string {
	return filepath.Join(os.Getenv("HOME"), ".terraform", "bin", "terraform")
}

func getTerraformVersion() (string, error) {
	exePath := getTerraformExecutablePath()
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		return "", fmt.Errorf("terraform not installed")
	}
	cmd := exec.Command(exePath, "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func generateTableName(table_name string, env string) string {
	return fmt.Sprintf("%s_%s", table_name, env)
}
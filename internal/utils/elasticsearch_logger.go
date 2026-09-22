package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type LogEntry struct {
	Timestamp   time.Time              `json:"timestamp"`
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	Error       string                 `json:"error,omitempty"`
	StackTrace  string                 `json:"stack_trace,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	ServiceName string                 `json:"service_name"`
}

type ESLogger struct {
	client      *elasticsearch.Client
	indexPrefix string
	serviceName string
}

func LogError(err error, route, method string) {
	metadata := map[string]interface{}{
		"endpoint":    route,
		"http_method": method,
	}
	logger, loggerErr := NewESLogger()
	if loggerErr != nil {
		log.Printf("Failed to initialize ESLogger: %v", loggerErr)
		return
	}
	if logErr := logger.LogError(
		context.Background(),
		err,
		"Failed to process user request",
		metadata,
	); logErr != nil {
		log.Printf("Failed to log error: %v", logErr)
	}
}

func NewESLogger() (*ESLogger, error) {
	config := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch:9200", // Use your Elasticsearch address
		},
		// Add other configuration options as needed
		// Username: "user",
		// Password: "pass",
	}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error creating elasticsearch client: %w", err)
	}

	return &ESLogger{
		client:      client,
		indexPrefix: "app-logs",
		serviceName: "service-name",
	}, nil
}

func (l *ESLogger) LogError(ctx context.Context, err error, message string, metadata map[string]interface{}) error {
	if metadata == nil {
		metadata = make(map[string]interface{})
	}

	entry := LogEntry{
		Timestamp:   time.Now().UTC(),
		Level:       "error",
		Message:     message,
		Error:       err.Error(),
		Metadata:    metadata,
		ServiceName: l.serviceName,
	}

	// Create index name with date suffix for better organization
	indexName := fmt.Sprintf("%s-%s", l.indexPrefix, time.Now().UTC().Format("2006.01.02"))

	// Convert entry to JSON
	payload, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("error marshaling log entry: %w", err)
	}

	// Index the document
	res, err := l.client.Index(
		indexName,
		strings.NewReader(string(payload)),
		l.client.Index.WithContext(ctx),
		l.client.Index.WithRefresh("true"),
	)

	if err != nil {
		return fmt.Errorf("error indexing log entry: %w", err)
	}
	defer res.Body.Close()

	// Check for indexing errors
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.String())
	}

	return nil
}

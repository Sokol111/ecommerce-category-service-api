package events

import (
	"fmt"
	"sync"

	"github.com/hamba/avro/v2"
)

var (
	// Pre-parsed schemas (initialized once at startup)
	compiledSchemas     = make(map[string]avro.Schema)
	compiledSchemasOnce sync.Once
	schemaInitError     error
)

// initSchemas parses all embedded schemas once
func initSchemas() error {
	compiledSchemasOnce.Do(func() {
		schemas := map[string][]byte{
			"CategoryCreatedEvent": CategoryCreatedSchema,
			"CategoryUpdatedEvent": CategoryUpdatedSchema,
			"EventMetadata": EventMetadataSchema,
		}

		for name, schemaJSON := range schemas {
			schema, err := avro.Parse(string(schemaJSON))
			if err != nil {
				schemaInitError = fmt.Errorf("failed to parse schema %s: %w", name, err)
				return
			}
			compiledSchemas[name] = schema
		}
	})
	return schemaInitError
}

// Marshal serializes an event using its pre-compiled schema
func Marshal(schemaName string, v interface{}) ([]byte, error) {
	if err := initSchemas(); err != nil {
		return nil, err
	}
	
	schema, ok := compiledSchemas[schemaName]
	if !ok {
		return nil, fmt.Errorf("schema %s not found", schemaName)
	}
	
	return avro.Marshal(schema, v)
}

// Unmarshal deserializes an event using its pre-compiled schema
func Unmarshal(schemaName string, data []byte, v interface{}) error {
	if err := initSchemas(); err != nil {
		return err
	}
	
	schema, ok := compiledSchemas[schemaName]
	if !ok {
		return fmt.Errorf("schema %s not found", schemaName)
	}
	
	return avro.Unmarshal(schema, data, v)
}

// Typed serialization helpers

// MarshalCategoryCreatedEvent serializes CategoryCreatedEvent to Avro bytes
func MarshalCategoryCreatedEvent(event *CategoryCreatedEvent) ([]byte, error) {
	return Marshal("CategoryCreatedEvent", event)
}

// UnmarshalCategoryCreatedEvent deserializes CategoryCreatedEvent from Avro bytes
func UnmarshalCategoryCreatedEvent(data []byte) (*CategoryCreatedEvent, error) {
	var event CategoryCreatedEvent
	if err := Unmarshal("CategoryCreatedEvent", data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// MarshalCategoryUpdatedEvent serializes CategoryUpdatedEvent to Avro bytes
func MarshalCategoryUpdatedEvent(event *CategoryUpdatedEvent) ([]byte, error) {
	return Marshal("CategoryUpdatedEvent", event)
}

// UnmarshalCategoryUpdatedEvent deserializes CategoryUpdatedEvent from Avro bytes
func UnmarshalCategoryUpdatedEvent(data []byte) (*CategoryUpdatedEvent, error) {
	var event CategoryUpdatedEvent
	if err := Unmarshal("CategoryUpdatedEvent", data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}


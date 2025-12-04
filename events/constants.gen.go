package events

import "reflect"

// Event type constants - single source of truth for event names
// These match the schema names and should be used in EventMetadata.EventType
const (
	EventTypeCategoryCreated = "CategoryCreatedEvent"
	EventTypeCategoryUpdated = "CategoryUpdatedEvent"
)

// Event is a sealed interface that all event types must implement.
// This enables exhaustiveness checking when using type switches.
//
// When using exhaustive linter (golangci-lint), it will warn you if you
// don't handle all concrete Event types in a type switch.
type Event interface {
	isEvent() // unexported method seals the interface
}

// CategoryCreatedEvent implements Event interface
func (*CategoryCreatedEvent) isEvent() {}

// CategoryUpdatedEvent implements Event interface
func (*CategoryUpdatedEvent) isEvent() {}

// Topic constants - Kafka topics defined in AsyncAPI spec
const (
	TopicCatalogCategoryEvents = "catalog.category.events"
)

// Schema name constants - Avro schema full names (namespace.name)
const (
	SchemaNameCategoryCreated = "com.ecommerce.events.category.CategoryCreatedEvent"
	SchemaNameCategoryUpdated = "com.ecommerce.events.category.CategoryUpdatedEvent"
)

// TypeRegistrations contains schema registration data for TypeMapping.
// Use this to register all schemas in your microservice:
//
// Example:
//   func RegisterSchemas(tm *mapping.TypeMapping) error {
//       for _, reg := range events.TypeRegistrations {
//           if err := tm.Register(reg.GoType, reg.SchemaJSON, reg.SchemaName, reg.Topic); err != nil {
//               return err
//           }
//       }
//       return nil
//   }
type TypeRegistration struct {
	GoType     reflect.Type
	SchemaJSON []byte
	SchemaName string
	Topic      string
}

var TypeRegistrations = []TypeRegistration{
	{
		GoType:     reflect.TypeOf(CategoryCreatedEvent{}),
		SchemaJSON: CategoryCreatedSchema,
		SchemaName: SchemaNameCategoryCreated,
		Topic:      TopicCatalogCategoryEvents,
	},
	{
		GoType:     reflect.TypeOf(CategoryUpdatedEvent{}),
		SchemaJSON: CategoryUpdatedSchema,
		SchemaName: SchemaNameCategoryUpdated,
		Topic:      TopicCatalogCategoryEvents,
	},
}

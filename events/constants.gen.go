package events

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

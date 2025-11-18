package events

import _ "embed"

//go:embed asyncapi.yaml
var AsyncAPISpec []byte

//go:embed schemas/category_created.avsc
var CategoryCreatedSchema []byte

//go:embed schemas/category_updated.avsc
var CategoryUpdatedSchema []byte

//go:embed schemas/event_metadata.avsc
var EventMetadataSchema []byte

// Combined schemas with EventMetadata embedded for Avro serialization
// These resolve named type references like "com.ecommerce.events.EventMetadata"

//go:embed schemas/category_created_combined.json
var CategoryCreatedCombinedSchema []byte // Use this for Serialize() - includes EventMetadata

//go:embed schemas/category_updated_combined.json
var CategoryUpdatedCombinedSchema []byte // Use this for Serialize() - includes EventMetadata


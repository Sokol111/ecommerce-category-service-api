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

// Combined schemas with EventMetadata inlined for Avro serialization
// These schemas have the EventMetadata type definition embedded inline

//go:embed schemas/category_created_combined.json
var CategoryCreatedCombinedSchema []byte // Use this for Serialize() - has EventMetadata inlined

//go:embed schemas/category_updated_combined.json
var CategoryUpdatedCombinedSchema []byte // Use this for Serialize() - has EventMetadata inlined


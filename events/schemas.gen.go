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


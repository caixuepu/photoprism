package entity

import "github.com/photoprism/photoprism/internal/classify"

const (
	// Data sources.
	SrcAuto     = ""
	SrcManual   = "manual"
	SrcEstimate = "estimate"
	SrcName     = "name"
	SrcMeta     = "meta"
	SrcXmp      = "xmp"
	SrcYaml     = "yaml"
	SrcLocation = classify.SrcLocation
	SrcImage    = classify.SrcImage

	// Sort orders.
	SortOrderRelevance = "relevance"
	SortOrderNewest    = "newest"
	SortOrderOldest    = "oldest"
	SortOrderImported  = "imported"
	SortOrderSimilar   = "similar"
	SortOrderName      = "name"

	// Unknown values.
	YearUnknown  = -1
	MonthUnknown = -1
	TitleUnknown = "Unknown"

	// Content types.
	TypeDefault = ""
	TypeImage   = "image"
	TypeLive    = "live"
	TypeVideo   = "video"
	TypeRaw     = "raw"
	TypeText    = "text"

	// Root directories.
	RootOriginals = ""
	RootExamples  = "examples"
	RootSidecar   = "sidecar"
	RootImport    = "import"
	RootPath      = "/"

	// Event names.
	Updated = "updated"
	Created = "created"
	Deleted = "deleted"
)

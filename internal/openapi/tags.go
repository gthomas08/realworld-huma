package openapi

type Tag string

const (
	AuthTag Tag = "Auth"
	UserTag Tag = "User"
)

// AddTags takes a variable number of tags and returns a slice of strings.
// It is intended to be used with OpenAPI tags.
func AddTags(tags ...Tag) []string {
	var result []string
	for _, tag := range tags {
		result = append(result, string(tag))
	}
	return result
}

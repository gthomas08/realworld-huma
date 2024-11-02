package security

import "github.com/danielgtaylor/huma/v2"

type AuthFlow string

const Bearer AuthFlow = "bearer"

// Protected returns a security requirement for a given flow.
func Protected(flow AuthFlow) []map[string][]string {
	return []map[string][]string{
		{string(flow): {}},
	}
}

// Public returns a security requirement for no authentication.
func Public() []map[string][]string {
	return []map[string][]string{}
}

func GetSecuritySchemes() map[string]*huma.SecurityScheme {
	return map[string]*huma.SecurityScheme{
		string(Bearer): {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
}

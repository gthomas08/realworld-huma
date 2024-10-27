package security

import "github.com/danielgtaylor/huma/v2"

type AuthFlow string

const Bearer AuthFlow = "bearer"

// RequireAuth returns a security requirement for a given flow.
func RequireAuth(flow AuthFlow) []map[string][]string {
	return []map[string][]string{
		{string(flow): {}},
	}
}

// RequireNoAuth returns a security requirement for no authentication.
func RequireNoAuth() []map[string][]string {
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

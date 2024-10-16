package security

type AuthFlow string

const Bearer AuthFlow = "bearer"

// RequireAuth returns a security requirement for a given flow.
func RequireAuth(flow AuthFlow) []map[string][]string {
	return []map[string][]string{
		{string(flow): {}},
	}
}

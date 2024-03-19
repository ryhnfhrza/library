package exception

type Conflict struct {
	Error string
}

func NewConflictError(error string) Conflict {
	return Conflict{Error: error}
}

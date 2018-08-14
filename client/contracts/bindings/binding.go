package bindings

// Binding represents the golang binding of a specific contract
type Binding interface {
	Domain() string
}

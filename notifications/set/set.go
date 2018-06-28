package set

type Set interface {
	Add(key string) error
	Remove(key string) error
	Contains(key string) (bool, error)
}

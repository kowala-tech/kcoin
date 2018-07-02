package notifier

const (
	EmailToKey   = "TO"
	EmailFromKey = "FROM"
)

//go:generate moq -out notifier_mock.go . Notifier
type Notifier interface {
	Send(vars map[string]string) error
}

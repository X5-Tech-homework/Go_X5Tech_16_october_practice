package practice

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	Channel string
}

func (b BaseNotifier) Label(message string) string {
	return "[" + b.Channel + "] " + message
}

type EmailNotifier struct {
	BaseNotifier
	Address  string
	Messages []string
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
	Messages []string
}

func (n *EmailNotifier) Send(message string) error {
	// TODO: добавьте форматированное сообщение в Messages.
	output := n.Label(message)
	n.Messages = append(n.Messages, output)
	return nil
}

func (n *PushNotifier) Send(message string) error {
	// TODO: добавьте форматированное сообщение в Messages.
	output := n.Label(message)
	n.Messages = append(n.Messages, output)
	return nil
}

func Dispatch(all []Notifier, message string) error {
	// TODO: вызовите Send у каналов по порядку, верните первую ошибку.
	for _, n := range all {
		if err := n.Send(message); err != nil {
			return err
		}
	}
	return nil
}

func EmailAddress(n Notifier) (string, bool) {
	// TODO: безопасно проверьте тип и указатель.
	if email, ok := n.(*EmailNotifier); ok {
		if email != nil {
			return email.Address, true
		}
	}
	return "", false
}

func Describe(n Notifier) string {
	// TODO: опишите email и push, остальные значения дают unknown.
	switch v := n.(type) {
	case *EmailNotifier:
		if v == nil {
			return "unknown"
		}
		return "email: " + v.Address
	case *PushNotifier:
		if v == nil {
			return "unknown"
		}
		return "push: " + v.DeviceID
	default:
		return "unknown"
	}
}

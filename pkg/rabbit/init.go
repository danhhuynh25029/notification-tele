package rabbit

type RabbitMessage interface {
	InitRabbit() error
}

func InitRabbit() error {
	return nil
}

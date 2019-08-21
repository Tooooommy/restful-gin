package connect

type Broker interface {
	Publish(topic string, msg *Message)
	Subscribe(topic string, handleFunc func(msg *Message))
	Unsubscribe(topic string)
}

type RedisEngine struct {
}

func NewEngine() *RedisEngine {
	return &RedisEngine{}
}

func (engine *RedisEngine) Publish(topic string, msg *Message) {
}

func (engine *RedisEngine) Subscribe(topic string, handleFunc func(msg *Message)) {
}

func (engin *RedisEngine) Unsubscribe(topic string) {

}

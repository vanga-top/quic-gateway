package gateway

type Pipeline interface {
	Chain() []Module
}

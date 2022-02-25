package gateway

type Pipeline interface {
	First() Module
	List() Module
	Chain() []Module
	// Add return true/ false if module id is same
	Add(Module) (bool, string)
	//Remove param:ID module ID
	Remove(ID string) bool
}

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

type PipelineInstance struct {
	moduleChain []Module
}

func (p PipelineInstance) First() Module {
	//TODO implement me
	panic("implement me")
}

func (p PipelineInstance) List() Module {
	//TODO implement me
	panic("implement me")
}

func (p PipelineInstance) Chain() []Module {
	return p.moduleChain
}

func (p PipelineInstance) Add(module Module) (bool, string) {
	return true, "ok"
}

func (p PipelineInstance) Remove(ID string) bool {
	//TODO implement me
	panic("implement me")
}

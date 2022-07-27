package group

type Producer interface {
	Run() error
	Done() <-chan struct{}
}

type producer struct {
	t    thread
	done chan struct{}
}

func (p *producer) Run() error {
	go p.runner()
	return nil
}

func (p *producer) Done() <-chan struct{} {
	return p.done
}

// thread = chan any
func NewProducer(t thread) Producer {
	return &producer{
		t:    t,
		done: make(chan struct{}),
	}
}

func (p *producer) runner() {

}

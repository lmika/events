package events

type Dispatcher struct {
	topics map[string]*topic
}

func New() *Dispatcher {
	return &Dispatcher{topics: make(map[string]*topic)}
}

func (d *Dispatcher) On(event string, receiver interface{}) error {
	// TODO: make thread safe
	t, hasTopic := d.topics[event]
	if !hasTopic {
		t = &topic{}
		d.topics[event] = t
	}

	sub, err := newSubscriptionFromFunc(receiver)
	if err != nil {
		return err
	}

	t.addSubscriber(sub)
	return nil
}

func (d *Dispatcher) Fire(event string, args ...interface{}) {
	// TODO: make thead safe
	topic, hasTopic := d.topics[event]
	if !hasTopic {
		return
	}

	preparedArgs := prepareArgs(args)

	for sub := topic.head; sub != nil; sub = sub.next {
		sub.handler.invoke(preparedArgs)
	}
}

type topic struct {
	head *subscription
	tail *subscription
}

func (t *topic) addSubscriber(sub *subscription) {
	if t.head == nil {
		t.head = sub
	}
	if t.tail != nil {
		t.tail.next = sub
	}
	t.tail = sub
}

type subscription struct {
	handler receiptHandler
	next    *subscription
}

func newSubscriptionFromFunc(fn interface{}) (*subscription, error) {
	handler, err := newReceiptHandler(fn)
	if err != nil {
		return nil, err
	}
	return &subscription{handler: handler, next: nil}, nil
}
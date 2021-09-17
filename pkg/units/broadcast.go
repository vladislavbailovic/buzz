package units

import "sync"

type Consumer interface {
	Listen(...interface{})
}

type Listener func(...interface{})

type EventType string

type Broadcaster struct {
	sync.Mutex
	consumers map[EventType][]Listener
}

func (cast *Broadcaster) Subscribe(event EventType, listener Consumer) {
	cast.Lock()
	defer cast.Unlock()

	if cast.consumers == nil {
		cast.consumers = make(map[EventType][]Listener)
	}
	cast.consumers[event] = append(cast.consumers[event], listener.Listen)
}

func (cast Broadcaster) Publish(event EventType, what ...interface{}) {
	cast.Lock()
	defer cast.Unlock()

	listeners, ok := cast.consumers[event]
	if ok {
		var wg sync.WaitGroup
		wg.Add(len(listeners))
		for _, listener := range listeners {
			go func(handler Listener) {
				defer wg.Done()
				handler(what...)
			}(listener)
		}
		wg.Wait()
	}
}

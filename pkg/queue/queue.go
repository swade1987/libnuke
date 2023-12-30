package queue

type IQueue interface {
	Total() int
	Count(states ...ItemState) int
}

type Queue struct {
	Items []Item
}

func (q Queue) Total() int {
	return len(q.Items)
}

func (q Queue) Count(states ...ItemState) int {
	count := 0
	for _, item := range q.Items {
		for _, state := range states {
			if item.State == state {
				count = count + 1
				break
			}
		}
	}
	return count
}

package hw4

type List struct {
	first *Item
	last  *Item
	len   int
}

func (list *List) PushFront(v int) {
	newItem := &Item{value: v}

	if list.first == nil {
		list.first = newItem
		list.last = newItem
	} else {
		list.pushBefore(list.first, newItem)
	}

	list.len++
}

func (list *List) PushBack(v int) {
	newItem := &Item{value: v}

	if list.last == nil {
		list.first = newItem
		list.last = newItem
	} else {
		list.pushAfter(list.last, newItem)
	}

	list.len++
}

func (list *List) First() *Item {
	return list.first
}

func (list *List) Last() *Item {
	return list.last
}

func (list *List) Remove(item *Item) {
	if item.prev == nil {
		list.first = item.next
	} else {
		item.prev.next = item.next
	}

	if item.next == nil {
		list.last = item.prev
	} else {
		item.next.prev = item.prev
	}
}

func (list *List) Len() int {
	return list.len
}

func (list *List) pushAfter(item *Item, newItem *Item) {
	newItem.prev = item

	if item.next == nil {
		newItem.next = nil
		list.last = newItem
	} else {
		newItem.next = item.next
		item.next.prev = newItem
	}

	item.next = newItem
}

func (list *List) pushBefore(item *Item, newItem *Item) {
	newItem.next = item

	if item.prev == nil {
		newItem.prev = nil
		list.first = newItem
	} else {
		newItem.prev = item.prev
		item.prev.next = newItem
	}

	item.prev = newItem
}

type Item struct {
	value int
	next  *Item
	prev  *Item
}

func (item *Item) Value() int {
	return item.value
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}

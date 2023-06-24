package utils

type txtnode struct {
	tid   int
	title string
	prev  *txtnode
	next  *txtnode
}

type list struct {
	front *txtnode
	back  *txtnode
	len   int
}

// 新建双向链表
func NewList() *list {
	f, b := &txtnode{}, &txtnode{}
	f.next = b
	b.prev = f
	l := &list{
		front: f,
		back:  b,
		len:   0,
	}
	return l
}

// 命中缓存：将指定的节点从原来的位置移到链表头部
func (l *list) move2Front(tn *txtnode) {
	if tn.prev != l.front {
		// 从原来的地方解开
		prevn, nextn := tn.prev, tn.next
		prevn.next = nextn
		nextn.prev = prevn
		// 放到头部
		l.pushFront(tn)
	}

}

// 将指定节点插入链表头部
func (l *list) pushFront(tn *txtnode) {
	secondn, frontn := l.front.next, l.front
	tn.next, tn.prev = secondn, frontn
	secondn.prev = tn
	frontn.next = tn
}

// 缓存已满：将最近最少使用的内容删除——将链表尾部的节点移除
func (l *list) delBack() (tid int, title string) {
	tid, title = l.back.prev.tid, l.back.prev.title
	second2lastn := l.back.prev.prev
	second2lastn.next = l.back
	l.back.prev = second2lastn
	return
}

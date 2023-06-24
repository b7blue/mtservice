package utils

import (
	"os"
	"sync"
)

// 要写并发安全，加锁。。
// 链表和哈希表要不要分开加锁？
var mu sync.Mutex

var lruList *list
var lruMap map[int]*txtnode

const cacheCap = 100

func init() {
	lruList = NewList()
	lruMap = make(map[int]*txtnode, cacheCap)
}

// 有一个问题：
// 第一次请求获取文章a的TXT，下载到服务器本地，存入缓存
// 然后文章a发生更新
// 第二次请求获取文章a，用tid查询缓存发现已经在本地，直接拿到服务器本地的txt
// 但是这样就无法获取到更新内容了。怎么办？
// 我可以建议用户只下载完结的文吗，哈哈哈哈哈哈
func NewFile(tid int, title string) {
	mu.Lock()
	defer mu.Unlock()

	tn := &txtnode{
		tid:   tid,
		title: title,
	}
	// 假如缓存没满
	if lruList.len < cacheCap {
		// 链表长度加一
		lruList.len++
	} else {
		// 缓存已满：删除最近最少使用的缓存，删除哈希表key-value
		tid, title = lruList.delBack()
		delete(lruMap, tid)
		os.Remove("txt/" + title + ".txt")
	}
	// 直接插入链表，新增哈希表
	lruList.pushFront(tn)
	lruMap[tid] = tn

}

func GetFilebyTid(tid int) string {
	mu.Lock()
	defer mu.Unlock()

	if tn, ok := lruMap[tid]; ok {
		lruList.move2Front(tn)
		return tn.title
	} else {
		return ""
	}
}

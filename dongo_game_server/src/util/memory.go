package util

import (
	"fmt"
	"time"
)

var DefMemory *Memory

type Memory struct {
	Name  string       //内存容器名
	Items []MemoryItem //多个项
}

type MemoryItem struct {
	Ct   int64       //创建时间
	Et   int64       //过期时间
	Item interface{} //扩展变量
}

func DefaultMemory(name string) *Memory {
	if DefMemory == nil {
		DefMemory = NewMemory(name)
	}
	return DefMemory
}

func NewMemory(name string) *Memory {
	t := &Memory{
		Name: name,
	}
	go func() {
		for {
			t.Repair(time.Now().Unix())
			time.Sleep(time.Second)
		}
	}()
	return t
}

func (p *Memory) Put(newItem interface{}) bool {
	now := time.Now().Unix()
	fmt.Println(now)
	timeQueueItem := MemoryItem{
		Ct:   now,
		Et:   now + 2,
		Item: newItem,
	}
	if len(p.Items) == 0 {
		p.Items = append(p.Items, timeQueueItem)
		return true
	} else {
		for _, item := range p.Items {
			if item.Item == newItem {
				fmt.Println("newItem 存在")
				return true
			}
		}
		p.Items = append(p.Items, timeQueueItem)
		return true
	}
}

func (p *Memory) Delete(newItem interface{}) bool {
	var newItems []MemoryItem
	if len(p.Items) == 0 {
		return true
	} else {
		for _, item := range p.Items {
			if item.Item != newItem {
				newItems = append(newItems, item)
			}
		}
		if len(newItems) > 0 && len(newItems) != len(p.Items) {
			p.Items = newItems
		}
		if len(newItems) == 0 && len(p.Items) == 1 {
			p.Items = newItems
		}
		return true
	}
}

func (p *Memory) Get() interface{} {
	if len(p.Items) == 0 {
		fmt.Println("no records")
		return nil
	} else {
		item := p.Items[0].Item
		p.Delete(item)
		return item
	}
}

func (p *Memory) Length() int {
	fmt.Println(len(p.Items))
	return len(p.Items)
}

func (p *Memory) Repair(time int64) {
	for p.Length() > 0 {
		item := p.Items[0]
		if item.Et > time {
			return
		}
		p.Delete(item.Item)
		fmt.Println("delete后items的长度", p.Length())
	}
}

package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

//初始化
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

//增加元素
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	} else {
		return false
	}
}

//删除元素
func (set *HashSet) Delete(e interface{}) {
	delete(set.m, e)
}

//清空Set
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

//查询是否存在某元素
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//Set长度
func (set *HashSet) Len() int {
	return len(set.m)
}

//比较两个Set是否相同
func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
 	}
 	return true
}

//保存Set在某一个时刻的映像于快照中用于后续的迭代
func (set *HashSet) Element() []interface{} {
	initLen := len(set.m)
	actLen := 0
	snapshot := make([]interface{}, initLen)
	for key := range set.m {
		if actLen < initLen {
			snapshot[actLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actLen++
	}
	//如果在迭代完成前，m的值中的元素有所减少，致使快照值的尾部存在若干nil应该去掉
	if actLen < initLen {
		snapshot = snapshot[:actLen]
	}
	return snapshot
}

//获取自身字符串的表现形式
func (set *HashSet) String() string {
	var buf = bytes.Buffer{}
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first == true{
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}









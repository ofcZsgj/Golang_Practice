package set

//有多个集合类型的时候，应该抽取一个接口类型以标识它们共有的行为方法
type Set interface {
	Add(e interface{}) bool
	Delete(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	//不能在接口类型的方法中包含它的实现类型，即修改Same方法前面让*HashSet类型成为Set接口类型的一个实现
	Same(other Set) bool//func (set *HashSet) Same(other *HashSet) bool
	Element() []interface{}
	String() string
}

//高级方法，核心逻辑是通过对实现了基本方法的组合来实现
//将高级方法抽离出来时间
//判断集合 one 是否是集合 other 的超集
func IsSuperset(one Set, other Set) bool {
	if other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Element() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}
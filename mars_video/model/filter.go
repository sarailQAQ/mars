package model

//实现一个筛，用于对包含某些特定内容的弹幕进行过滤,以及弹幕密度控制
//用AC自动机提高查询效率

type Filter struct {
	set *AC_Automation
}

func (filter *Filter) Add(word string) bool {
	filter.set.Add(word)
	return true
}

func (filter *Filter) Build()  {
	filter.set.Build()
}

func (filter *Filter) Check(word string) bool {
	return !filter.set.Find(word)
}

func NewFilter() *Filter{
	return &Filter{
		set: NewAC(),
	}
}
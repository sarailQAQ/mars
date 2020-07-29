package model

//实现AC自动机，一种高效的字符匹配算法
//修改下build的写法可以取消失配指针，将失配指针加入到儿子节点中
//时间问题就先不实现了

type AC_Automation struct {
	root *node
}

//trie node
type node struct {
	val uint
	ch map[rune]*node

	fail *node //失配指针
}

func NewAC() *AC_Automation{
	return &AC_Automation{
		root: &node{
			val:  0,
			ch:   make(map[rune]*node),
			fail: nil,
		},
	}
}

func (n *node) insert(pattern string,v uint) {
	if v == 0 {
		return
	}

	chars := []rune(pattern)
	p := n
	ok := true
	for _,s := range chars {
		_,ok = p.ch[s]
		if !ok {
			p.ch[s] = &node{
				val: 0,
				ch:  make(map[rune]*node),
				fail:nil,
			}
		}

		p = p.ch[s]
	}

	p.val = v
}

func (ac *AC_Automation) Build() {
	queue := []*node{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.ch {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				failAcNode := parent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.ch[char]; ok {
						child.fail = parent.fail.ch[char]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					child.fail = ac.root
				}
			}
			queue = append(queue, child)
		}
	}
}

//如果字符串集里至少有一个可以和content的子串匹配，返回true 否则返回false
func (ac *AC_Automation) Find(content string) bool {
	chars := []rune(content)
	iter := ac.root //iterator

	for _, c := range chars {
		_, ok := iter.ch[c]
		for !ok && iter != ac.root {
			//匹配失败，沿着失配指针走
			iter = iter.fail
		}
		if _, ok = iter.ch[c]; ok {

			iter = iter.ch[c]
			if iter.val > 0 {
				return true
			}
		}
	}
	return false
}

func (ac *AC_Automation) Add(Pattern string) {
	ac.root.insert(Pattern,1)
}
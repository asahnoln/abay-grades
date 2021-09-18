package grades

type Group struct {
	items []*Student
}

func NewGroup(ss ...*Student) *Group {
	g := &Group{}
	for _, s := range ss {
		g.items = append(g.items, s)
	}
	return g
}

func (g *Group) Add(s *Student) {
	g.items = append(g.items, s)
}

func (g *Group) Len() int {
	return len(g.items)
}

func (g *Group) Student(i int) *Student {
	return g.items[i]
}

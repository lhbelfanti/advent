package day18

func add(a *Pair, b *Pair) *Pair {
	p := &Pair{
		parent: nil,
		depth:  0,
		left:   a,
		right:  b,
	}

	a.parent = p
	b.parent = p
	incDepth(a)
	incDepth(b)
	reduce(p)
	return p
}

func incDepth(p *Pair) {
	p.depth++
	switch p2 := p.left.(type) {
	case *Pair:
		incDepth(p2)
	case *Num:
		p2.depth++
	}

	switch p2 := p.right.(type) {
	case *Pair:
		incDepth(p2)
	case *Num:
		p2.depth++
	}
}

func iter(p *Pair) []*Num {
	var out []*Num
	switch v := p.left.(type) {
	case *Num:
		out = append(out, v)
	case *Pair:
		out = append(out, iter(v)...)
	}

	switch v := p.right.(type) {
	case *Num:
		out = append(out, v)
	case *Pair:
		out = append(out, iter(v)...)
	}

	return out
}

func reduce(p *Pair) *Pair {
	modified := true
	for modified {
		modified = false

		it := iter(p)
		for i, v := range it {
			if v.depth > 4 {
				if i > 0 {
					it[i-1].value += v.value
				}
				if i < len(it)-2 {
					it[i+2].value += it[i+1].value
				}
				parent := v.parent
				if parent.parent.left == parent {
					parent.parent.left = &Num{
						parent: parent.parent,
						depth:  parent.depth,
						value:  0,
					}
				} else {
					parent.parent.right = &Num{
						parent: parent.parent,
						depth:  parent.depth,
						value:  0,
					}
				}
				modified = true
				break
			}
		}
		if modified {
			continue
		}

		for _, v := range it {
			if v.value >= 10 {
				p := &Pair{
					parent: v.parent,
					depth:  v.depth,
				}
				p.left = &Num{
					parent: p,
					depth:  p.depth + 1,
					value:  v.value / 2,
				}
				p.right = &Num{
					parent: p,
					depth:  p.depth + 1,
					value:  (v.value + 1) / 2,
				}

				if v.parent.left == v {
					v.parent.left = p
				} else {
					v.parent.right = p
				}
				modified = true
				break
			}
		}
	}

	return p
}

func magnitude(p *Pair) int {
	var l, r int
	switch v := p.left.(type) {
	case *Pair:
		l = magnitude(v)
	case *Num:
		l = v.value
	}

	switch v := p.right.(type) {
	case *Pair:
		r = magnitude(v)
	case *Num:
		r = v.value
	}

	return 3*l + 2*r
}

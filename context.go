package svg

type context struct {
	style map[string]string

	previous *context
}

var ctx = &context{style: make(map[string]string)}

func (c *context) pop() {
	if c.previous == nil {
		ctx = c
		return
	}
	ctx = c.previous
}

func (c *context) push() {
	prev := new(context)
	prev.style = ctx.style
	prev.previous = ctx
	ctx = prev
}

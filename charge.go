package socdTool

type ChargeMode struct {
	cur, output Directions
}

func (c ChargeMode) next(d Directions) ChargeMode {
	if !d.left || !d.right {
		return ChargeMode{cur: d, output: d}
	}

	if c.cur.left && c.cur.right {
		return c
	}

	return ChargeMode{
		cur: d,
		output: Directions{left: !c.output.left, right: !c.output.right}}
}

func (c ChargeMode) out() Directions {
	return c.output
}
package socdTool

import "testing"

func TestChargeMode(t *testing.T) {
	var Left = Directions{left: true, right: false}
	var Right = Directions{left: false, right: true}
	var Both = Directions{left: true, right: true}
	var None = Directions{left: false, right: false}

	tests := []struct {
		cur, out, in, want Directions
	}{
		{
			cur: Left,
			out: Left,
			in: Both,
			want: Right,
		},
		{
			cur: Right,
			out: Right,
			in: Both,
			want: Left,
		},
		{
			cur: Both,
			out: Right,
			in: Left,
			want: Left,
		},
		{
			cur: Both,
			out: Left,
			in: Right,
			want: Right,
		},
		{
			cur: Both,
			out: Left,
			in: Both,
			want: Left,
		},
		{
			cur: Both,
			out: Right,
			in: Both,
			want: Right,
		},
		{
			cur: Left,
			out: Left,
			in: None,
			want: None,
		},
		{
			cur: Right,
			out: Right,
			in: None,
			want: None,
		},
	}

	for _, c := range tests {
		mode := ChargeMode{cur: c.cur, output: c.out}
		got := mode.next(c.in)
		if got.out() != c.want {
			t.Errorf("Didn't get what we wanted")
		}
		if got.cur != c.in {
			t.Errorf("Didn't keep state of previous")
		}
	}
}

package socdTool

type Directions struct {
	left, right bool
}

type DirectionsInputStream interface {
  poll() Directions
}

type DirectionsRenderer interface {
  render(Directions) error
}
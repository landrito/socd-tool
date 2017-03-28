package socdTool

type Mode interface {
  next(Directions) Mode
  out() Directions
}
package socdTool

type DirectionsHandler struct {
  m Mode
  in DirectionsInputStream
  renderer DirectionsRenderer
}

func (h *DirectionsHandler) next() (error, *DirectionsHandler) {
  d := h.in.poll()
  m := h.m.next(d)
  error := h.renderer.render(m.out())

  return error, &DirectionsHandler{m: m, in: h.in, renderer: h.renderer}
}
package oebb

func (o *OebbStyleRenderer) Red() {
	o.surface.SetSourceRGB(1.0, 0.0, 0.0)
}

func (o *OebbStyleRenderer) Black() {
	o.surface.SetSourceRGB(0.0, 0.0, 0.0)
}

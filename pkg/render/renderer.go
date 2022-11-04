package render

import (
	"github.com/joushx/traingraph/pkg/model"
	"github.com/ungerik/go-cairo"
)

type Renderer interface {
	GetPageSize() (float64, float64)
	SetData(infrastructure []model.InfrastructureObject, journeys []model.Journey)
	SetSurface(surface *cairo.Surface)
	Render()
}

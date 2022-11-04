package traingraph

import (
	"github.com/joushx/traingraph/pkg/model"
	"github.com/joushx/traingraph/pkg/render"
	"github.com/ungerik/go-cairo"
)

type TrainGraph struct {
	infrastructure []model.InfrastructureObject
	journeys       []model.Journey
	renderer       render.Renderer
}

func NewTrainGraph(infrastructure []model.InfrastructureObject, journeys []model.Journey, renderer render.Renderer) TrainGraph {
	return TrainGraph{
		infrastructure: infrastructure,
		journeys:       journeys,
		renderer:       renderer,
	}
}

func (t TrainGraph) GeneratePDF(filename string) {
	pageWidth, pageHeight := t.renderer.GetPageSize()
	surface := cairo.NewPDFSurface(
		filename,
		pageWidth,
		pageHeight,
		cairo.PDF_VERSION_1_5,
	)

	t.renderer.SetData(t.infrastructure, t.journeys)
	t.renderer.SetSurface(surface)
	t.renderer.Render()

	surface.Finish()
}

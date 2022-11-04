package oebb

import (
	"log"
	"time"

	"github.com/joushx/traingraph/pkg/model"
	"github.com/joushx/traingraph/pkg/util"
	"github.com/ungerik/go-cairo"
)

// DIN A1
const pageWidth = 84.1 * util.PT_PER_CM
const pageHeight = 59.4 * util.PT_PER_CM

type OebbStyleRenderer struct {
	title             string
	id                string
	validityStartDate time.Time
	validityEndDate   time.Time
	infrastructure    []model.InfrastructureObject
	journeys          []model.Journey
	surface           *cairo.Surface
}

func NewOebbStyleRenderer(title string, id string, validityStartDate time.Time, validityEndDate time.Time) OebbStyleRenderer {
	return OebbStyleRenderer{
		title:             title,
		id:                id,
		validityStartDate: validityStartDate,
		validityEndDate:   validityEndDate,
	}
}

func (o *OebbStyleRenderer) GetPageSize() (float64, float64) {
	return pageWidth, pageHeight
}

func (o *OebbStyleRenderer) SetData(infrastructure []model.InfrastructureObject, journeys []model.Journey) {
	o.infrastructure = infrastructure
	o.journeys = journeys
}

func (o *OebbStyleRenderer) SetSurface(surface *cairo.Surface) {
	o.surface = surface
}

func (o *OebbStyleRenderer) Render() {
	if o.surface == nil {
		log.Fatal("Cannot render without surface")
	}

	if o.infrastructure == nil || o.journeys == nil {
		log.Fatal("Cannot render without data")
	}

	o.renderPages()
}

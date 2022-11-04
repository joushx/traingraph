package oebb

import (
	"log"
	"math"

	"github.com/joushx/traingraph/pkg/model"
	"github.com/joushx/traingraph/pkg/util"
	"github.com/ungerik/go-cairo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func (o *OebbStyleRenderer) renderInfrastructureObjects() {
	o.Red()

	for _, object := range o.infrastructure {
		log.Printf("Draw object %s", object.Name)

		switch object.Type {
		case "station":
			// dots on the station line
			o.renderStationLineDot(object)
			// vertical line with interval markers
			o.renderStationTimeLine(object)
		default:
			// short vertical lines at top and bottom
			o.renderObjectTimeLines(object)
		}

		// kilometer location over the station line
		o.renderObjectLocation(object)

		// name and abbreviation of object
		o.renderObjectName(object)
	}
}

func (o *OebbStyleRenderer) renderStationLineDot(object model.InfrastructureObject) {
	position := o.getXPosition(object)
	o.surface.Arc(chartMarginLeft+position, 5.1*util.PT_PER_CM, 2.5, 0, math.Pi*2)
	o.surface.Fill()
}

func (o *OebbStyleRenderer) renderObjectLocation(object model.InfrastructureObject) {
	position := o.getXPosition(object)

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(7)
	o.surface.MoveTo(chartMarginLeft+position+0.1*util.PT_PER_CM, 4.9*util.PT_PER_CM)

	o.surface.Rotate(-0.5 * math.Pi)
	p := message.NewPrinter(language.German)
	o.surface.ShowText(p.Sprintf("%5.1f", object.Location))
	o.surface.Rotate(0.5 * math.Pi)
}

func (o *OebbStyleRenderer) renderObjectTimeLines(object model.InfrastructureObject) {
	position := o.getXPosition(object)
	o.surface.SetLineWidth(0.5)

	o.surface.MoveTo(chartMarginLeft+position, chartMarginTop)
	o.surface.LineTo(chartMarginLeft+position, chartMarginTop+0.2*util.PT_PER_CM)
	o.surface.Stroke()

	o.surface.MoveTo(chartMarginLeft+position, pageHeight-chartMarginBottom)
	o.surface.LineTo(chartMarginLeft+position, pageHeight-chartMarginBottom-0.2*util.PT_PER_CM)
	o.surface.Stroke()
}

func (o *OebbStyleRenderer) renderStationTimeLine(object model.InfrastructureObject) {

	position := o.getXPosition(object)

	o.surface.SetLineWidth(1)
	o.surface.MoveTo(chartMarginLeft+position, chartMarginTop)
	o.surface.LineTo(chartMarginLeft+position, pageHeight-chartMarginBottom)
	o.surface.Stroke()

	o.surface.SetLineWidth(0.5)
	for i := 1; i <= hoursPerPage*60; i++ {
		var offset = 1.5
		if i%5 == 0 {
			offset = 2.5
		}

		positionY := getYPositionFromMinutes(0, i)
		o.surface.MoveTo(chartMarginLeft+position-offset, chartMarginTop+positionY)
		o.surface.LineTo(chartMarginLeft+position+offset, chartMarginTop+positionY)
		o.surface.Stroke()
	}
}

func (o *OebbStyleRenderer) renderObjectName(object model.InfrastructureObject) {
	position := o.getXPosition(object)

	o.Red()

	fontSize := 10.0
	if object.Type != "station" {
		fontSize = 7.0
	}

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(fontSize)

	// name
	o.surface.MoveTo((chartMarginLeft+position)+fontSize/2, 9*util.PT_PER_CM)
	o.surface.Rotate(-0.5 * math.Pi)
	o.surface.ShowText(object.Name)
	o.surface.Rotate(0.5 * math.Pi)

	// DB640 abbreviation top
	o.surface.MoveTo((chartMarginLeft+position)+fontSize/2, 10.2*util.PT_PER_CM)
	o.surface.Rotate(-0.5 * math.Pi)
	o.surface.ShowText(object.Id.Db640)
	o.surface.Rotate(0.5 * math.Pi)

	// DB640 abbreviation bottom
	o.surface.MoveTo((chartMarginLeft+position)+fontSize/2, pageHeight-6*util.PT_PER_CM)
	o.surface.Rotate(-0.5 * math.Pi)
	o.surface.ShowText(object.Id.Db640)
	o.surface.Rotate(0.5 * math.Pi)
}

package oebb

import "github.com/joushx/traingraph/internal/pkg/renderutils"

const chartMarginLeft = 6.1 * renderutils.PT_PER_CM
const chartMarginTop = 10.7 * renderutils.PT_PER_CM
const chartMarginRight = 3 * renderutils.PT_PER_CM
const chartMarginBottom = 7.4 * renderutils.PT_PER_CM
const stationLineMarginTop = 5.1 * renderutils.PT_PER_CM
const timeIntervalSize = 10
const chartWidth = pageWidth - (chartMarginLeft + chartMarginRight)
const chartHeight = pageHeight - (chartMarginTop + chartMarginBottom)

func (o *OebbStyleRenderer) renderLayout() {
	o.renderStationLine()
	o.renderTimeIntervals()
}

// renderStationLine renders the red line on top of the chart
// where the stations are represented with dots
func (o *OebbStyleRenderer) renderStationLine() {
	o.Red()

	o.surface.SetLineWidth(0.75)
	o.surface.MoveTo(chartMarginLeft, stationLineMarginTop)
	o.surface.LineTo(pageWidth-chartMarginRight, stationLineMarginTop)
	o.surface.Stroke()
}

// renderTimeIntervals renders horizontal lines in different thickness
// every 60, 30 and 10 minutes
func (o *OebbStyleRenderer) renderTimeIntervals() {
	o.Red()

	for i := 0; i <= hoursPerPage*60; i += timeIntervalSize {
		switch i % 60 {
		case 0:
			o.surface.SetLineWidth(1.5)
		case 30:
			o.surface.SetLineWidth(0.75)
		default:
			o.surface.SetLineWidth(0.25)
		}

		positionY := getYPositionFromMinutes(0, i) // we use hour = 0 here, because we don't care

		o.surface.MoveTo(chartMarginLeft, chartMarginTop+positionY)
		o.surface.LineTo(pageWidth-chartMarginRight, chartMarginTop+positionY)
		o.surface.Stroke()
	}
}

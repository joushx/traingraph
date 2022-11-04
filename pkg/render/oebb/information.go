package oebb

import (
	"strconv"

	"github.com/joushx/traingraph/pkg/util"
	"github.com/ungerik/go-cairo"
)

const titlePositionX = 12.3 * util.PT_PER_CM
const titlePositionY = 2.5 * util.PT_PER_CM
const idPositionX = 1.4 * util.PT_PER_CM
const idPositionY = titlePositionY
const hoursPositionX = idPositionX
const hoursPositionY = 3 * util.PT_PER_CM

func (o *OebbStyleRenderer) renderPageInformation(startHour int) {
	o.renderTitle()
	o.renderID()
	o.renderTimeRange(startHour)
	o.renderTimeLabels(startHour)
	o.renderValidity()
}

func (o *OebbStyleRenderer) renderTitle() {
	o.Red()

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(14)
	o.surface.MoveTo(titlePositionX, titlePositionY)
	o.surface.ShowText(o.title)
}

func (o *OebbStyleRenderer) renderID() {
	o.Red()

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(14)
	o.surface.MoveTo(idPositionX, idPositionY)
	o.surface.ShowText(o.id)
}

func (o *OebbStyleRenderer) renderTimeRange(startHour int) {
	o.Black()

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(12)
	o.surface.MoveTo(hoursPositionX, hoursPositionY)
	o.surface.ShowText(strconv.Itoa(startHour) + " - " + strconv.Itoa(startHour+hoursPerPage))
}

func (o *OebbStyleRenderer) renderValidity() {
	o.Black()

	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)
	o.surface.SetFontSize(12)
	o.surface.MoveTo(hoursPositionX, 3.5*util.PT_PER_CM)
	o.surface.ShowText("Gültig vom   " + o.validityStartDate.Format("02.01.2006"))

	o.surface.MoveTo(hoursPositionX, 4*util.PT_PER_CM)
	o.surface.ShowText("bis einschl. " + o.validityEndDate.Format("02.01.2006"))
}

func (o *OebbStyleRenderer) renderTimeLabels(startHour int) {
	o.Black()

	o.surface.SetFontSize(16)
	o.surface.SelectFontFace("monospace", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)

	for i := 0; i <= 6; i++ {
		positionY := getYPositionFromMinutes(0, i*60)
		o.surface.MoveTo(2*util.PT_PER_CM, (chartMarginTop + positionY + 0.2*util.PT_PER_CM))
		o.surface.ShowText(strconv.Itoa(i + startHour))

		o.surface.MoveTo(pageWidth-2*util.PT_PER_CM, (chartMarginTop + positionY + 0.2*util.PT_PER_CM))
		o.surface.ShowText(strconv.Itoa(i + startHour))
	}
}

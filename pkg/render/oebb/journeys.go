package oebb

import (
	"log"
	"math"

	"github.com/joushx/traingraph/internal/pkg/renderutils"
	"github.com/joushx/traingraph/pkg/model"
)

func (o *OebbStyleRenderer) renderJourneys(startHour int) {
	for _, journey := range o.journeys {
		log.Printf("Render journey '%s' (%s)", journey.Name, journey.ID)
		o.renderJourney(journey, startHour)
	}
}

func (o *OebbStyleRenderer) renderJourney(journey model.Journey, startHour int) {
	o.Black()
	o.surface.SetLineWidth(0.5)

	// vertical lines that represent the time inside a station or stop
	o.renderStationStopLines(journey, startHour)

	// lines between the objects
	o.renderConnectionLines(journey, startHour)

	// name of journey
	o.renderJourneyName(journey, startHour)
}

func (o *OebbStyleRenderer) renderStationStopLines(journey model.Journey, startHour int) {
	for stopIndex, stop := range journey.Stops {

		infrastructureObject, found := o.getInfrastructureObjectForId(stop.Id)
		if !found {
			continue
		}

		if infrastructureObject.Type == "station" && !o.shouldDrawStationStopLine(stopIndex, stop, journey.Stops) {
			continue
		}

		positionX := o.getXPosition(infrastructureObject)

		arrivalY := getYPositionFromTime(startHour, stop.Arrival)
		departureY := getYPositionFromTime(startHour, stop.Departure)

		if (arrivalY <= 0 || arrivalY >= chartHeight) && (departureY <= 0 || departureY >= chartHeight) {
			// both arrival and departure are outside of this page
			continue
		} else if arrivalY <= 0 {
			// arrival is on previous page: we draw a line starting at zero
			arrivalY = 0
		} else if departureY >= chartHeight {
			// departure is on next page: we draw until end of page
			departureY = chartHeight
		}

		// lines are longer for stops
		if infrastructureObject.Type == "stop" {
			arrivalY = arrivalY - 0.1*renderutils.PT_PER_CM
			departureY = departureY + 0.1*renderutils.PT_PER_CM
		}

		o.surface.MoveTo(chartMarginLeft+positionX, chartMarginTop+arrivalY)
		o.surface.LineTo(chartMarginLeft+positionX, chartMarginTop+departureY)

		o.surface.Stroke()
	}
}

func (o *OebbStyleRenderer) renderConnectionLines(journey model.Journey, startHour int) {
	for i, stop := range journey.Stops {

		// no line to next stop for last entry
		if i == len(journey.Stops)-1 {
			continue
		}

		nextStop := journey.Stops[i+1]
		o.renderConnectionLine(stop, nextStop, startHour)
	}
}

func (o *OebbStyleRenderer) renderConnectionLine(a model.StopTime, b model.StopTime, startHour int) {
	aPositionY := getYPositionFromTime(startHour, a.Departure)
	bPositionY := getYPositionFromTime(startHour, b.Arrival)

	stopA, found := o.getInfrastructureObjectForId(a.Id)
	if !found {
		return
	}

	stopB, found := o.getInfrastructureObjectForId(b.Id)
	if !found {
		return
	}

	aPositionX := o.getXPosition(stopA)
	bPositionX := o.getXPosition(stopB)

	if (aPositionY <= 0 && bPositionY <= 0) || (aPositionY >= chartHeight && bPositionY >= chartHeight) {
		// both ends are on another page
		return
	} else if aPositionY <= 0 {
		ratioToDraw := 1 - (bPositionY / float64(bPositionY-aPositionY))
		aPositionX = aPositionX + ratioToDraw*(bPositionX-aPositionX)
		aPositionY = 0
	} else if bPositionY >= chartHeight {
		ratioToDraw := math.Abs((chartHeight - aPositionY) / float64((chartHeight-bPositionY)-(chartHeight-aPositionY)))
		bPositionX = aPositionX + ratioToDraw*(bPositionX-aPositionX)
		bPositionY = chartHeight
	}

	o.surface.MoveTo(chartMarginLeft+aPositionX, chartMarginTop+aPositionY)
	o.surface.LineTo(chartMarginLeft+bPositionX, chartMarginTop+bPositionY)
	o.surface.Stroke()
}

// shouldDrawStationStopLine decides whether we want to draw a vertical stop line at this object
// This is not desired if it is the first or last object to draw, the journey does stop for less than
// a minute or the first and last stop of the journey (sometimes we still get data for it)
func (o *OebbStyleRenderer) shouldDrawStationStopLine(index int, stopTime model.StopTime, stopTimes []model.StopTime) bool {

	// no line at first station (e.g. when coming from outside the chart)
	if index == 0 {
		return false
	}

	// no line at last station (e.g. when leaving the chart)
	if index == len(stopTimes)-1 {
		return false
	}

	// no line when time is the same
	if stopTime.Arrival == stopTime.Departure {
		return false
	}

	// no line not both of arrival and departure are defined
	if stopTime.Arrival == "" || stopTime.Departure == "" {
		return false
	}

	// not if the previous stop is not on the chart
	previousStop := stopTimes[index-1]
	_, previousStopFound := o.getInfrastructureObjectForId(previousStop.Id)
	if !previousStopFound {
		return false
	}

	// not if the next stop is not in this chart any more
	// e.g. train leaves route to another route
	nextStop := stopTimes[index+1]
	_, nextStopFound := o.getInfrastructureObjectForId(nextStop.Id)
	if !nextStopFound {
		return false
	}

	return true
}

// getInfrastructureObjectForId searches for an infrastructure object
// that matches the given ID (e.g. from a StopTime of a journey)
func (o *OebbStyleRenderer) getInfrastructureObjectForId(id model.ObjectID) (model.InfrastructureObject, bool) {
	for _, object := range o.infrastructure {
		if renderutils.IsSameObject(object.Id, id) {
			return object, true
		}
	}

	return model.InfrastructureObject{}, false
}

func (o *OebbStyleRenderer) renderJourneyName(journey model.Journey, startHours int) {

	for i, stop := range journey.Stops {
		if i+1 > len(journey.Stops)-1 {
			return
		}

		stopObject, found := o.getInfrastructureObjectForId(stop.Id)
		nextStopObject, nextStopFound := o.getInfrastructureObjectForId(journey.Stops[i+1].Id)
		if found && nextStopFound {
			positionX := o.getXPosition(stopObject)
			positionY := getYPositionFromTime(startHours, stop.Departure)
			nextPositionX := o.getXPosition(nextStopObject)
			nextPositionY := getYPositionFromTime(startHours, journey.Stops[i+1].Arrival)

			if positionY <= 0 || positionY >= chartHeight {
				return
			}

			sizeX := nextPositionX - positionX
			sizeY := nextPositionY - positionY
			rotation := math.Atan(1 / (sizeX / sizeY))

			o.surface.SetFontSize(8)
			o.surface.MoveTo(chartMarginLeft+positionX+0.5*renderutils.PT_PER_CM, chartMarginTop+positionY-0.1*renderutils.PT_PER_CM)
			o.surface.Rotate(rotation)
			o.surface.ShowText(journey.Name)
			o.surface.Rotate(-rotation)

			break
		}
	}
}

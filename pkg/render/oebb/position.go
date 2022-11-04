package oebb

import (
	"github.com/joushx/traingraph/internal/pkg/idutil"
	"github.com/joushx/traingraph/pkg/model"
	"github.com/joushx/traingraph/pkg/util"
)

// getYPosition calculates the position inside the graph. One has to provide
// the starting hour of the page we are on and the time of the day in minutes
func getYPositionFromMinutes(pageStartingHour int, minutes int) float64 {
	minutesOnPage := minutes - pageStartingHour*60
	positionOnChartRatio := float64(minutesOnPage) / (float64(hoursPerPage) * 60)
	return positionOnChartRatio * chartHeight
}

func getYPositionFromTime(startHours int, time string) float64 {
	minutes := util.ParseTime(time)
	return getYPositionFromMinutes(startHours, minutes)
}

func (o *OebbStyleRenderer) getXPosition(object model.InfrastructureObject) float64 {
	var pathLength float32 = 0.0
	for _, object := range o.infrastructure {
		pathLength += object.Distance
	}

	var distanceFromStart float32 = 0
	for _, currentObject := range o.infrastructure {
		distanceFromStart += currentObject.Distance

		if idutil.IsSameObject(currentObject.Id, object.Id) {
			break
		}
	}

	return float64(distanceFromStart/pathLength) * chartWidth
}

package oebb

import "log"

const hoursPerPage = 6
const hoursPerDay = 24

func (o *OebbStyleRenderer) renderPages() {
	for hour := 0; hour < hoursPerDay; hour += hoursPerPage {
		log.Printf("Draw page for %vh", hour)
		o.drawPage(hour)
		o.surface.ShowPage()
	}
}

func (o *OebbStyleRenderer) drawPage(startHour int) {
	log.Printf("Render layout")
	o.renderLayout()

	log.Printf("Render page information")
	o.renderPageInformation(startHour)

	log.Printf("Render stations")
	o.renderInfrastructureObjects()

	log.Printf("Render journeys")
	o.renderJourneys(startHour)
}

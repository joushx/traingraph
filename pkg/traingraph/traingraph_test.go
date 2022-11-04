package traingraph_test

import (
	"testing"
	"time"

	"github.com/joushx/traingraph/pkg/model"
	"github.com/joushx/traingraph/pkg/render/oebb"
	"github.com/joushx/traingraph/pkg/traingraph"
	"github.com/joushx/traingraph/pkg/util"
)

func TestOebb(t *testing.T) {
	infrastructure := []model.InfrastructureObject{
		{
			Name:     "A",
			Distance: 0,
			Location: 123.45,
			Type:     "station",
			Id: model.ObjectID{
				Db640: "Aa",
				Ifopt: util.NewIfOpt("at:44:1234:0:1"),
			},
		},
		{
			Name:     "A Ort",
			Distance: 0.5,
			Location: 123.7,
			Type:     "stop",
			Id: model.ObjectID{
				Db640: "Aa H1",
				Ifopt: util.NewIfOpt("at:44:1235"),
			},
		},
		{
			Name:     "B",
			Distance: 1,
			Location: 124.45,
			Type:     "station",
			Id: model.ObjectID{
				Db640: "Bb",
				Ifopt: util.NewIfOpt("at:44:5678:0:2"),
			},
		},
		{
			Name:     "C",
			Distance: 1,
			Location: 125.0,
			Type:     "station",
			Id: model.ObjectID{
				Db640: "Cc",
				Ifopt: util.NewIfOpt("at:44:9876:0:2"),
			},
		},
	}

	timetable := []model.Journey{
		{
			ID:   "1234",
			Name: "REX 1234",
			Stops: []model.StopTime{
				{
					Id: model.ObjectID{
						Ifopt: util.NewIfOpt("at:44:1234:0:1"),
					},
					Departure: "09:00",
				},
				{
					Id: model.ObjectID{
						Ifopt: util.NewIfOpt("at:44:1235"),
					},
					Arrival:   "09:05",
					Departure: "09:06",
				},
				{
					Id: model.ObjectID{
						Ifopt: util.NewIfOpt("at:44:5678:0:2"),
					},
					Arrival:   "09:15",
					Departure: "09:17",
				},
				{
					Id: model.ObjectID{
						Ifopt: util.NewIfOpt("at:44:9876:0:2"),
					},
					Arrival: "09:30",
				},
			},
		},
	}

	renderer := oebb.NewOebbStyleRenderer(
		"A - B",
		"Blatt 1234",
		time.Now(),
		time.Now().AddDate(0, 0, 1),
	)

	traingraph := traingraph.NewTrainGraph(
		infrastructure,
		timetable,
		&renderer,
	)
	traingraph.GeneratePDF("out.pdf")
}

package function_chart_calculator

import (
	"github.com/wcharczuk/go-chart"
	"net/http"
)

//func main(){
//	http.HandleFunc("/", drawChart)
//	http.ListenAndServe(":8080", nil)
//}

func drawChart(res http.ResponseWriter, req *http.Request) {

	/*
	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
	*/

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "The N Axis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "The Y Axis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: []float64{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 100, 200, 300},
				YValues: []float64{0, 2.772588722239781, 9.887510598012986, 22.18070977791825, 40.23594781085251, 230.25850929940458, 1198.2929094215963, 3061.07764349594, 5902.207126582298, 9780.057513570366, 46051.701859880915, 211932.69466192144, 513340.42271905806},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

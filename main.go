package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/yosssi/ace"

	datatypes "app/datatypes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"GetOrderNumber": func(n int) int {
			incremented := n + 1
			return incremented
		},
		"UppetCaseTracker": func(s string) string {
			result := strings.ToUpper(s)
			return result
		},
		"DeleteMinus": func(n float32) float32 {
			plused := n * -1
			return plused
		},
	}

	tpl, err := ace.Load("templates/privatStatistic/base", "templates/privatStatistic/inner", &ace.Options{
		FuncMap: funcMap,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := datatypes.User{
		Name: "Aaron Ross",
	}

	timerange := datatypes.TimeRange{
		From: "February 19th, 2018",
		To:   "February 26th, 2018",
	}

	domen := datatypes.Domen{
		Title:   "facebook.com",
		Address: "http://facebook.com",
	}

	statistic := datatypes.Statistic{
		TotalTime:           99,
		TotalTimeComparison: -8,
		Useful:              82,
		UsefulComparison:    9,
		MostUsed:            []string{"Adobe Photoshop", "Illustrator", "Sketch", "Pollaris Office", "Photoshop"},
	}

	teammate1 := datatypes.Teammate{
		Name:   "Tom Delonge",
		Avatar: "public/images/gentleman.svg",
		Time:   6.4,
	}

	teammate2 := datatypes.Teammate{
		Name:   "Stive Churchill",
		Avatar: "public/images/detective.svg",
		Time:   1,
	}

	teammate3 := datatypes.Teammate{
		Name:   "Jack White",
		Avatar: "public/images/astronaut.svg",
		Time:   5.5,
	}

	project2 := datatypes.Project{
		Name:      "CodeQuality Redesign",
		Tracker:   "Jira",
		Time:      29,
		Teammates: []datatypes.Teammate{teammate1, teammate2, teammate3},
	}

	project1 := datatypes.Project{
		Name:      "Glimmer (StartingFive)",
		Tracker:   "Redmine",
		Time:      23,
		Teammates: []datatypes.Teammate{teammate1, teammate2, teammate3},
	}

	info := datatypes.TemplateInfo{
		Title:     "Your Weekly Statistics",
		User:      user,
		TimeRange: timerange,
		Domen:     domen,
		Statistic: statistic,
		Projects:  []datatypes.Project{project1, project2},
	}

	if err := tpl.Execute(w, info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlerFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	fmt.Println("open  http://localhost:8080/")
	http.HandleFunc("/", handler)
	http.HandleFunc("/public/", handlerFiles)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/yosssi/ace"

	"golang/GoTemplates/datatypes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"GetOrderNumber": func(n int) int {
			return 4
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
		Title:   "tgms.qarea.org",
		Address: "http://tgms.qarea.org",
	}

	statistic := datatypes.Statistic{
		TotalTime:           999,
		TotalTimeComparison: 2,
		Useful:              82,
		UsefulComparison:    21,
		MostUsed:            []string{"Adobe Photoshop", "Illustrator", "Sketch", "Pollaris Office", "Photoshop"},
	}

	teammate1 := datatypes.Teammate{
		Name: "Tom Delonge",
		Time: 6.4,
	}

	teammate2 := datatypes.Teammate{
		Name: "Stive Churchill",
		Time: 6.0,
	}

	teammate3 := datatypes.Teammate{
		Name: "Jack White",
		Time: 5.5,
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
		Projects:  []datatypes.Project{project1},
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

package datatypes

type User struct {
	Name string
}

type TimeRange struct {
	From string
	To   string
}

type Domen struct {
	Title   string
	Address string
}

type Statistic struct {
	TotalTime           float32
	TotalTimeComparison float32
	Useful              float32
	UsefulComparison    float32
	MostUsed            []string
}

type Teammate struct {
	Name string
	Time float32
}

type Project struct {
	Name      string
	Tracker   string
	Time      float32
	Teammates []Teammate
}

type TemplateInfo struct {
	Title string
	User
	TimeRange
	Domen
	Statistic
	Projects []Project
}

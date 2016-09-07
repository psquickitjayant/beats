package console

type config struct {
	Url string `config:"Url"`
}

var (
	defaultConfig = config{
        Url: "http://logs-01.loggly.com/inputs/token/45866f2f-57c3-459c-bc3d-a19214c34edf/tag/beats",
	}
)

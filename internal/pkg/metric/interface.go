package metric

import "time"

type CLI struct {
	Name       string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewCLI(name string) *CLI {
	return &CLI{Name: name}
}

func (c *CLI) StartedApp() {
	c.StartedAt = time.Now()
}

func (c *CLI) FinishedApp() {
	c.FinishedAt = time.Now()
	c.Duration = time.Since(c.StartedAt).Seconds()
}

type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewHTTP(handler string, method string) *HTTP {
	return &HTTP{Handler: handler, Method: method}
}

func (h *HTTP) StartedHttp() {
	h.StartedAt = time.Now()
}

func (h *HTTP) FinishedHttp() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}

type Service interface {
	SaveCLI(c *CLI) error
	SaveHTTP(h *HTTP)
}

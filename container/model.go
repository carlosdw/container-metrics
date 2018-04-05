package container

type Container struct {
	Id     string
	Name   string
	Labels map[string]string

	ComputedLabels map[string]string
}

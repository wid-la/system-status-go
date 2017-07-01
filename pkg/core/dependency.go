package core

import "fmt"

// Dependency ...
type Dependency struct {
	Interval int // In Milliseconds
	Version  string
	Port     string
	Timeout  int // Service Timeout in Seconds
	Token    string
}

const (
	defaultTag      = "Develop"
	defaultInterval = 300
	defaultPort     = ":8080"
	defaultTimeout  = 1800
	defaultToken    = "xxx"
)

// NewDependency ...
func NewDependency() *Dependency {
	coreDeps := &Dependency{
		Version:  envString("VERSION_TAG", defaultTag),
		Interval: envInt("INTERVAL", defaultInterval),
		Port:     envString("HTTP_PORT", defaultPort),
		Timeout:  envInt("TIMEOUT", defaultTimeout),
		Token:    envString("TOKEN", defaultToken),
	}

	return coreDeps
}

// FieldsReport ...
func (dep *Dependency) FieldsReport() string {
	report := fmt.Sprintf("Version Tag '%v' \n", dep.Version)
	report += fmt.Sprintf("Process Interval %v \n", dep.Interval)
	report += fmt.Sprintf("Service Timeout %v \n", dep.Timeout)
	report += fmt.Sprintf("Auth Token %v \n", dep.Token)

	report += fmt.Sprintf("Listening to port %v \n", dep.Port)

	return report
}

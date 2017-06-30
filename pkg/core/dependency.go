package core

import "fmt"

// Dependency ...
type Dependency struct {
	Interval int
	Version  string
	Port     string
	// DB       *Database
}

const (
	defaultTag        = "Develop"
	defaultInterval   = 300
	defaultPort       = ":8080"
	defaultDBEndpoint = ""
	defaultDBUser     = ""
	defaultDBPassword = ""
	defaultDBName     = ""
)

// NewDependency ...
func NewDependency() *Dependency {
	coreDeps := &Dependency{
	// Version:  envString("VERSION_TAG", defaultTag),
	// Interval: envInt("INTERVAL", defaultInterval),
	// Port:     envString("HTTP_PORT", defaultPort),
	// DB: NewDatabase(envString("DB_ENDPOINT", defaultDBEndpoint),
	// 	envString("DB_ENDPOINT", defaultDBEndpoint),
	// 	envString("DB_ENDPOINT", defaultDBEndpoint),
	// 	envString("DB_ENDPOINT", defaultDBEndpoint)),
	}

	return coreDeps
}

// FieldsReport ...
func (dep *Dependency) FieldsReport() string {
	report := fmt.Sprintf("Version Tag '%v' \n", dep.Version)
	// report += fmt.Sprintf("Process Interval %v \n", dep.Interval)

	// report += fmt.Sprintf("Db Endpoint %v \n", dep.DB.Endpoint)
	// report += fmt.Sprintf("Db User %v \n", dep.DB.User)
	// report += fmt.Sprintf("Db Pass %v \n", dep.DB.Pass)
	// report += fmt.Sprintf("Db Name %v \n", dep.DB.Name)

	// report += fmt.Sprintf("Listening to port %v \n", dep.Port)

	return report
}

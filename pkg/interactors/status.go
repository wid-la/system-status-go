package interactors

// CheckStatuses ...
func (inter Inter) CheckStatuses() {
	for i, srv := range inter.Services {
		srv.CheckStatus()
		inter.Services[i] = srv
	}
}

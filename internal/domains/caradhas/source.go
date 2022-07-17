package caradhas

type infoDTO struct {
	Version string
}

//New Caradhas
func New(version string) {

	dto := infoDTO{}

	dto.Version = version
}

//Info retrieves intormation about the caradhas version
func Info() string {
	return "v2.0.0"
}

package version

var (
	deployVersion string
	deployMode    string
)

func GetDeployVersion() string {
	return deployVersion
}

func GetDeployMode() string {
	return deployMode
}

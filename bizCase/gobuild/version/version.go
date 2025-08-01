package version

var (
	DeployVersion string
	DeployMode    string
)

func GetDeployVersion() string {
	return DeployVersion
}

func GetDeployMode() string {
	return DeployMode
}

package constant


type Environment string

const (
	EnvironmentLocal Environment = "LOCAL"
	EnvironmentTest  Environment = "TEST"
	EnvironmentDev   Environment = "DEV"
	EnvironmentProd  Environment = "PROD"
)

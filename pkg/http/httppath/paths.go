package httppath

const (
	AnalyticBasePath = "/v2"
	ListenerBasePath = "/v1"

	AuthGroup        = "/auth"
	LoginPath        = AuthGroup + "/login"
	AuthorizePath    = AuthGroup + "/authorize"
	TokenPath        = AuthGroup + "/token"
	ActivateUserPath = AuthGroup + "/activate"

	SystemGroup = "/system"
	VersionPath = SystemGroup + "/version"
)

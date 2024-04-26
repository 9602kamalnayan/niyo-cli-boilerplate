package constants

// RouterBasePrefix BasePrefix added before all router paths
type RouterBasePrefix string

// RVersion ConfigVersion management for routes
type RVersion string

const (
	RV1 RVersion = "v1"
)

// RoutesGroup route groups
type RoutesGroup string

// RoutesPath route sub-paths
type RoutesPath string

const (
	RGDocs            RoutesGroup = "docs"
	RGInternalGroup   RoutesGroup = "testRoutes"
	RGCheckRouterPath RoutesPath  = "routerCheck"
)

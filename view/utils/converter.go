package utils

import "github.com/PavlushaSource/Radar/view/config"

var ConvertStringToDistanceType = map[string]config.DistanceType{
	"Euclidean":   config.Euclidean,
	"Manhattan":   config.Manhattan,
	"Curvilinear": config.Curvilinear,
}

var ConvertDistanceTypeToString = map[config.DistanceType]string{
	config.Euclidean:   "Euclidean",
	config.Manhattan:   "Manhattan",
	config.Curvilinear: "Curvilinear",
}

var ConvertStringToGeometryType = map[string]config.GeometryType{
	"Simple": config.Simple,
	"Vector": config.Vector,
}

var ConvertGeometryTypeToString = map[config.GeometryType]string{
	config.Simple: "Simple",
	config.Vector: "Vector",
}

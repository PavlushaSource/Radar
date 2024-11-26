package utils

import (
	"github.com/PavlushaSource/Radar/view/api"
)

var ConvertStringToDistanceType = map[string]api.DistanceType{
	"Euclidean":   api.Euclidean,
	"Manhattan":   api.Manhattan,
	"Curvilinear": api.Curvilinear,
}

var ConvertDistanceTypeToString = map[api.DistanceType]string{
	api.Euclidean:   "Euclidean",
	api.Manhattan:   "Manhattan",
	api.Curvilinear: "Curvilinear",
}

var ConvertStringToGeometryType = map[string]api.GeometryType{
	"Simple": api.Simple,
	"Vector": api.Vector,
}

var ConvertGeometryTypeToString = map[api.GeometryType]string{
	api.Simple: "Simple",
	api.Vector: "Vector",
}

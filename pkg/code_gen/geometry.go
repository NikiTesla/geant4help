package codegen

type Geometry struct {
	Shape string `json:"shape"`

	SizeX       float64 `json:"size_x"`
	SizeY       float64 `json:"size_y"`
	SizeZ       float64 `json:"size_z"`
	InnerRadius float64 `json:"inner_r"`
	OuterRadius float64 `json:"outer_r"`

	Material string `json:"material"`
}

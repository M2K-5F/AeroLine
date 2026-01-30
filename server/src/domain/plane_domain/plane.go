package plane_domain

type Plane struct {
	id   PlaneID
	name Name
}

func (ths Plane) ID() PlaneID {
	return ths.id
}

func (ths Plane) Name() Name {
	return ths.name
}

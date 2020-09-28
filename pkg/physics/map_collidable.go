package physics

type Axis int

const (
	XAxis Axis = iota
	YAxis
)

func UpdateCollidableY(
	accelerator Accelerator,
	kinematics *Kinematics2D,
) {
	UpdateCollidable(
		accelerator,
		kinematics,
		YAxis,
	)
}

func UpdateCollidable(
	accelerator Accelerator,
	kinematics *Kinematics2D,
	axis Axis,
) {
	var k *Kinematics
	switch axis {
	case XAxis:
		k = &kinematics.X
	case YAxis:
		k = &kinematics.Y
	}
	accelerator.UpdateVelocity(k)
	delta := k.Delta()
	k.Position += delta
}

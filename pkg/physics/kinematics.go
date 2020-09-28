package physics

import "kokos_quest/pkg/units"

type Kinematics struct {
	Position float64
	Velocity float64
}

func (k Kinematics) Delta() float64 {
	return k.Velocity
}

type Kinematics2D struct {
	X, Y Kinematics
}

func MakeKinematics2D(pos units.Vec, vel units.VecFloat) Kinematics2D {
	return Kinematics2D {
		X: Kinematics{
			Position: float64(pos.X),
			Velocity: vel.X,
		},
		Y: Kinematics{
			Position: float64(pos.Y),
			Velocity: vel.Y,
		},
	}
}

func (k Kinematics2D) Position() units.Vec {
	return units.Vec{
		X: int(k.X.Position),
		Y: int(k.Y.Position),
	}
}

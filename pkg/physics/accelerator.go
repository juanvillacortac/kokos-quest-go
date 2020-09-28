package physics

import (
	"fmt"
	"math"
)

const (
	GravityAcceleration float64 = 1.00078125
	TerminalSpeed       float64 = 1.2998046875
)

var (
	Zero    = &ZeroAccelerator{}
	Gravity = &ConstantAccelerator{
		Acceleration: GravityAcceleration,
		MaxVelocity:  TerminalSpeed,
	}
)

type Accelerator interface {
	UpdateVelocity(kinematics *Kinematics)
}

type ZeroAccelerator struct{}

func (a *ZeroAccelerator) UpdateVelocity(kinematics *Kinematics) {}

type ResetAccelerator struct{}

func (a *ResetAccelerator) UpdateVelocity(kinematics *Kinematics) {
	kinematics.Velocity = 0
}

type ConstantAccelerator struct {
	Acceleration float64
	MaxVelocity  float64
}

func (a *ConstantAccelerator) UpdateVelocity(kinematics *Kinematics) {
	velocity := kinematics.Velocity + a.Acceleration*1.6
	if a.Acceleration < 0 {
		kinematics.Velocity = math.Max(velocity, a.MaxVelocity)
	} else {
		kinematics.Velocity = math.Min(velocity, a.MaxVelocity)
	}
	fmt.Println(kinematics.Velocity)
}

type FrictionAccelerator struct {
	Friction float64
}

func (a *FrictionAccelerator) UpdateVelocity(kinematics *Kinematics) {
	if kinematics.Velocity > 0 {
		kinematics.Velocity = math.Max(0, kinematics.Velocity-a.Friction)
	} else {
		kinematics.Velocity = math.Min(0, kinematics.Velocity+a.Friction)
	}
}

type BidirectionalAccelerator struct {
	Positive, Negative *ConstantAccelerator
}

func (a *BidirectionalAccelerator) Init(acceleration, maxVelocity float64) {
	a.Positive = &ConstantAccelerator{
		Acceleration: acceleration,
		MaxVelocity:  maxVelocity,
	}
	a.Negative = &ConstantAccelerator{
		Acceleration: -acceleration,
		MaxVelocity:  -maxVelocity,
	}
}

package player

import "kokos_quest/pkg/physics"

func (p *Player) updateY() {
	var accelerator physics.Accelerator
	if /*jump_active &&*/ p.kinematics.Y.Velocity < 0 {
		// accelerator = kJumpGravityAccelerator
		accelerator = physics.Gravity
	} else {
		accelerator = physics.Gravity
	}
	physics.UpdateCollidableY(accelerator, &p.kinematics)
}

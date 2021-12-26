package razor

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
)

func (c *char) Attack(p map[string]int) (int, int) {

	f, a := c.ActionFrames(core.ActionAttack, p)
	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
		AttackTag:  core.AttackTagNormal,
		ICDTag:     core.ICDTagNormalAttack,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeBlunt,
		Element:    core.Physical,
		Durability: 25,
		Mult:       attack[c.NormalCounter][c.TalentLvlAttack()],
	}
	r := 0.3
	if c.Core.Status.Duration("noelleq") > 0 {
		r = 2
	}

	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(r, false, core.TargettableEnemy), 0, f-1)

	c.AdvanceNormalIndex()

	return f, a
}

func (c *char) Skill(p map[string]int) (int, int) {
	f, a := c.ActionFrames(core.ActionSkill, p)

	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Claw And Thunder",
		AttackTag:  core.AttackTagElementalArt,
		ICDTag:     core.ICDTagElementalArt,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeBlunt,
		Element:    core.Electro,
		Durability: 50,
		Mult:       skillPress[c.TalentLvlSkill()],
	}
	snap := c.Snapshot(&ai)

	c.SetCD(core.ActionSkill, 24*60)
	return f, a
}

func (c *char) Burst(p map[string]int) (int, int) {
	f, a := c.ActionFrames(core.ActionSkill, p)

	c.Core.Status.AddStatus("noelleq", 900+f)

	ai := core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Sweeping Time (Burst)",
		AttackTag:  core.AttackTagElementalBurst,
		ICDTag:     core.ICDTagElementalBurst,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeBlunt,
		Element:    core.Electro,
		Durability: 25,
		Mult:       burst[c.TalentLvlBurst()],
	}
	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(6.5, false, core.TargettableEnemy), f-10, f-10)

	ai = core.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Sweeping Time (Skill)",
		AttackTag:  core.AttackTagElementalBurst,
		ICDTag:     core.ICDTagElementalBurst,
		ICDGroup:   core.ICDGroupDefault,
		StrikeType: core.StrikeTypeBlunt,
		Element:    core.Electro,
		Durability: 25,
		Mult:       burst[c.TalentLvlBurst()],
	}
	c.Core.Combat.QueueAttack(ai, core.NewDefCircHit(4.5, false, core.TargettableEnemy), f-10, f-10)

	c.SetCD(core.ActionBurst, 900)
	c.ConsumeEnergy(0)
	return f, a
}

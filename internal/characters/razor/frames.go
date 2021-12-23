package razor

import "github.com/genshinsim/gcsim/pkg/core"

func (c *char) ActionFrames(a core.ActionType, p map[string]int) (int, int) {
	switch a {
	case core.ActionAttack:
		f := 0
		switch c.NormalCounter {
		//TODO: need to add atkspd mod
		case 0:
			f = 25 //frames from keqing lib
		case 1:
			f = 71 - 25
		case 2:
			f = 109 - 71
		case 3:
			f = 192 - 109
		}
		f = int(float64(f) / (1 + c.Stats[core.AtkSpd]))
		return f, f
	case core.ActionSkill:
		if p["hold"] == 0 {
			return 74, 74
		}
		return 92, 92
	case core.ActionBurst:
		return 62, 62 //ok
	default:
		c.Core.Log.Warnf("%v: unknown action (%v), frames invalid", c.Base.Key.String(), a)
		return 0, 0
	}
}

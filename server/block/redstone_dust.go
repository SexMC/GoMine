package block

type RedstoneDust struct {
	empty
	transparent
	Power int
}

// EncodeItem ...
func (RedstoneDust) EncodeItem() (name string, meta int16) {
	return "minecraft:redstone", 0
}

// EncodeBlock ...
func (r RedstoneDust) EncodeBlock() (string, map[string]any) {
	return "minecraft:redstone_wire", map[string]any{
		"redstone_signal": int32(r.Power),
	}
}

// BreakInfo ...
func (r RedstoneDust) BreakInfo() BreakInfo {
	return newBreakInfo(0.5, neverHarvestable, nil, oneOf(r))
}

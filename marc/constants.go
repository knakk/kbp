package marc

type ControlTag int

const (
	tagIllegalControlTag ControlTag = iota
	Tag001
	Tag002
	Tag003
	Tag004
	Tag005
	Tag006
	Tag007
	Tag008
	Tag009
)

type DataTag int

const (
	tagIllegalDataTag DataTag = iota
	Tag010                    = iota + 10
	Tag020
	Tag041
	Tag111
	Tag245
	Tag260
	Tag300
	Tag440
	Tag508
	Tag546
	Tag700
	Tag911
)

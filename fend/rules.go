package fend

import "github.com/foomo/fender/rule"

func Int(v int, rules ...rule.IntRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func Int8(v int8, rules ...rule.Int8Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func Int32(v int32, rules ...rule.Int32Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func Int64(v int64, rules ...rule.Int64Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func UInt(v uint, rules ...rule.UIntRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func UInt8(v uint8, rules ...rule.UInt8Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func UInt32(v uint32, rules ...rule.UInt32Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func UInt64(v uint64, rules ...rule.UInt64Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func Float32(v float32, rules ...rule.Float32Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}
func Float64(v float64, rules ...rule.Float64Rule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}

func Bool(v bool, rules ...rule.BoolRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}

func String(v string, rules ...rule.StringRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}

func Interface(v interface{}, rules ...rule.InterfaceRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}

func Validator(v rule.Validator, rules ...rule.ValidatorRule) Fend {
	return func() []rule.Rule {
		ret := make([]rule.Rule, 0, len(rules))
		for _, rule := range rules {
			ret = append(ret, rule(v))
		}
		return ret
	}
}

func Var(v interface{}, tag string) Fend {
	return func() []rule.Rule {
		return []rule.Rule{
			rule.Var(tag)(v),
		}
	}
}

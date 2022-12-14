package waops

import "go/types"

// 对于指针获取指针指向的类型, 或者返回当前类型.
func Deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}

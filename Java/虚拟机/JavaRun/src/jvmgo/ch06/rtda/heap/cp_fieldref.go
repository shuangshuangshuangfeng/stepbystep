package heap

import "jvmgo/ch06/classfile"
type FieldRef struct { // 字段符号引用
	MemberRef
	field *Field
}
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}


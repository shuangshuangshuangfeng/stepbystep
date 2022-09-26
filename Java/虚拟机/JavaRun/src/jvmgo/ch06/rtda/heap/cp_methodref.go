package heap

import "jvmgo/ch06/classfile"
type MethodRef struct { // 方法符号的引用
	MemberRef
	method *Method
}
func newMethodRef(cp *ConstantPool,
	refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

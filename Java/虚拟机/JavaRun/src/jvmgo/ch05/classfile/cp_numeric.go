package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = int32(bytes)
}


type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}





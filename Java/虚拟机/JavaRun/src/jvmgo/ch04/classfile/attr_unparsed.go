package classfile


type UnparseAttribute struct {
	name string
	length uint32
	info []byte
}

func (self *UnparseAttribute) readInfo(reader * ClassReader){
	self.info = reader.readBytes(self.length)
}




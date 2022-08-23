package main

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/instructions"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 编写简单的解释器

func catchErr(frame *rtda.Frame){
	if r := recover(); r != nil{
		fmt.Printf("LocalVars : %v \n", frame.LocalVars())
		fmt.Printf("OperandStack: %v \n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// excute
		fmt.Printf("pc :%2d inst:%T %v \n", pc, inst, inst)
		inst.Execute(frame)

	}
}

func interpret(methodInfo *classfile.MemberInfo){
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
	// ...
}

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Op string

const (
	Acc Op = "acc"
	Jmp Op = "jmp"
	Nop Op = "nop"
)

type Instruction struct {
	Op    Op
	Param int
}

type HH struct {
	Mem []Instruction
	IP  int
	Acc int
}

func (v HH) Copy() HH {
	mem := make([]Instruction, len(v.Mem))
	copy(mem, v.Mem)
	return HH{mem, v.IP, v.Acc}
}

func (v *HH) Exec() bool {
	ip := (*v).IP
	mem := (*v).Mem
	if ip >= len(mem) {
		return false
	}
	i := mem[ip]
	switch i.Op {
	case Acc:
		(*v).Acc += i.Param
		(*v).IP++
	case Jmp:
		(*v).IP += i.Param
	case Nop:
		(*v).IP++
	}
	return true
}

func parse(code string) HH {
	lines := strings.Split(strings.TrimSpace(code), "\n")
	inst := make([]Instruction, len(lines))
	for i, line := range lines {
		parsedLine := strings.Split(line, " ")
		param, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			log.Fatal(err.Error())
		}
		op := Op(parsedLine[0])
		inst[i] = Instruction{op, param}
	}
	return HH{inst, 0, 0}
}

func loopOne(vm *HH) {
	visitedIP := make(map[int]struct{})
	ok := true
	for ok {
		if _, ok := visitedIP[vm.IP]; !ok {
			visitedIP[vm.IP] = struct{}{}
		} else {
			break
		}

		ok = vm.Exec()
	}
}

func loopTwo(v HH) HH {
	mutatedIP := make(map[int]struct{})
	for {
		visitedIP := make(map[int]struct{})
		running := true
		mutated := false

		vm := v.Copy()
		for running {
			// trying to mutate Op
			if _, ok := mutatedIP[vm.IP]; !mutated && !ok {
				op := vm.Mem[vm.IP].Op
				if op == Jmp {
					vm.Mem[vm.IP].Op = Nop
					mutated = true
				} else if op == Nop {
					vm.Mem[vm.IP].Op = Jmp
					mutated = true
				}
				if mutated {
					mutatedIP[vm.IP] = struct{}{}
				}
			}

			if _, ok := visitedIP[vm.IP]; !ok {
				visitedIP[vm.IP] = struct{}{}
			} else {
				break
			}

			running = vm.Exec()
		}

		// is succeed
		if !running {
			return vm
		}
	}
}

func Fn() {
	code, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err.Error())
	}
	vm := parse(string(code))

	firstVM := vm.Copy()
	loopOne(&firstVM)
	secVM := loopTwo(vm.Copy())

	log.Printf("%v", firstVM.Acc)
	log.Printf("%v", secVM.Acc)
}

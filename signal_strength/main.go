package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const addxPattern = `^addx\s+-?\d+$` // matches (addx -11) or (addx 5)

type Signal struct {
	cycle               int        // Instruction cycle value
	currentValue        int        // Current sum of value of addx
	targetStrengthCycle int        //Target cycle to calculate strength e.g 20, 60, 100..
	strength            []Strength // Cycle signal strength accumulator
}

type Strength struct {
	cycle, activeCycleValue, strengthValue int
}

// watchCycle processes instruction cycle and value
func (s *Signal) watchCycle(value int, instruction string) {
	//Incase the target cycle is reached before the end of the cycle of the given instruction
	// this is expected during [addx] of 2 cycles
	if instruction == "addx" {
		s.cycle += 1
		if s.cycle == s.targetStrengthCycle {
			s.targetStrengthCycle += 40 // Intial target, 20 would have been elapsed already
			s.strength = append(s.strength, Strength{s.cycle, s.currentValue, s.cycle * s.currentValue})
		}
	}

	s.currentValue += value // Compute the value during the last cycle of a given instruction
	s.cycle += 1
	if s.cycle == s.targetStrengthCycle {
		s.targetStrengthCycle += 40 // Intial target, 20 would have been elapsed already
		s.strength = append(s.strength, Strength{s.cycle, s.currentValue, s.cycle * s.currentValue})
	}
}

func main() {
	fmt.Println("\nEnter [noop] or [addx (integer)] \n ")
	//Collect input from terminal

	scanner := bufio.NewScanner(os.Stdin)
	signal := Signal{targetStrengthCycle: 20, currentValue: 1}
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())

		if len(input) == 0 {
			break
		}

		matchdAddx, err := regexp.MatchString(addxPattern, input)
		panicError(err)
		if input != "noop" && !matchdAddx {
			panicError(fmt.Errorf("Invalid input. Enter [noop] or [addx (integer)] as the correct format"))
		}

		value := 0
		instruction := "noop"

		if matchdAddx {
			//Extract the value after 'addx'
			inputValue := strings.TrimSpace(input[4:])
			value, err = strconv.Atoi(inputValue) // covert value to integer
			instruction = "addx"
		}
		signal.watchCycle(value, instruction)
	}
	panicError(scanner.Err())
	// fmt.Println(signal)
	totalStrength := 0
	for _, strength := range signal.strength {
		totalStrength += strength.strengthValue
		fmt.Printf("\n During the %vth cycle, register X has the value %v, so the signal strength is %v * %v = %v",
			strength.cycle, strength.activeCycleValue, strength.cycle, strength.activeCycleValue,
			strength.strengthValue,
		)
	}
	fmt.Println("\n Total sum of the signal strengths is: ", totalStrength)
}

func panicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

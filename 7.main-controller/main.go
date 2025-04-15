package main

import (
	"fmt"
	"github/elliot9/class7/core"
	"github/elliot9/class7/entities/commands"
	"github/elliot9/class7/entities/devices"
	"github/elliot9/class7/interfaces"
)

func main() {
	tank := devices.NewTank()
	telecom := devices.NewTelecom()
	controller := core.NewMainController()

	allCommands := []interfaces.Command{
		commands.NewMoveForwardTankCommand(tank),
		commands.NewMoveBackwardTankCommand(tank),
		commands.NewConnectTelecomCommand(telecom),
		commands.NewDisconnectTelecomCommand(telecom),
		commands.NewResetMainControlKeyboardCommand(controller),
	}

	controller.SetCommands(allCommands)

	for {
		printOperation(controller)
		var input string
		fmt.Scanln(&input)

		if input == "1" {
			setHotKey(controller)
		} else if input == "2" {
			controller.Undo()
		} else if input == "3" {
			controller.Redo()
		} else {
			err := controller.Invoke(input)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func printOperation(controller *core.MainController) {
	controller.PrintfKeyboard()
	fmt.Print("(1) 快捷鍵設置 (2) Undo (3) Redo (字母) 按下按鍵: ")
}

func setHotKey(controller *core.MainController) {
	fmt.Println("設置巨集指令 (y/n)：")
	var input string
	fmt.Scanln(&input)

	if input == "y" {
		fmt.Println("請輸入快捷鍵: ")
		var key string
		fmt.Scanln(&key)

		fmt.Printf("要將哪些指令設置成快捷鍵 %s 的巨集（輸入多個數字，以空白隔開）:  \n", key)
		controller.PrintfAllCommands()

		command := controller.RecordMarco()
		controller.Bind(key, command)
		return
	}

	fmt.Println("請輸入快捷鍵: ")
	var key string
	fmt.Scanln(&key)

	fmt.Printf("要將哪一道指令設置到快捷鍵 %s 上: \n", key)
	controller.PrintfAllCommands()
	var commandIndex int
	fmt.Scanln(&commandIndex)

	controller.Bind(key, controller.GetCommands(commandIndex))
}

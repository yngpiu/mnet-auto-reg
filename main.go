package main

import "os"

func main() {
	cfg := loadConfig()

	for {
		showMenu(cfg)
		choice := askQuestion("Select")

		switch choice {
		case "1":
			handleCreateAccounts(cfg)
		case "2":
			cfg = handleToggleMode(cfg)
		case "3":
			cfg = handleChangePassword(cfg)
		case "4":
			printBold("\n  Bye!\n")
			os.Exit(0)
		default:
			printYellow("  Invalid choice!")
		}
	}
}

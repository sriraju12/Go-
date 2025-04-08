package main

import (
	"fmt"
	"example.com/myapp/fileops"
)



func main() {

	fileName := "balance.txt"

	var accountBalance, err = fileops.ReadValueFloatFromFile(fileName)

	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("Creating a new file with default balance of 1000")
        fmt.Println("#######################################################")
	}

	fmt.Println("Welcome to the Bank!")

	for {

		selectOptions()
		
		var choice int
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)

		case 2:
			var depositAmount float64
			fmt.Println("Enter the amount to deposit:")
			fmt.Scanln(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Please Deposit amount greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your updated balance is:", accountBalance)
			fileops.WriteValueFloatToFile(accountBalance,fileName)

		case 3:
			var withdrawAmount float64
			fmt.Println("Enter the amount to withdraw:")
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("Please withdraw amount greater than 0 and less than or equal to your balance")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Your updated balance is:", accountBalance)
				fileops.WriteValueFloatToFile(accountBalance,fileName)
			}

		case 4:
			fmt.Println("Thank you for using the Bank!")
			return
		default:
			fmt.Println("Invalid choice!")

		}
	}
}


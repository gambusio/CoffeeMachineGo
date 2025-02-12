/*
	Coffee machine simulation
	2025 Felipe

	Added features:
		- plant-based milk alternatives (soy milk)
		- add the extra-large cup for regular black coffee
		- extras: sugar and caramel syrup
*/

// Source https://github.com/gambusio/CoffeeMachineGo.git

package main

import (
	"fmt"
)

var waterAmount, milkAmount, coffeeAmount, dispoCups, dispoCupsXL int
var cash int
var sugarAmount, caramel, soyMilkAmount int

/*
initMachine() Inicializa la cafetera con las cantidades inciales
*/
func initMachine() {
	cash = 550
	waterAmount = 400
	milkAmount = 540
	soyMilkAmount = 540
	coffeeAmount = 120
	dispoCups = 9
	dispoCupsXL = 9
	sugarAmount = 1000
	caramel = 1000
}

/*
fill() actualiza la cantidad de consumibles en la cafetera con las cantidades que indique
el operario
*/
func fill() {
	var amount int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&amount)
	waterAmount += amount
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&amount)
	milkAmount += amount
	fmt.Println("Write how many ml of soy milk you want to add:")
	fmt.Scan(&amount)
	soyMilkAmount += amount
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&amount)
	coffeeAmount += amount
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&amount)
	dispoCups += amount
	fmt.Println("Write how many XL disposable cups you want to add:")
	fmt.Scan(&amount)
	dispoCupsXL += amount
	fmt.Println()
}

/*
take() Simula que un operio saca el dinero de la máquina
*/
func take() {
	fmt.Printf("I gave you $%d\n", cash)
	cash = 0
	fmt.Println()
}

/*
printState() Imprime por pantalla el estado actual de la cafetera
*/
func printState() {
	fmt.Printf("The coffee machine has:\n")
	fmt.Printf("%d ml of water\n", waterAmount)
	fmt.Printf("%d ml of milk\n", milkAmount)
	fmt.Printf("%d ml of soy milk\n", soyMilkAmount)
	fmt.Printf("%d g of coffee beans\n", coffeeAmount)
	fmt.Printf("%d g of sugars\n", sugarAmount)
	fmt.Printf("%d disposable cups\n", dispoCups)
	fmt.Printf("%d xl disposable cups\n", dispoCupsXL)
	fmt.Printf("$%d of money\n", cash)
	fmt.Println()
}

/*
Simula el comportamiento de la cafetera cuando un cliente pide un café
*/
func buy() {
	var selection string
	var option string
	//var sugar string

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
	fmt.Scan(&selection)
	if dispoCups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
		fmt.Println()
		return
	}
	switch selection {
	case "1": //expresso 250 ml agua 16 gr café 4 dólares
		// normal size or XL
		fmt.Println("What size do you want to buy? 1 - regular, 2 - XL Cup")
		fmt.Scan(&option)
		if option == "1" {
			if waterAmount < 250 {
				fmt.Println("Sorry, not enough water!")
				break
			}
			if coffeeAmount < 16 {
				fmt.Println("Sorry, not enough coffee!")
				break
			}
			waterAmount -= 250
			coffeeAmount -= 16
			cash += 4
			dispoCups--
		} else if option == "2" {
			if waterAmount < 400 {
				fmt.Println("Sorry, not enough water!")
				break
			}
			if coffeeAmount < 25 {
				fmt.Println("Sorry, not enough coffee!")
				break
			}
			waterAmount -= 400
			coffeeAmount -= 25
			cash += 6
			dispoCupsXL--
		} else {
			break
		}
		fmt.Println("I have enough resources, making you a coffee!")
	case "2": //latte 350 ml agua 75 ml leche 20 gr café 7 dólares
		if waterAmount < 350 {
			fmt.Println("Sorry, not enough water!")
			break
		}
		if coffeeAmount < 20 {
			fmt.Println("Sorry, not enough coffee!")
			break
		}
		// Milk or soy milk
		fmt.Println("What milk do you want? 1 - Regular, 2 - Soy Milk")
		fmt.Scan(&option)
		if option == "1" {
			if milkAmount < 75 {
				fmt.Println("Sorry, not enough milk!")
				break
			}
			milkAmount -= 75
		} else if option == "2" {
			if soyMilkAmount < 75 {
				fmt.Println("Sorry, not enough  soy milk!")
				break
			}
			soyMilkAmount -= 75
		} else {
			break
		}
		waterAmount -= 350
		coffeeAmount -= 20
		cash += 7
		dispoCups--
		fmt.Println("I have enough resources, making you a coffee!")
	case "3": //cappuccino 200 ml agua 100 ml leche 12 g café 6 dólares
		if waterAmount < 200 {
			fmt.Println("Sorry, not enough water!")
			break
		}
		if coffeeAmount < 12 {
			fmt.Println("Sorry, not enough coffee!")
			break
		}
		// Milk or soy milk
		fmt.Println("What milk do you want? 1 - Regular, 2 - Soy Milk")
		fmt.Scan(&option)
		if option == "1" {
			if milkAmount < 100 {
				fmt.Println("Sorry, not enough milk!")
				break
			}
			milkAmount -= 100
		} else if option == "2" {
			if soyMilkAmount < 100 {
				fmt.Println("Sorry, not enough  soy milk!")
				break
			}
			soyMilkAmount -= 100
		} else {
			break
		}
		milkAmount -= 100
		waterAmount -= 200
		coffeeAmount -= 12
		cash += 6
		dispoCups--
		fmt.Println("I have enough resources, making you a coffee!")
	}
	fmt.Println()
}

func main() {
	var action string
	initMachine()
	for action != "exit" {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)
		fmt.Println()
		switch action {
		case "fill":

			fill()
		case "take":

			take()
		case "buy":

			buy()
		case "remaining":
			printState()
		}
	}
}

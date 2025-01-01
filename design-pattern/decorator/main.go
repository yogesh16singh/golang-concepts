package main

import "fmt"

type IIceCream interface {
	GetIceCream() string
}

type VanillaIceCream struct{}

func (v *VanillaIceCream) GetIceCream() string {
	return "Vanilla IceCream"
}

type ChocolateFrostingDecorator struct {
	IceCream IIceCream
}

func (c *ChocolateFrostingDecorator) GetIceCream() string {
	return c.IceCream.GetIceCream() + " with Chocolate Frosting"
}

type CaramelSauceDecorator struct {
	IceCream IIceCream
}

func (c *CaramelSauceDecorator) GetIceCream() string {
	return c.IceCream.GetIceCream() + " with Caramel Sauce"
}

type BananaIceCream struct{}

func (b *BananaIceCream) GetIceCream() string {
	return "Banana IceCream"
}

func main() {
	iceCreamOne := &VanillaIceCream{}
	iceCreamOneWithChocolateFrosting := &ChocolateFrostingDecorator{IceCream: iceCreamOne}
	fmt.Println("iceCream1: " + iceCreamOneWithChocolateFrosting.GetIceCream())

	iceCreamTwo := &VanillaIceCream{}
	iceCreamTwoWithChocoFrosting := &ChocolateFrostingDecorator{IceCream: iceCreamTwo}
	iceCreamTwoWithChocoFrWithCaramelSauce := &CaramelSauceDecorator{IceCream: iceCreamTwoWithChocoFrosting}
	fmt.Println("iceCream2: " + iceCreamTwoWithChocoFrWithCaramelSauce.GetIceCream())

	iceCreamThree := &VanillaIceCream{}
	iceCreamThreeWithCaramelFrostingOnly := &CaramelSauceDecorator{IceCream: iceCreamThree}
	fmt.Println("iceCream3: " + iceCreamThreeWithCaramelFrostingOnly.GetIceCream())

	iceCreamFour := &BananaIceCream{}
	iceCreamFourWithCaramelSauceOnly := &CaramelSauceDecorator{IceCream: iceCreamFour}
	fmt.Println("iceCream4: " + iceCreamFourWithCaramelSauceOnly.GetIceCream())
}

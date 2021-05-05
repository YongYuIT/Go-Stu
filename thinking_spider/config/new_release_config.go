package config

import "strings"

func InitHomeAndKitchen(HomeAndKitchen map[string]interface{}) {
	HomeAndKitchen["Bath"] = make(map[string]interface{})
	HomeAndKitchen["Bedding"] = make(map[string]interface{})
	HomeAndKitchen["Cleaning Supplies"] = make(map[string]interface{})
	HomeAndKitchen["Furniture"] = make(map[string]interface{})
	HomeAndKitchen["Heating, Cooling & Air Quality"] = make(map[string]interface{})
	HomeAndKitchen["Home Décor"] = make(map[string]interface{})
	HomeAndKitchen["Irons & Steamers"] = make(map[string]interface{})
	HomeAndKitchen["Kids' Home Store"] = make(map[string]interface{}) //Need
	HomeAndKitchen["Kitchen & Dining"] = make(map[string]interface{}) //Need
	HomeAndKitchen["Seasonal Décor"] = make(map[string]interface{})
	HomeAndKitchen["Storage & Organization"] = make(map[string]interface{})
	HomeAndKitchen["Vacuums & Floor Care"] = make(map[string]interface{})
	HomeAndKitchen["Wall Décor"] = make(map[string]interface{})

	initKitchenAndDining(HomeAndKitchen["Kitchen & Dining"].(map[string]interface{}))
	initKidsHomeStore(HomeAndKitchen["Kids' Home Store"].(map[string]interface{}))
}

func initKidsHomeStore(KidsHomeStore map[string]interface{}) {
	KidsHomeStore["Kids' Baking Supplies"] = make(map[string]interface{})
	KidsHomeStore["Kids' Bedding"] = make(map[string]interface{})
	KidsHomeStore["Kids' Furniture"] = make(map[string]interface{})
	KidsHomeStore["Kids' Room Décor"] = make(map[string]interface{})
	KidsHomeStore["Nursery Décor"] = make(map[string]interface{})
	KidsHomeStore["Nursery Furniture"] = make(map[string]interface{})
	KidsHomeStore["Nursery Bedding"] = make(map[string]interface{})
	KidsHomeStore["Kids' Bath"] = make(map[string]interface{})
}

func initKitchenAndDining(KitchenAndDining map[string]interface{}) {
	KitchenAndDining["Bakeware"] = make(map[string]interface{})
	KitchenAndDining["Coffee, Tea & Espresso Appliances"] = make(map[string]interface{})
	KitchenAndDining["Cookware"] = make(map[string]interface{})
	KitchenAndDining["Dining & Entertaining"] = make(map[string]interface{})
	KitchenAndDining["Home Brewing & Wine Making"] = make(map[string]interface{})
	KitchenAndDining["Kitchen & Table Linens"] = make(map[string]interface{})
	KitchenAndDining["Kitchen Utensils & Gadgets"] = make(map[string]interface{})
	KitchenAndDining["Small Appliance Parts & Accessories"] = make(map[string]interface{})
	KitchenAndDining["Small Appliances"] = make(map[string]interface{})
	KitchenAndDining["Storage & Organization"] = make(map[string]interface{})
	KitchenAndDining["Travel & To-Go Drinkware"] = make(map[string]interface{})
	KitchenAndDining["Water Coolers & Filters"] = make(map[string]interface{})
	KitchenAndDining["Wine Accessories"] = make(map[string]interface{})

	initBakeware(KitchenAndDining["Bakeware"].(map[string]interface{}))
	initDiningAndEntertaining(KitchenAndDining["Dining & Entertaining"].(map[string]interface{}))
	initHomeBrewingAndWineMaking(KitchenAndDining["Home Brewing & Wine Making"].(map[string]interface{}))
	initKitchenAndTableLinens(KitchenAndDining["Kitchen & Table Linens"].(map[string]interface{}))
	initKitchenUtensilsAndGadgets(KitchenAndDining["Kitchen Utensils & Gadgets"].(map[string]interface{}))
	initSmallAppliancePartsAndAccessories(KitchenAndDining["Small Appliance Parts & Accessories"].(map[string]interface{}))
	initSmallAppliances(KitchenAndDining["Small Appliances"].(map[string]interface{}))
	initStorageAndOrganization(KitchenAndDining["Storage & Organization"].(map[string]interface{}))
	initTravelAndToGoDrinkware(KitchenAndDining["Travel & To-Go Drinkware"].(map[string]interface{}))
	initWaterCoolersAndFilters(KitchenAndDining["Water Coolers & Filters"].(map[string]interface{}))
	initWineAccessories(KitchenAndDining["Wine Accessories"].(map[string]interface{}))

}

func initWineAccessories(WineAccessories map[string]interface{}) {
	WineAccessories["Wine Education & Games"] = make(map[string]interface{})
	WineAccessories["Wine Racks & Cabinets"] = make(map[string]interface{})
	WineAccessories["Wine Accessory Sets"] = make(map[string]interface{})
}

func initWaterCoolersAndFilters(WaterCoolersAndFilters map[string]interface{}) {
	WaterCoolersAndFilters["Pitcher Water Filters"] = make(map[string]interface{})
	WaterCoolersAndFilters["Water Coolers"] = make(map[string]interface{})
	WaterCoolersAndFilters["Water Filter Cleaners"] = make(map[string]interface{})
}

func initTravelAndToGoDrinkware(TravelAndToGoDrinkware map[string]interface{}) {
	TravelAndToGoDrinkware["Water Bottles"] = make(map[string]interface{})
}

func initStorageAndOrganization(StorageAndOrganization map[string]interface{}) {
	StorageAndOrganization["Cabinet & Drawer Organization"] = make(map[string]interface{})
	StorageAndOrganization["Countertop & Wall Organization"] = make(map[string]interface{})
	StorageAndOrganization["Dinnerware & Stemware Storage"] = make(map[string]interface{})
	StorageAndOrganization["Flatware & Utensil Storage"] = make(map[string]interface{})
	StorageAndOrganization["Food Storage"] = make(map[string]interface{})
	StorageAndOrganization["Kitchen Storage & Organization Accessories"] = make(map[string]interface{})
	StorageAndOrganization["Thermoses"] = make(map[string]interface{})
	StorageAndOrganization["Travel & To-Go Food Containers"] = make(map[string]interface{})
	StorageAndOrganization["Wine Racks & Cabinets"] = make(map[string]interface{})
}

func initSmallAppliances(SmallAppliances map[string]interface{}) {
	SmallAppliances["Blenders"] = make(map[string]interface{})
	SmallAppliances["Bread Machines"] = make(map[string]interface{})
	SmallAppliances["Coffee, Tea & Espresso Appliances"] = make(map[string]interface{})
	SmallAppliances["Compact Refrigerators"] = make(map[string]interface{})
	SmallAppliances["Contact Grills"] = make(map[string]interface{})
	SmallAppliances["Countertop Burners"] = make(map[string]interface{})
	SmallAppliances["Electric Cake Pop & Mini Cake Makers"] = make(map[string]interface{})
	SmallAppliances["Electric Griddles"] = make(map[string]interface{})
	SmallAppliances["Electric Pressure Cookers"] = make(map[string]interface{})
	SmallAppliances["Electric Skillets"] = make(map[string]interface{})
	SmallAppliances["Electric Woks"] = make(map[string]interface{})
	SmallAppliances["Food Processors"] = make(map[string]interface{})
	SmallAppliances["Fryers"] = make(map[string]interface{})
	SmallAppliances["Hot Pots"] = make(map[string]interface{})
	SmallAppliances["Ice Cream Machines"] = make(map[string]interface{})
	SmallAppliances["Juicers"] = make(map[string]interface{})
	SmallAppliances["Microwave Ovens"] = make(map[string]interface{})
	SmallAppliances["Mixers"] = make(map[string]interface{})
	SmallAppliances["Ovens & Toasters"] = make(map[string]interface{})
	SmallAppliances["Rice Cookers"] = make(map[string]interface{})
	SmallAppliances["Slow Cookers"] = make(map[string]interface{})
	SmallAppliances["Soda Makers"] = make(map[string]interface{})
	SmallAppliances["Specialty Appliances"] = make(map[string]interface{})
	SmallAppliances["Steamers"] = make(map[string]interface{})
	SmallAppliances["Waffle Irons"] = make(map[string]interface{})
	SmallAppliances["Wine Cellars"] = make(map[string]interface{})

	initBlenders(SmallAppliances["Blenders"].(map[string]interface{}))
	initCoffeeTeaAndEspressoAppliances(SmallAppliances["Coffee, Tea & Espresso Appliances"].(map[string]interface{}))
	initElectricCakePopAndMiniCakeMakers(SmallAppliances["Electric Cake Pop & Mini Cake Makers"].(map[string]interface{}))
	initFryers(SmallAppliances["Fryers"].(map[string]interface{}))
	initJuicers(SmallAppliances["Juicers"].(map[string]interface{}))
	initMicrowaveOvens(SmallAppliances["Microwave Ovens"].(map[string]interface{}))
	initMixers(SmallAppliances["Mixers"].(map[string]interface{}))
	initOvensAndToasters(SmallAppliances["Ovens & Toasters"].(map[string]interface{}))
	initWineCellars(SmallAppliances["Wine Cellars"].(map[string]interface{}))

}

func initWineCellars(WineCellars map[string]interface{}) {
	WineCellars["Built-In Wine Cellars"] = make(map[string]interface{})
	WineCellars["Freestanding Wine Cellars"] = make(map[string]interface{})
	WineCellars["Wine Cellar Cooling Systems"] = make(map[string]interface{})
}

func initOvensAndToasters(OvensAndToasters map[string]interface{}) {
	OvensAndToasters["Convection Ovens"] = make(map[string]interface{})
	OvensAndToasters["Countertop Pizza Ovens"] = make(map[string]interface{})
	OvensAndToasters["Rotisseries & Roasters"] = make(map[string]interface{})
	OvensAndToasters["Toaster Ovens"] = make(map[string]interface{})
	OvensAndToasters["Toasters"] = make(map[string]interface{})
}

func initMixers(Mixers map[string]interface{}) {
	Mixers["Hand Mixers"] = make(map[string]interface{})
	Mixers["Stand Mixers"] = make(map[string]interface{})
}

func initMicrowaveOvens(MicrowaveOvens map[string]interface{}) {
	MicrowaveOvens["Compact Microwave Ovens"] = make(map[string]interface{})
	MicrowaveOvens["Countertop Microwave Ovens"] = make(map[string]interface{})
	MicrowaveOvens["Microhood Microwave Ovens"] = make(map[string]interface{})
	MicrowaveOvens["Over-the-Range Microwave Ovens"] = make(map[string]interface{})
	MicrowaveOvens["Speed-Cooking Microwave Ovens"] = make(map[string]interface{})
}

func initJuicers(Juicers map[string]interface{}) {
	Juicers["Centrifugal Juicers"] = make(map[string]interface{})
	Juicers["Citrus Juicers"] = make(map[string]interface{})
	Juicers["Masticating Juicers"] = make(map[string]interface{})
}

func initFryers(Fryers map[string]interface{}) {
	Fryers["Air Fryers"] = make(map[string]interface{})
	Fryers["Deep Fryers"] = make(map[string]interface{})
}

func initElectricCakePopAndMiniCakeMakers(ElectricCakePopAndMiniCakeMakers map[string]interface{}) {
	ElectricCakePopAndMiniCakeMakers["Cake Pop Makers"] = make(map[string]interface{})
	ElectricCakePopAndMiniCakeMakers["Cupcake Makers"] = make(map[string]interface{})
	ElectricCakePopAndMiniCakeMakers["Mini Donut Makers"] = make(map[string]interface{})
}

func initCoffeeTeaAndEspressoAppliances(CoffeeTeaAndEspressoAppliances map[string]interface{}) {
	CoffeeTeaAndEspressoAppliances["Beverage Warmers"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Coffee Makers"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Electric Blade Grinders"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Electric Burr Grinders"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Espresso Machine & Coffeemaker Combos"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Espresso Machines"] = make(map[string]interface{})
	CoffeeTeaAndEspressoAppliances["Kettles & Tea Machines"] = make(map[string]interface{})

}

func initBlenders(Blenders map[string]interface{}) {
	Blenders["Countertop Blenders"] = make(map[string]interface{})
	Blenders["Hand Blenders"] = make(map[string]interface{})
	Blenders["Personal Size Blenders"] = make(map[string]interface{})
}

func initSmallAppliancePartsAndAccessories(SmallAppliancePartsAndAccessories map[string]interface{}) {
	SmallAppliancePartsAndAccessories["Blender Replacement Parts"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Bread Machine Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Coffee & Espresso Machine Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Deep Fryer Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Food Processor Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Juicer Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Microwave Oven Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Mixer Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Pressure Cooker Parts & Accessories"] = make(map[string]interface{})
	SmallAppliancePartsAndAccessories["Soda Maker Parts & Accessories"] = make(map[string]interface{})

	initCoffeeAndEspressoMachinePartsAndAccessories(SmallAppliancePartsAndAccessories["Coffee & Espresso Machine Parts & Accessories"].(map[string]interface{}))
	initMicrowaveOvenPartsAndAccessories(SmallAppliancePartsAndAccessories["Microwave Oven Parts & Accessories"].(map[string]interface{}))
	initPressureCookerPartsAndAccessories(SmallAppliancePartsAndAccessories["Pressure Cooker Parts & Accessories"].(map[string]interface{}))
}

func initPressureCookerPartsAndAccessories(PressureCookerPartsAndAccessories map[string]interface{}) {
	PressureCookerPartsAndAccessories["Accessories"] = make(map[string]interface{})
	PressureCookerPartsAndAccessories["Replacement Parts"] = make(map[string]interface{})
}

func initMicrowaveOvenPartsAndAccessories(MicrowaveOvenPartsAndAccessories map[string]interface{}) {
	MicrowaveOvenPartsAndAccessories["Filters"] = make(map[string]interface{})
	MicrowaveOvenPartsAndAccessories["Microwave Oven Replacement Parts"] = make(map[string]interface{})
	MicrowaveOvenPartsAndAccessories["Turntables"] = make(map[string]interface{})
}

func initCoffeeAndEspressoMachinePartsAndAccessories(CoffeeAndEspressoMachinePartsAndAccessories map[string]interface{}) {
	CoffeeAndEspressoMachinePartsAndAccessories["Coffee & Espresso Machine Cleaning Products"] = make(map[string]interface{})
	CoffeeAndEspressoMachinePartsAndAccessories["Coffee Machine Accessories"] = make(map[string]interface{})
	CoffeeAndEspressoMachinePartsAndAccessories["Espresso Machine Accessories"] = make(map[string]interface{})
	CoffeeAndEspressoMachinePartsAndAccessories["Espresso Machine Replacement Parts"] = make(map[string]interface{})
}

func initKitchenUtensilsAndGadgets(KitchenUtensilsAndGadgets map[string]interface{}) {
	KitchenUtensilsAndGadgets["Cooking Utensils"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Measuring Tools & Scales"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Tool & Gadget Sets"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Graters, Peelers & Slicers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Can Openers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Jar Openers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Seasoning & Spice Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Colanders & Food Strainers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Specialty Tools & Gadgets"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Thermometers & Timers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Oil Sprayers & Dispensers"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Fruit & Vegetable Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Pasta & Pizza Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Meat & Poultry Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Salad Tools & Spinners"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Cheese Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Baking Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Barbecue Tools"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Kitchen Accessories"] = make(map[string]interface{})
	KitchenUtensilsAndGadgets["Straws"] = make(map[string]interface{})
}

func initKitchenAndTableLinens(KitchenAndTableLinens map[string]interface{}) {
	KitchenAndTableLinens["Aprons"] = make(map[string]interface{})
	KitchenAndTableLinens["Chair Pads"] = make(map[string]interface{})
	KitchenAndTableLinens["Dish Cloths & Dish Towels"] = make(map[string]interface{})
	KitchenAndTableLinens["Kitchen Linen Sets"] = make(map[string]interface{})
	KitchenAndTableLinens["Kitchen Rugs"] = make(map[string]interface{})
	KitchenAndTableLinens["Potholders & Oven Mitts"] = make(map[string]interface{})
	KitchenAndTableLinens["Tea Cozies"] = make(map[string]interface{})
	KitchenAndTableLinens["Accessories"] = make(map[string]interface{})

	initPotholdersAndOvenMitts(KitchenAndTableLinens["Potholders & Oven Mitts"].(map[string]interface{}))
	initAccessories(KitchenAndTableLinens["Accessories"].(map[string]interface{}))

}

func initAccessories(Accessories map[string]interface{}) {
	Accessories["Place Cards & Place Card Holders"] = make(map[string]interface{})
}

func initPotholdersAndOvenMitts(PotholdersAndOvenMitts map[string]interface{}) {
	PotholdersAndOvenMitts["Oven Mitts"] = make(map[string]interface{})
	PotholdersAndOvenMitts["Potholders"] = make(map[string]interface{})
}

func initHomeBrewingAndWineMaking(HomeBrewingAndWineMaking map[string]interface{}) {
	HomeBrewingAndWineMaking["Beer Brewing"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Cleaning & Sanitizing"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Fermentation & More"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Labeling Supplies"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Measuring & Testing"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Racking & Storage"] = make(map[string]interface{})
	HomeBrewingAndWineMaking["Wine Making"] = make(map[string]interface{})

	initBeerBrewing(HomeBrewingAndWineMaking["Beer Brewing"].(map[string]interface{}))
	initFermentationAndMore(HomeBrewingAndWineMaking["Fermentation & More"].(map[string]interface{}))
	initWineMaking(HomeBrewingAndWineMaking["Wine Making"].(map[string]interface{}))
}

func initWineMaking(WineMaking map[string]interface{}) {
	WineMaking["Barrels"] = make(map[string]interface{})
	WineMaking["Bottling & Corking"] = make(map[string]interface{})
	WineMaking["Crushing, Pressing & Stemming"] = make(map[string]interface{})
	WineMaking["Wine Filters"] = make(map[string]interface{})
	WineMaking["Wine Making Starter Sets"] = make(map[string]interface{})
}

func initFermentationAndMore(FermentationAndMore map[string]interface{}) {
	FermentationAndMore["Aeration Equipment"] = make(map[string]interface{})
	FermentationAndMore["Airlocks"] = make(map[string]interface{})
	FermentationAndMore["Carboys"] = make(map[string]interface{})
	FermentationAndMore["Fermenters"] = make(map[string]interface{})
	FermentationAndMore["Heating & Temperature Control"] = make(map[string]interface{})
	FermentationAndMore["Siphoning"] = make(map[string]interface{})
}

func initBeerBrewing(BeerBrewing map[string]interface{}) {
	BeerBrewing["Bottles & Bottling"] = make(map[string]interface{})
	BeerBrewing["Brew Pots, Kettles & Accessories"] = make(map[string]interface{})
	BeerBrewing["Burners"] = make(map[string]interface{})
	BeerBrewing["Filters & Straining"] = make(map[string]interface{})
	BeerBrewing["Grain Mills"] = make(map[string]interface{})
	BeerBrewing["Home Brewing Starter Sets"] = make(map[string]interface{})
	BeerBrewing["Kegs & Kegging"] = make(map[string]interface{})
	BeerBrewing["Stirring & Utensils"] = make(map[string]interface{})
	BeerBrewing["Wort Chillers"] = make(map[string]interface{})
}

func initDiningAndEntertaining(DiningAndEntertaining map[string]interface{}) {
	DiningAndEntertaining["Bar Tools & Drinkware"] = make(map[string]interface{})
	DiningAndEntertaining["Novelty"] = make(map[string]interface{})
	DiningAndEntertaining["Tabletop Accessories"] = make(map[string]interface{})

	initBarToolsAndDrinkware(DiningAndEntertaining["Bar Tools & Drinkware"].(map[string]interface{}))
	initNovelty(DiningAndEntertaining["Novelty"].(map[string]interface{}))
	initTabletopAccessories(DiningAndEntertaining["Tabletop Accessories"].(map[string]interface{}))
}

func initTabletopAccessories(TabletopAccessories map[string]interface{}) {

	TabletopAccessories["Candlesticks"] = make(map[string]interface{})
	TabletopAccessories["Linens"] = make(map[string]interface{})

}

func initNovelty(Novelty map[string]interface{}) {
	Novelty["Bowls"] = make(map[string]interface{})
	Novelty["Drinkware"] = make(map[string]interface{})
	Novelty["Flatware"] = make(map[string]interface{})
	Novelty["Plates"] = make(map[string]interface{})
	Novelty["Serveware"] = make(map[string]interface{})
}

func initBarToolsAndDrinkware(BarToolsAndDrinkware map[string]interface{}) {
	BarToolsAndDrinkware["Bar Tools"] = make(map[string]interface{})
	initBarTools(BarToolsAndDrinkware["Bar Tools"].(map[string]interface{}))
}

func initBarTools(BarTools map[string]interface{}) {
	BarTools["Cocktail Picks"] = make(map[string]interface{})
	BarTools["Swizzle Sticks"] = make(map[string]interface{})
}

func initBakeware(Bakeware map[string]interface{}) {
	Bakeware["Bakers & Casseroles"] = make(map[string]interface{})
	Bakeware["Bakeware Sets"] = make(map[string]interface{})
	Bakeware["Baking & Cookie Sheets"] = make(map[string]interface{})
	Bakeware["Baking Tools & Accessories"] = make(map[string]interface{})
	Bakeware["Bread & Loaf Pans"] = make(map[string]interface{})
	Bakeware["Cake Pans"] = make(map[string]interface{})
	Bakeware["Electric Cake Pop & Mini Cake Makers"] = make(map[string]interface{})
	Bakeware["Candy Making Supplies"] = make(map[string]interface{})
	Bakeware["Decorating Tools"] = make(map[string]interface{})
	Bakeware["Jelly-Roll Pans"] = make(map[string]interface{})
	Bakeware["Kids' Baking Supplies"] = make(map[string]interface{})
	Bakeware["Mixing Bowls"] = make(map[string]interface{})
	Bakeware["Muffin & Cupcake Pans"] = make(map[string]interface{})
	Bakeware["Pastry & Baking Molds"] = make(map[string]interface{})
	Bakeware["Pie, Tart & Quiche Pans"] = make(map[string]interface{})
	Bakeware["Pizza Pans & Stones"] = make(map[string]interface{})
	Bakeware["Popover Pans"] = make(map[string]interface{})
	Bakeware["Ramekins & Soufflé Dishes"] = make(map[string]interface{})
	Bakeware["Roasting Pans"] = make(map[string]interface{})

	initBakersAndCasseroles(Bakeware["Bakers & Casseroles"].(map[string]interface{}))
	initBakingToolsAndAccessories(Bakeware["Baking Tools & Accessories"].(map[string]interface{}))
	initBreadAndLoafPans(Bakeware["Bread & Loaf Pans"].(map[string]interface{}))
	initCakePans(Bakeware["Cake Pans"].(map[string]interface{}))
	initCandyMakingSupplies(Bakeware["Candy Making Supplies"].(map[string]interface{}))
	initDecoratingTools(Bakeware["Decorating Tools"].(map[string]interface{}))
	initPastryAndBakingMolds(Bakeware["Pastry & Baking Molds"].(map[string]interface{}))
	initPieTartAndQuichePans(Bakeware["Pie, Tart & Quiche Pans"].(map[string]interface{}))
	initRamekinsAndSouffléDishes(Bakeware["Ramekins & Soufflé Dishes"].(map[string]interface{}))

}

func initRamekinsAndSouffléDishes(RamekinsAndSouffléDishes map[string]interface{}) {
	RamekinsAndSouffléDishes["Ramekins"] = make(map[string]interface{})
	RamekinsAndSouffléDishes["Soufflé Dishes"] = make(map[string]interface{})
}

func initPieTartAndQuichePans(PieTartAndQuichePans map[string]interface{}) {
	PieTartAndQuichePans["Pie Pans"] = make(map[string]interface{})
	PieTartAndQuichePans["Quiche Pans"] = make(map[string]interface{})
	PieTartAndQuichePans["Tart Pans"] = make(map[string]interface{})
}

func initPastryAndBakingMolds(PastryAndBakingMolds map[string]interface{}) {
	PastryAndBakingMolds["Brioche Pans"] = make(map[string]interface{})
	PastryAndBakingMolds["Cake Pop Pans"] = make(map[string]interface{})
	PastryAndBakingMolds["Donut Pans"] = make(map[string]interface{})
	PastryAndBakingMolds["Fondant & Gum Paste Molds"] = make(map[string]interface{})
	PastryAndBakingMolds["Macaron Baking Mats & Pans"] = make(map[string]interface{})
	PastryAndBakingMolds["Madeleine Pans"] = make(map[string]interface{})
}

func initDecoratingTools(DecoratingTools map[string]interface{}) {
	DecoratingTools["Cake Decorating Kits"] = make(map[string]interface{})
	DecoratingTools["Decorating & Pastry Bags"] = make(map[string]interface{})
	DecoratingTools["Edible Ink Printers"] = make(map[string]interface{})
	DecoratingTools["Edible Printer Ink & Paper"] = make(map[string]interface{})
	DecoratingTools["Food Coloring"] = make(map[string]interface{})
	DecoratingTools["Icing & Decorating Spatulas"] = make(map[string]interface{})
	DecoratingTools["Icing & Piping Tips"] = make(map[string]interface{})
	DecoratingTools["Icing Dispensers"] = make(map[string]interface{})
	DecoratingTools["Sculpting & Modeling Tools"] = make(map[string]interface{})
	DecoratingTools["Stencils"] = make(map[string]interface{})
	DecoratingTools["Wrapping & Packaging"] = make(map[string]interface{})
}

func initCandyMakingSupplies(CandyMakingSupplies map[string]interface{}) {
	CandyMakingSupplies["Candy, Chocolate & Fondant Molds"] = make(map[string]interface{})
	CandyMakingSupplies["Candy Making Accessories"] = make(map[string]interface{})
}

func initCakePans(CakePans map[string]interface{}) {
	CakePans["Angel Food"] = make(map[string]interface{})
	CakePans["Bundt Pans"] = make(map[string]interface{})
	CakePans["Ring Mold Pans"] = make(map[string]interface{})
	CakePans["Round"] = make(map[string]interface{})
	CakePans["Specialty & Novelty Cake Pans"] = make(map[string]interface{})
	CakePans["Springform"] = make(map[string]interface{})
	CakePans["Square & Rectangular"] = make(map[string]interface{})
}

func initBreadAndLoafPans(BreadAndLoafPans map[string]interface{}) {
	BreadAndLoafPans["French Bread & Baguette Pans"] = make(map[string]interface{})
	BreadAndLoafPans["Loaf Pans"] = make(map[string]interface{})
	BreadAndLoafPans["Specialty Forms"] = make(map[string]interface{})
}

func initBakersAndCasseroles(BakersAndCasseroles map[string]interface{}) {
	BakersAndCasseroles["Au Gratin Pans"] = make(map[string]interface{})
	BakersAndCasseroles["Bake & Serve Sets"] = make(map[string]interface{})
	BakersAndCasseroles["Bakers"] = make(map[string]interface{})
	BakersAndCasseroles["Casseroles"] = make(map[string]interface{})
	BakersAndCasseroles["Lasagna Pans"] = make(map[string]interface{})
	BakersAndCasseroles["Terrines"] = make(map[string]interface{})
}

func initBakingToolsAndAccessories(BakingToolsAndAccessories map[string]interface{}) {
	BakingToolsAndAccessories["Baking & Pastry Utensils"] = make(map[string]interface{})
	BakingToolsAndAccessories["Baking Cups"] = make(map[string]interface{})
	BakingToolsAndAccessories["Baking Mats"] = make(map[string]interface{})
	BakingToolsAndAccessories["Beaters"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cake Rings"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cookie Cutters"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cookie Presses"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cookie Stamps"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cooking Torches"] = make(map[string]interface{})
	BakingToolsAndAccessories["Cooling Racks"] = make(map[string]interface{})
	BakingToolsAndAccessories["Mixing Bowls"] = make(map[string]interface{})
	BakingToolsAndAccessories["Parchment"] = make(map[string]interface{})
	BakingToolsAndAccessories["Rolling Pins"] = make(map[string]interface{})
	BakingToolsAndAccessories["Sifters"] = make(map[string]interface{})
}

func GetMapInMap(title string, bigMap map[string]interface{}) map[string]interface{} {
	for key := range bigMap {
		if strings.EqualFold(key, title) {
			return bigMap[key].(map[string]interface{})
		} else {
			result := GetMapInMap(title, bigMap[key].(map[string]interface{}))
			if result != nil {
				return result
			}
		}
	}
	return nil
}

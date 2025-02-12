package main

import "BudgetApp/database/seeds"

func main() {
	seeds.SeedUsers(10)
	seeds.CategorySeed(10)
}

package main

import (
    "fmt"
)

func main() {

    fmt.Println("✓ Hello, 世界")
    // OR
    fmt.Println("\u2713 Hello, 世界")

fmt.Println("\u250C\u2500\u2500\u2500\u2510\n\u2502 X \u2502\n\u2514\u2500\u2500\u2500\u2518\n")
fmt.Println("┌───┐\n│ O │\n└───┘")
fmt.Println()
//fmt.Println("\u250C\u2500\u2500\u2500\u252C\u2500\u2500\u2500\u252C\u2500\u2500\u2500\u2510")
//fmt.Println("\u251C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u253C\u2500\u2500\u2500\u2524")
//fmt.Println("\u2502 O \u2502 X \u2502 O \u2502")
//fmt.Println("\u2514\u2500\u2500\u2500\u2534\u2500\u2500\u2500\u2534\u2500\u2500\u2500\u2518")
fmt.Println("┌───┬───┬───┐")
fmt.Println("│ O │ X │ O │")
fmt.Println("├───┼───┼───┤")
fmt.Println("│ X │ X │ X │")
fmt.Println("├───┼───┼───┤")
fmt.Println("│   │ X │   │")
fmt.Println("└───┴───┴───┘")

fmt.Printf("┌───┬───┬───┐\n")
fmt.Printf("│ %s│ %s│ %s│\n","X ","O ","O ")
fmt.Printf("├───┼───┼───┤\n")
fmt.Printf("│ %s│ %s│ %s│\n","  ","X ","X ")
fmt.Printf("├───┼───┼───┤\n")
fmt.Printf("│ %s│ %s│ %s│\n","  ","  ","  ")
fmt.Printf("└───┴───┴───┘\n")

l1:="┌───┬───┬───┐   "
l2:="│   │   │   │   "
l3:="├───┼───┼───┤   "
l4:="│   │   │   │   "
l5:="├───┼───┼───┤   "
l6:="│   │   │   │   "
l7:="└───┴───┴───┘   "


fmt.Printf("┌───┬───┬───┐\n")
fmt.Printf("│   │   │   │\n")
fmt.Printf("├───┼───┼───┤\n")
fmt.Printf("│   │   │   │\n")
fmt.Printf("├───┼───┼───┤\n")
fmt.Printf("│   │   │   │\n")
fmt.Printf("└───┴───┴───┘\n")

fmt.Println(l1,l1,l1)
fmt.Println(l2,l2,l2)
fmt.Println(l3,l3,l3)
fmt.Println(l4,l4,l4)
fmt.Println(l5,l5,l5)
fmt.Println(l6,l6,l6)
fmt.Println(l7,l7,l7)

// replaced font 'Droid Sans Mono', 'monospace', monospace with Jetbrains Mono which has better box drawing









}



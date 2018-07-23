package main

import (
    "fmt"
    "flag"
)

func main() {
    dani := flag.Int("dani", 0, "Dani laptopjanak toltottsegi szintje")
    gellert := flag.Int("gellert", 0, "Gellert laptopjanak toltottsegi szintje")

    flag.Parse()

    if (*dani > *gellert) {
        fmt.Println("Gellert laptopjat kell tolteni")
    } else {
        fmt.Println("Dani laptopjat kell tolteni")
    }
}

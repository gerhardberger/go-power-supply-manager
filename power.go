package main

import (
    "fmt"
    "flag"
)

func main() {
    dani := flag.Int("dani", 100, "Dani laptopjanak toltottsegi szintje")
    gellert := flag.Int("gellert", 100, "Gellert laptopjanak toltottsegi szintje")
    gergo := flag.Int("gergo", 100, "Gergo laptopjanak toltottsegi szintje")


    flag.Parse()
    
    switch {
        *dani < *gellert && *dani < *gergo:
            fmt.Println("Dani laptopjat kell tolteni")
        *gellert < *dani && *gellert < *gergo:
            fmt.Println("Gellert laptopjat kell tolteni")
        default:
            fmt.Println("Geri laptopjat kell tolteni")     
    }
}

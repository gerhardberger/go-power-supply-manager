package main

import (
    "fmt"
    "flag"
)

func main() {
    dani := flag.Int("dani", 100, "Dani laptopjanak toltottsegi szintje")
    gabor := flag.Int("gabor", 100, "Gabor laptopjanak toltottsegi szintje")
    gergo := flag.Int("gergo", 100, "Gergo laptopjanak toltottsegi szintje")


    flag.Parse()
    
    switch {
        case *dani < *gabor && *dani < *gergo:
            fmt.Println("Dani laptopjat kell tolteni")
        case *gabor < *dani && *gabor < *gergo:
            fmt.Println("Gabor laptopjat kell tolteni")
        default:
            fmt.Println("Geri laptopjat kell tolteni")     
    }
}

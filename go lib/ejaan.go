package main

import "fmt"

func main() {
    var input string

    fmt.Print("Masukkan kata: ")
    fmt.Scanln(&input)

    for i := 0; i < len(input); i++ {
        fmt.Printf("%c ", input[i])
    }
}

package main

import "fmt"

func main() {
    var input []byte

    fmt.Print("Masukkan teks: ")
    fmt.Scanln(&input)

    fmt.Println("Input: ", string(input))
}

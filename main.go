package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "mini-quiz/models"
    "mini-quiz/utils"
)

func main() {
    // Minta nama pemain
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Masukkan nama Anda: ")
    nameInput, _ := reader.ReadString('\n')
    playerName := strings.TrimSpace(nameInput)

    // Daftar pertanyaan (bisa diperbanyak lagi)
    questions := []models.Question{
        {Text: "Siapa penemu bahasa Go?", Choices: []string{"Ken Thompson & Rob Pike", "Guido van Rossum", "James Gosling", "Bjarne Stroustrup"}, Answer: 1, Category: "Teknologi", Difficulty: "Sedang"},
        {Text: "Tahun berapa Go dirilis pertama kali?", Choices: []string{"2007", "2009", "2012", "2015"}, Answer: 2, Category: "Teknologi", Difficulty: "Mudah"},
        {Text: "Planet terbesar di tata surya?", Choices: []string{"Bumi", "Mars", "Jupiter", "Venus"}, Answer: 3, Category: "Sains", Difficulty: "Mudah"},
        {Text: "Gunung tertinggi di dunia?", Choices: []string{"K2", "Kilimanjaro", "Everest", "Elbrus"}, Answer: 3, Category: "Geografi", Difficulty: "Sedang"},
        {Text: "Siapa pelukis Mona Lisa?", Choices: []string{"Van Gogh", "Leonardo da Vinci", "Picasso", "Michelangelo"}, Answer: 2, Category: "Seni", Difficulty: "Sulit"},
    }

    utils.RunQuiz(questions, playerName)
}

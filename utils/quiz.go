package utils

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
    "mini-quiz/models"
)

// Acak urutan pertanyaan
func ShuffleQuestions(questions []models.Question) []models.Question {
    rand.Seed(time.Now().UnixNano())
    shuffled := make([]models.Question, len(questions))
    perm := rand.Perm(len(questions))
    for i, v := range perm {
        shuffled[i] = questions[v]
    }
    return shuffled
}

// Hitung skor berdasarkan tingkat kesulitan
func getPoints(difficulty string) int {
    switch difficulty {
    case "Mudah":
        return 1
    case "Sedang":
        return 2
    case "Sulit":
        return 3
    default:
        return 1
    }
}

// Baca high score dari file
func ReadHighScore() int {
    data, err := os.ReadFile("highscore.txt")
    if err != nil {
        return 0
    }
    score, _ := strconv.Atoi(strings.TrimSpace(string(data)))
    return score
}

// Simpan high score ke file
func SaveHighScore(score int) {
    _ = os.WriteFile("highscore.txt", []byte(fmt.Sprintf("%d", score)), 0644)
}

// Tambahkan riwayat skor
func AppendScoreHistory(player string, score int) {
    f, err := os.OpenFile("score_history.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Gagal menyimpan riwayat skor:", err)
        return
    }
    defer f.Close()

    timestamp := time.Now().Format("2006-01-02 15:04:05")
    entry := fmt.Sprintf("%s | %s | Skor: %d\n", timestamp, player, score)
    f.WriteString(entry)
}

// Timer untuk membaca input dengan batas waktu
func readInputWithTimeout(reader *bufio.Reader, timeout time.Duration) (string, bool) {
    inputCh := make(chan string)
    go func() {
        text, _ := reader.ReadString('\n')
        inputCh <- strings.TrimSpace(text)
    }()

    select {
    case <-time.After(timeout):
        return "", false
    case input := <-inputCh:
        return input, true
    }
}

func RunQuiz(questions []models.Question, playerName string) {
    reader := bufio.NewReader(os.Stdin)
    questions = ShuffleQuestions(questions)
    score := 0

    for i, q := range questions {
        fmt.Printf("\n[%s | %s] Pertanyaan %d: %s\n", q.Category, q.Difficulty, i+1, q.Text)
        for idx, choice := range q.Choices {
            fmt.Printf("%d) %s\n", idx+1, choice)
        }
        fmt.Printf("Jawaban Anda (waktu 10 detik): ")

        input, ok := readInputWithTimeout(reader, 10*time.Second)
        if !ok {
            fmt.Println("\n⏰ Waktu habis! Pertanyaan dilewati.")
            continue
        }

        answer, err := strconv.Atoi(input)
        if err != nil || answer < 1 || answer > len(q.Choices) {
            fmt.Println("⚠️ Input tidak valid. Lewati pertanyaan.")
            continue
        }

        if answer == q.Answer {
            points := getPoints(q.Difficulty)
            fmt.Printf("✅ Benar! +%d poin\n", points)
            score += points
        } else {
            fmt.Printf("❌ Salah. Jawaban benar: %s\n", q.Choices[q.Answer-1])
        }
    }

    fmt.Printf("\n🎉 Skor Akhir Anda: %d\n", score)

    highScore := ReadHighScore()
    if score > highScore {
        fmt.Println("🏆 Selamat! Anda membuat high score baru!")
        SaveHighScore(score)
    } else {
        fmt.Printf("📈 High score saat ini: %d\n", highScore)
    }

    AppendScoreHistory(playerName, score)
    fmt.Println("📝 Riwayat skor tersimpan di score_history.txt")
}

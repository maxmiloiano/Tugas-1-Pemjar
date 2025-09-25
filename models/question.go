package models

type Question struct {
    Text       string   // Pertanyaan
    Choices    []string // Pilihan jawaban
    Answer     int      // Indeks jawaban benar (mulai dari 1)
    Category   string   // Kategori soal
    Difficulty string   // Tingkat kesulitan: Mudah, Sedang, Sulit
}

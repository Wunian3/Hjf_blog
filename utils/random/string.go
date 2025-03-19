package random

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//func TestApproach1(t *testing.T) {
//	rand.Seed(time.Now().UnixNano())
//	fmt.Println(randStr(10))
//}
//func BenchmarkApproach1(b *testing.B) {
//	rand.Seed(time.Now().UnixNano())
//	for i := 0; i < b.N; i++ {
//		_ = randStr(10)
//	}
//}

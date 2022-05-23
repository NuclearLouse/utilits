package utilits

import (
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
)

// Round ...функция округления
func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	return rounder / pow
}

// Cut ...обрезает текст до заданного числа знаков
func Cut(text string, limit int) string {
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit])
	}
	return text
}

//GetIntersection нахождение пересечений 2-х срезов для разных mode:
// 1 - те что есть только в левом массиве
// 2 - те что только в правом
// 3 - пересечение
// 4 - это те что в пересечение не попали
func GetIntersection(a []string, b []string, mode byte) []string {
	m := make(map[string]byte)

	for _, k := range a {
		m[k] += 1
	}

	for _, k := range b {
		m[k] += 2
	}

	result := []string{}

	if mode == 4 {
		for k, v := range m {
			if v < 3 {
				result = append(result, k)
			}
		}
	} else {
		for k, v := range m {
			if v == mode {
				result = append(result, k)
			}
		}
	}

	return result
}

//RemoveByIndexes вернет массив без укзанных индексов
func RemoveByIndexes(s []string, indexes []int) []string {
	var r []string
	for _, i := range indexes {
		r = append(r, s[i])
	}
	return GetIntersection(s, r, 4)
}

//OnlySpecifiedIndexes вернет массив только указанных индексов
func OnlySpecifiedIndexes(s []string, indexes []int) []string {
	var r []string
	for _, i := range indexes {
		r = append(r, s[i])
	}
	return GetIntersection(s, r, 3)
}

//MonthRus вывод месяца на русском
func MonthRus(m time.Month) string {
	switch m {
	case 1:
		return "января"
	case 2:
		return "февраля"
	case 3:
		return "марта"
	case 4:
		return "апреля"
	case 5:
		return "мая"
	case 6:
		return "июня"
	case 7:
		return "июля"
	case 8:
		return "августа"
	case 9:
		return "сентября"
	case 10:
		return "октября"
	case 11:
		return "ноября"
	default:
		return "декабря"
	}
}

//CompareHashAndPassword сравнение хэш-строки и пароля
func CompareHashAndPassword(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

// GenerateHash ...генерация хэш из пароля-строки
func GenerateHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 8)
}

//GenerateUUID генерирует UUID вида ca901b15-4089-4a96-a544-7275a7bbf3a1
func GenerateUUID() string {
	return uuid.New().String()
}

//GenerateShortUUID если lenght==nil генерирует короткий UUID вида WRxcWgJAnJkJhr9fP3FtRP либо обрезает до указанного числа знаков
func GenerateShortUUID(lenght ...int) string {
	if lenght == nil {
		return shortuuid.New()
	}
	return shortuuid.New()[:lenght[0]]
}

//RandomPass случайный пароль-строка указанной длины, если passwordLength=0 сгенерирует пароль из 10 знаков
func RandomPass(passwordLength int) string {
	if passwordLength <= 0 {
		passwordLength = 10
	}
	var (
		lowerCharSet   = "abcdedfghijklmnopqrstuvwxyz"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*()-_+={}[]~/?,.;:^"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	)
	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1

	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

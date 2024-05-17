package data

import (
	"fmt"
	"math/rand"
	"strings"
)

func Generate(dataType string) any {
	switch dataType {
	case TYPE_NAME:
		return generateName()
	case TYPE_DATE:
		return generateDate()
	case TYPE_ADDRESS:
		return generateAddress()
	case TYPE_PHONE:
		return generatePhone()
	}
	return ""
}

func generateName() string {
	nameLen := len(name)
	index := rand.Intn(nameLen)
	return name[index]
}

func generateDate() string {
	year := 1950 + rand.Intn(100)
	month := 1 + rand.Intn(12)
	day := 1 + rand.Intn(28)

	return fmt.Sprintf("%02d-%02d-%d", day, month, year)
}

func generateAddress() string {
	streetLen := len(address[SUBTYPE_STREET])
	cityLen := len(address[SUBTYPE_CITY])

	streetIndex := rand.Intn(streetLen)
	cityIndex := rand.Intn(cityLen)
	number := rand.Intn(100)

	return fmt.Sprintf("Jl. %s No. %d, %s", address[SUBTYPE_STREET][streetIndex], number, address[SUBTYPE_CITY][cityIndex])
}

func generatePhone() string {
	prefixLen := 6 + rand.Intn(4)

	var sb strings.Builder
	sb.WriteString("081")

	for i := 0; i < prefixLen; i++ {
		digit := rand.Intn(10)
		digitString := fmt.Sprintf("%d", digit)

		sb.WriteString(digitString)
	}

	result := sb.String()

	return result
}

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 編號 (ID)：正整數 (>0)
var AutoIncrementId = 1

type Individual struct {
	id     int
	gender Gender
	age    int
	intro  [200]byte
	habits []Habit
	coord  Coord
}

func NewIndividual(gender Gender, age int, intro string, habitStrings string, x, y int) (*Individual, error) {
	if gender != Male && gender != Female {
		return nil, fmt.Errorf("[性別] 只能是 MALE 或 FEMALE")
	}

	if age < 18 {
		return nil, fmt.Errorf("[年齡] 小於 18 歲")
	}

	if len(intro) > 200 {
		return nil, fmt.Errorf("[自我介紹] 超過 200 字上限")
	}

	var introArr [200]byte
	copy(introArr[:], intro)

	habits := strings.Split(habitStrings, ",")
	var habitList []Habit
	for _, habit := range habits {
		if len(habit) > 10 {
			return nil, fmt.Errorf("[興趣] 超過 10 字上限")
		}

		var habitArr [10]byte
		copy(habitArr[:], habit)
		habitList = append(habitList, Habit{name: habitArr})
	}

	currentId := AutoIncrementId
	AutoIncrementId++

	return &Individual{currentId, gender, age, introArr, habitList, Coord{x, y}}, nil
}

func (i *Individual) GetId() int {
	return i.id
}

func (i *Individual) GetCoord() *Coord {
	return &i.coord
}

func (i *Individual) GetHabits() []Habit {
	return i.habits
}

func generateFakeIndividual(num int) []Individual {
	var users []Individual
	randomString := func(length int) string {
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
		rand.Seed(time.Now().UnixNano())
		builder := strings.Builder{}
		for i := 0; i < length; i++ {
			builder.WriteByte(charset[rand.Intn(len(charset))])
		}
		return builder.String()
	}
	for i := 0; i < num; i++ {
		user, err := NewIndividual(Gender(rand.Intn(2)), rand.Intn(100)+18, randomString(200), randomString(10)+","+randomString(10), rand.Intn(100), rand.Intn(100))
		if err != nil {
			panic(err)
		}
		users = append(users, *user)
	}
	return users
}

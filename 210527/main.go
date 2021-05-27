package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Score struct {
	Korean  uint `json:"korean,omitempty"`
	Math    uint `json:"math,omitempty"`
	English uint `json:"english,omitempty"`
}

type UserV1 struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint64 `json:"age"`
	Score    Score  `json:"score,omitempty"`
}

type UserV2 struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint64 `json:"age"`
	Score    *Score `json:"score,omitempty"`
}

var (
	userV1MockupScore = UserV1{
		UserName: "alice",
		Name:     "Alice",
		Email:    "alice@alice.com",
		Age:      28,
		Score: Score{
			Korean:  88,
			Math:    78,
			English: 98,
		},
	}

	userV1MockupNonScore = UserV1{
		UserName: "cathy",
		Name:     "Cathy",
		Email:    "cathy@cathy.com",
		Age:      26,
	}

	userV2MockUpScore = UserV2{
		UserName: "bob",
		Name:     "Bob",
		Email:    "bob@bob.com",
		Age:      38,
		Score: &Score{
			Korean:  10,
			Math:    20,
			English: 30,
		},
	}

	userV2MockUpNonScore = UserV2{
		UserName: "david",
		Name:     "David",
		Email:    "david@david.com",
		Age:      32,
	}
)

func main() {
	e := echo.New()

	e.GET("/api/v1/user/:username", GetUser)

	e.Logger.Fatal(e.Start(":30000"))
}

func GetUser(c echo.Context) error {
	username := c.Param("username")
	var userdata interface{}
	switch username {
	case "alice":
		userdata = userV1MockupScore
	case "bob":
		userdata = userV2MockUpScore
	case "cathy":
		userdata = userV1MockupNonScore
	case "david":
		userdata = userV2MockUpNonScore
	}
	return c.JSONPretty(http.StatusOK, userdata, "  ")
}

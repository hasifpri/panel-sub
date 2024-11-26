package infrastructureutilities

import (
	"fmt"
	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/bcrypt"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
	"strconv"
)

func EncodeID(id int64) string {
	ids := strconv.FormatInt(id, 10)
	var arrs = []string{ids}
	var arri = []int{}

	for _, i := range arrs {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		}
		arri = append(arri, j)
	}

	hd := hashids.NewData()
	hd.Salt = infrastructureconfiguration.HashSalt
	h, _ := hashids.NewWithData(hd)
	hashid, _ := h.Encode(arri)

	return hashid
}

func DecodeID(hashid string) int64 {
	hd := hashids.NewData()
	hd.Salt = infrastructureconfiguration.HashSalt
	h, _ := hashids.NewWithData(hd)
	arri, _ := h.DecodeWithError(hashid)

	ids := ""
	for _, i := range arri {
		s := strconv.Itoa(i)
		ids = ids + s
	}
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	return id
}

func HashPassword(naked string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(naked), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

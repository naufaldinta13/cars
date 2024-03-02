package rest

import (
	"strconv"
	"testing"

	"github.com/env-io/orm"
	"github.com/naufaldinta13/cars/entity"
)

func TestCreateFailedRequest(t *testing.T) {
	req := &createRequest{}

	v := req.Validate()

	if !v.Valid {
		t.Errorf("SALAH! harusnya error")
	}
}

func TestCreateSuccessRequest(t *testing.T) {
	req := &createRequest{Name: "Avanza G.E", MonthRate: 20000, Image: "https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Google_Images_2015_logo.svg/1200px-Google_Images_2015_logo.svg.png"}

	v := req.Validate()

	if v.Valid {
		t.Errorf("SALAH! harusnya tidak error")
	}
}

func TestDeleteFailedRequest(t *testing.T) {
	req := &deleteRequest{}

	v := req.Validate()

	if !v.Valid {
		t.Errorf("SALAH! harusnya error")
	}
}

func TestDeleteSuccessRequest(t *testing.T) {
	mx := &entity.Cars{
		CarName:   "Vios",
		MonthRate: 25000,
		Image:     "https://upload.wikimedia.org/wikipedia/commons/thumb/7/77/Google_Images_2015_logo.svg/1200px-Google_Images_2015_logo.svg.png",
	}

	oid, _ := orm.NewOrm().Insert(mx)

	req := &deleteRequest{ID: strconv.Itoa(int(oid))}

	v := req.Validate()

	if v.Valid {
		t.Errorf("SALAH! harusnya tidak error")
	}
}

package repository

import (
	"context"
	"course/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	getUsersURL = "/internal/users"
)

type McsUserRepo struct {
	HostName string
	Username string
	Password string
	Client   *http.Client
}

func NewMcsUserRepo() *McsUserRepo {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	return &McsUserRepo{
		HostName: "http://localhost:8080",
		Username: "user",
		Password: "abcd1234",
		Client:   &client,
	}
}

func (mcsUserRepo McsUserRepo) GetByID(ctx context.Context, ID int) (domain.User, error) {
	//http://localhost:8080/internal/users:id
	url := fmt.Sprintf("%s%s/%d", mcsUserRepo.HostName, getUsersURL, ID)
	log.Println("url:", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	var user domain.User
	if err != nil {
		return user, err
	}
	req.SetBasicAuth(mcsUserRepo.Username, mcsUserRepo.Password)
	log.Println("hit to:", req.Method, req.URL)
	//hit endpoint
	resp, err := mcsUserRepo.Client.Do(req)
	if err != nil {
		log.Println(err)
		return user, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return user, err
	}
	err = json.NewDecoder(resp.Body).Decode(&user)
	return user, err
}

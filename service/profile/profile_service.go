package profile

import (
	"fmt"
	"math/rand"
	"template/entity"
	"template/repository/profile"
	"time"
)

type profileService struct {
	repository profile.ProfileRepository
}

func NewProfileService(repository profile.ProfileRepository) ProfileService {
	return &profileService{
		repository: repository,
	}
}

func (p profileService) GetProfile(profileID string) (*entity.Profile, error) {
	profile, err := p.repository.Get(profileID)
	if err != nil {
		return nil, err
	}

	return profile, nil

}

func (p profileService) CreateProfile(profile *entity.Profile) (string, error) {

	profile.ProfileCode = GenerateProfileID()

	err := p.repository.Post(profile)
	if err != nil {
		return "", err
	}

	return profile.ProfileCode, nil
}

func (p profileService) UpdateProfile(profileID string, payload *entity.Profile) error {
	err := p.repository.Update(profileID, payload)
	if err != nil {
		return err
	}

	return nil
}

func GenerateProfileID() string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10000)

	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d%d", timestamp, randomNumber)

	referenceNumber := input

	return referenceNumber
}

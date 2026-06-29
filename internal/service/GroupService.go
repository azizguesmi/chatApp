package service

import (
	"fmt"

	"backend/internal/model"
	"backend/internal/repo"
)

func AddGroup(name string, creatorID int, memebers []int) (int64, error) {
	group := model.Group{
		Name:      name,
		CreatorID: creatorID,
		Members:   memebers,
	}
	id, err := repo.AddGroup(group)
	if err != nil {
		return 0, fmt.Errorf("error in adding group (servie) %w", err)
	}
	_, err = repo.AddMembersToGroup(memebers, int(id))
	if err != nil {
		return 0, fmt.Errorf("error in adding group memebers in (service) %w", err)
	}
	return id, nil
}

func DeleteGroup(groupID int) (bool, error) {
	test, err := repo.DeleteGroup(groupID)
	if err != nil {
		return false, fmt.Errorf("error while deleting group  %w", err)
	}
	return test, nil
}

func GetGroupByID(groupID int) (*model.Group, error) {
	group, err := repo.GetGroupByID(groupID)
	if err != nil {
		return nil, fmt.Errorf("error while getting group in (servie) %w", err)
	}
	return group, nil
}

func GetGroupMembers(groupID int) ([]int, error) {
	m, err := repo.GetMembersOfGroup(groupID)
	if err != nil {
		return nil, fmt.Errorf("error while getting memebrs (service) %w", err)
	}
	return m, nil
}

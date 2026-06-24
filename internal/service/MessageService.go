package service

import (
	"backend/internal/model"
	"backend/internal/repo"
	"fmt"
)

func AddMessage(content string, senderId, receiverId int, Rec_type string) (int64, error) {
	mess := model.Message{
		Content:    content,
		SenderID:   senderId,
		ReceiverID: receiverId,
		Rec_type:   Rec_type,
	}

	id, err := repo.AddMessage(mess)
	if err != nil {
		return 0, fmt.Errorf("error adding message (service) %W", err)
	}
	return id, nil
}

func RemoveMessage(id int) (bool, error) {
	test, err := repo.DeleteMessage(id)
	if err != nil {
		return false, fmt.Errorf("error deleting message %w", err)
	}
	return test, nil
}

func GetMessageById(id int) (*model.Message, error) {
	message, err := repo.GetMessageById(id)

	if err != nil {
		return nil, fmt.Errorf("error while getting message %W", err)
	}
	return message, nil
}

func GetAllMessage() ([]model.Message, error) {
	messages, err := repo.GetAllMessages()
	if err != nil {
		return nil, fmt.Errorf("error in getting all messages in service %w", err)
	}
	return messages, nil
}

func GetMessageBySender(senderId int) ([]model.Message, error) {
	messages, err := GetAllMessage()
	if err != nil {
		return nil, err
	}
	var finalRes []model.Message
	for _, m := range messages {
		if m.SenderID == senderId {
			finalRes = append(finalRes, m)
		}
	}
	return finalRes, nil
}

func GetMessageReceivedByAUser(id int) ([]model.Message, error) {
	mess, err := repo.GetAllMessagesReceivedBy(id, "USER")
	if err != nil {
		return nil, fmt.Errorf("error while getting messages in service %w", err)
	}
	return mess, nil
}

func GetMessageReceivedByAGroup(id int) ([]model.Message, error) {
	mess, err := repo.GetAllMessagesReceivedBy(id, "GROUP")
	if err != nil {
		return nil, fmt.Errorf("error while getting messages in service %w", err)
	}
	return mess, nil
}

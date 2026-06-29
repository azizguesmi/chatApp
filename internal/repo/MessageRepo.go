package repo

import (
	db "backend/internal/database"
	"backend/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func AddMessage(message model.Message) (int64, error) {
	conn, err := db.Connect()

	if err != nil {
		return 0, fmt.Errorf("adding message connecion error %w", err)
	}

	defer conn.Close()
	var res sql.Result
	if message.Rec_type == "USER" {
		res, err = conn.Conn.Exec(
			"INSERT INTO messages (sender_id,receiver_id_user,content,created_at,rec_type, receiver_id_group) VALUES (?,?,?,?,?,?)",
			message.SenderID,
			message.ReceiverID,
			message.Content,
			message.CreatedAt,
			message.Rec_type,
			nil,
		)
	} else if message.Rec_type == "GROUP" {
		res, err = conn.Conn.Exec(
			"INSERT INTO messages (sender_id,receiver_id_user,content,created_at,rec_type, receiver_id_group) VALUES (?,?,?,?,?,?)",
			message.SenderID,
			nil,
			message.Content,
			message.CreatedAt,
			message.Rec_type,
			message.ReceiverID,
		)
	} else {
		return 0, fmt.Errorf("invalid rec_type %s", message.Rec_type)
	}
	if err != nil {
		return 0, fmt.Errorf("error while exection of query in adding message %w", err)
	}

	lastid, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error geting last inseted id %w", err)
	}

	return lastid, nil
}

func DeleteMessage(id int) (bool, error) {
	conn, err := db.Connect()
	if err != nil {
		return false, fmt.Errorf("error while connection to db in deleting message %w", err)
	}
	defer conn.Close()
	res, err := conn.Conn.Exec(
		"DELETE FROM messages where id=?",
		id,
	)
	if err != nil {
		return false, fmt.Errorf("error while exec delete message %w", err)
	}
	rowEffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error while checking the number of affectied rows %w", err)
	}

	return rowEffected != 0, nil
}

func GetAllMessagesReceivedBy(id int, t string) ([]model.Message, error) {
	//t is the type of the receiver (GROUP | USER)
	conn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error while connecting to db in getAllMessagesReceivedByUser %w", err)
	}
	defer conn.Close()
	var rows *sql.Rows
	if t == "USER" {
		rows, err = conn.Conn.Query(
			"SELECT id,sender_id,receiver_id,content,created_at,rec_type FROM messages WHERE receiver_id_user=? and rec_type=?",
			id,
			t,
		)
	} else if t == "GROUP" {
		rows, err = conn.Conn.Query(
			"SELECT id,sender_id,receiver_id,content,created_at,rec_type FROM messages WHERE receiver_id_group=? and rec_type=?",
			id,
			t,
		)
	} else {
		return nil, fmt.Errorf("wrong t it should be (USER | GROUP")
	}
	var ms []model.Message
	for rows.Next() {
		var m model.Message

		err = rows.Scan(
			&m.ID,
			&m.SenderID,
			&m.ReceiverID,
			&m.Content,
			&m.CreatedAt,
			&m.Rec_type,
		)

		if err != nil {
			log.Println("error while scanning line in gettingmessageByReceiver %w", err)
			continue
		}
		ms = append(ms, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func GetMessageById(id int) (*model.Message, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error connection to db while connection to db %w", err)
	}
	defer conn.Close()
	var m model.Message
	err = conn.Conn.QueryRow(
		"SELECT id,sender_id,receiver_id_user,content,created_at,rec_type,receiver_id_group FROM messages WHERE id=?",
		id,
	).Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Content, &m.CreatedAt, &m.Rec_type, &m.ReceiverID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error in data base reading %W", err)
	}
	return &m, nil
}

func GetAllMessages() ([]model.Message, error) {
	fn := "GetAllMessages"
	conn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error connection to db while connection to db in %s %w", fn, err)
	}
	defer conn.Close()
	var ms []model.Message
	rows, err := conn.Conn.Query(
		"SELECT id, sender_id, receiver_id_user, content, created_at, rec_type, receiver_id_group FROM messages",
	)
	if err != nil {
		return nil, fmt.Errorf("error while executing query in %s %W", fn, err)
	}
	for rows.Next() {
		var m model.Message
		var receiverIdGroup *int
		var receiverIdUser *int
		err = rows.Scan(
			&m.ID,
			&m.SenderID,
			receiverIdUser,
			&m.Content,
			&m.CreatedAt,
			&m.Rec_type,
			receiverIdGroup,
		)
		if err != nil {
			log.Println("error while scanning line in gettingmessageByReceiver %w", err)
			continue
		}
		if m.Rec_type == "USER" {
			m.ReceiverID = *receiverIdUser
		} else if m.Rec_type == "GROUP" {
			m.ReceiverID = *receiverIdGroup
		} else {
			log.Fatal("error in receiver type")
			continue
		}
		ms = append(ms, m)
	}
	return ms, nil
}

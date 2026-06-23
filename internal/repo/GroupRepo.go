package repo

import (
	db "backend/internal/database"
	"backend/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetGroupById(id int) (*model.Group, error) {
	fn := "GetGroupById"
	conn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error while trying to connect to db in %s %w", fn, err)
	}
	defer conn.Close()
	var group model.Group
	err = conn.Conn.QueryRow(
		"SELECT id,name,creator,created_at FROM groups WHERE id=?",
		id,
	).Scan(&group.ID, &group.Name, &group.CreatorID, &group.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("group not found")
		}
		return nil, fmt.Errorf("error in exec query in %s %w", fn, err)
	}

	m, err := GetMembersOfGroup(id)
	if err != nil {
		return nil, fmt.Errorf("error in getting memebrs of group %s %w", fn, err)
	}
	group.Members = m
	return &group, nil
}

func GetMembersOfGroup(id int) ([]int, error) {
	fn := "GetMemberOfGroup"
	conn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("error while trying to connect to db in %s %w", fn, err)
	}
	var ids []int
	rows, err := conn.Conn.Query(
		"SELECT member_id FROM group_members WHERE group_id=?",
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("error while exec query in %s %w", fn, err)
	}

	for rows.Next() {
		var i int
		err = rows.Scan(&i)
		if err != nil {
			log.Println("error while scanning a line in %s %w", fn, err)
			continue
		}
		ids = append(ids, i)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

func AddGroup(group model.Group) (int64, error) {
	fn := "AddGroup"
	conn, err := db.Connect()
	if err != nil {
		return 0, fmt.Errorf("error while trying to connect to db in %s %w", fn, err)
	}
	defer conn.Close()
	res, err := conn.Conn.Exec(
		"INSERT INTO groups (name,creator,created_at) VALUES (?,?,?,?)",
		group.Name,
		group.CreatorID,
		group.CreatedAt,
	)
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error while trying to get last inserted row in %s %w ", fn, err)
	}

	return lastId, nil
}

func AddMembersToGroup(members []int, groupId int) (int64, error) {
	fn := "AddMembersToGroup"

	conn, err := db.Connect()
	if err != nil {
		return 0, fmt.Errorf("error while connecting to db in %s: %w", fn, err)
	}
	defer conn.Close()

	var totalInserted int64

	for _, id := range members {
		res, err := conn.Conn.Exec(
			`INSERT OR IGNORE INTO group_members (member_id, group_id)
			 VALUES (?, ?)`,
			id,
			groupId,
		)
		if err != nil {
			return totalInserted, fmt.Errorf("error inserting member %d in %s: %w", id, fn, err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return totalInserted, fmt.Errorf("error getting rows affected in %s: %w", fn, err)
		}

		totalInserted += rowsAffected
	}

	return totalInserted, nil
}

func DeleteMemberOfAGroup(id int, groupId int) (bool, error) {
	fn := "DeleteMemeberOfGroup"
	conn, err := db.Connect()
	if err != nil {
		return false, fmt.Errorf("error while connecting to db in %s: %w", fn, err)
	}
	defer conn.Close()
	res, err := conn.Conn.Exec(
		"DELETE FROM group_members WHERE member_id=? AND group_id=?",
		id,
		groupId,
	)
	if err != nil {
		return false, fmt.Errorf("error while deleting memebr  in %s %w", fn, err)
	}
	rowEffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error while checking the number of affectied rows in %s err : %w", fn, err)
	}
	return rowEffected != 0, nil
}

func DeleteGroup(id int) (bool, error) {
	fn := "DeleteGroup"
	conn, err := db.Connect()
	if err != nil {
		return false, fmt.Errorf("error while connecting to db in %s: %w", fn, err)
	}
	defer conn.Close()
	res, err := conn.Conn.Exec(
		"DELETE FROM groups WHERE id=?",
		id,
	)
	if err != nil {
		return false, fmt.Errorf("error while deleting memebr  in %s %w", fn, err)
	}
	rowEffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("error while checking the number of affectied rows in %s err : %w", fn, err)
	}
	return rowEffected != 0, nil
}

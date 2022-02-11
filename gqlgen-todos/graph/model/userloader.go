package model

import (
	"time"
)

const loadersKey = "dataloaders"

type Loaders struct {
	UserById UserLoader
}

var UserByIdLoader = &UserLoader{
	maxBatch: 100,
	wait:     1 * time.Millisecond,
	fetch: func(ids []string) ([]*User, []error) {
		var users []*User
		//重要，需要保证返回条数一致
		userRets := make([]*User, len(ids))
		sqlDB.Table("user").Where("id IN ?", ids).Scan(&users)
		userById := map[string]*User{}
		for _, userItem := range users {
			userById[userItem.ID] = userItem
		}
		//fmt.Println(userById)
		for i, id := range ids {
			value, flag := userById[id]
			if flag {
				//fmt.Println(i)
				userRets[i] = value
			} else {
				//fmt.Println(i, "not ok")
				userRets[i] = &User{}
			}
		}
		return userRets, nil
	},
}

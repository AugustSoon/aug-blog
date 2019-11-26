package service

import (
	"github.com/JumpSama/aug-blog/model"
	"sync"
)

func UserList(req *model.ListRequest) ([]*model.UserInfo, int) {
	list, count := model.GetUserList(req)

	ids := []uint{}

	for _, user := range list {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(list)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, u := range list {
		wg.Add(1)

		go func(u *model.User) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.ID] = &model.UserInfo{
				Id:        u.ID,
				Account:   u.Account,
				Username:  u.Username,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case <-errChan:
		return nil, count
	}

	lists := make([]*model.UserInfo, 0)

	for _, id := range ids {
		lists = append(lists, userList.IdMap[id])
	}

	return lists, count
}

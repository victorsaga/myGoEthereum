package RedisDistributedLockHelper

import (
	"errors"
	"myGoEthereum/Helper/RedisHelper"
	"time"

	"github.com/google/uuid"
)

const lockPrefix = "RedisDistributedLock_"

type RedisDistributedLock struct {
	lockKey string
	guid    string
}

func (obj RedisDistributedLock) Release() {
	v, _ := RedisHelper.GetString(obj.lockKey)

	if v != nil && *v == obj.guid {
		RedisHelper.DeleteKey(obj.lockKey)
	}
}

// lockKey 鎖定的 Key
// lockExpirySeconds 鎖定過多久後會自動失效(Redis TTL)，單位為秒
// retryWaitSeconds 嘗試鎖定失敗後，過多久後重新嘗試鎖定，單位為秒
// tryGetLockSeconds 要嘗試取得鎖定多長時間，單位為秒
func GetLock(lockKey string, lockExpirySeconds int64, retryWaitSeconds int64, tryGetLockSeconds int64) (lock RedisDistributedLock, err error) {
	startTime := time.Now()

	lock.lockKey = lockPrefix + lockKey
	lock.guid = uuid.New().String()
	continueTry := true
	isFristTimeTry := true

	for continueTry = true; continueTry; continueTry = (int64(time.Since(startTime)/time.Second) < tryGetLockSeconds) {
		//如果不是第一次嘗試
		if !isFristTimeTry {
			time.Sleep(time.Duration(retryWaitSeconds) * time.Second)
		}
		isFristTimeTry = false

		//取得鎖定
		v, e := RedisHelper.GetString(lock.lockKey)

		if e != nil {
			err = e
			return
		}

		if v != nil {
			//取得鎖定失敗
			continue
		} else {
			//取得鎖定成功
			RedisHelper.SetStringWithExpire(lock.lockKey, lock.guid, lockExpirySeconds)

			//確定鎖定的KEY是自己設定的
			v, e := RedisHelper.GetString(lock.lockKey)

			if e != nil {
				err = e
				lock.Release()
				return
			}

			//如果key不是自己設定的，則重新開始迴圈
			if *v != lock.guid {
				continue
			}

			break
		}
	}

	//如果超過嘗試時間
	if continueTry == false {
		err = errors.New("Failed Get Lock.")
	}

	return
}

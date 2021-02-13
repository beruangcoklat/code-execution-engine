package common

import "time"

const (
	CPPType = 1
)

const (
	REQUEUE_DELAY          = time.Duration(1 * time.Second)
	TOPIC_EXECUTE_CODE     = "execute_code"
	CHANNEL_EXECUTE_CODE   = "execute_code"
	TOPIC_EXECUTION_RESULT = "execution_result"
)

package tox

/*
#include <tox/tox.h>
#include <tox/toxav.h>
*/
import "C"

const (
	USER_STATUS_NONE = int(C.TOX_USER_STATUS_NONE)
	USER_STATUS_AWAY = int(C.TOX_USER_STATUS_AWAY)
	USER_STATUS_BUSY = int(C.TOX_USER_STATUS_BUSY)
)

const (
	FILE_CONTROL_RESUME = int(C.TOX_FILE_CONTROL_RESUME)
	FILE_CONTROL_PAUSE  = int(C.TOX_FILE_CONTROL_PAUSE)
	FILE_CONTROL_CANCEL = int(C.TOX_FILE_CONTROL_CANCEL)
	FILE_KIND_DATA      = int(C.TOX_FILE_KIND_DATA)
	FILE_KIND_AVATAR    = int(C.TOX_FILE_KIND_AVATAR)
)

const (
	GROUPCHAT_TYPE_TEXT = int(C.TOX_GROUPCHAT_TYPE_TEXT)
	GROUPCHAT_TYPE_AV   = int(C.TOX_GROUPCHAT_TYPE_AV)
)

const (
	FRIEND_CALL_STATE_ERROR       = int(C.TOXAV_FRIEND_CALL_STATE_ERROR)
	FRIEND_CALL_STATE_FINISHED    = int(C.TOXAV_FRIEND_CALL_STATE_FINISHED)
	FRIEND_CALL_STATE_SENDING_A   = int(C.TOXAV_FRIEND_CALL_STATE_SENDING_A)
	FRIEND_CALL_STATE_SENDING_V   = int(C.TOXAV_FRIEND_CALL_STATE_SENDING_V)
	FRIEND_CALL_STATE_ACCEPTING_A = int(C.TOXAV_FRIEND_CALL_STATE_ACCEPTING_A)
	FRIEND_CALL_STATE_ACCEPTING_V = int(C.TOXAV_FRIEND_CALL_STATE_ACCEPTING_V)
)

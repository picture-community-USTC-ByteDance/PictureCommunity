package initialize

import (
	"picture_community/global"
	"picture_community/utils"
)

func IdGeneratorInitialize() {
	global.CommentIDGenerator = utils.NewIDGenerator(0, 3)
	global.PostIDGenerator = utils.NewIDGenerator(0, 2)
	global.ForwardIDGenerator = utils.NewIDGenerator(0, 1)
	global.UserIDGenerator = utils.NewIDGenerator(0, 0)
}

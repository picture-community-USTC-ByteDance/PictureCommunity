package initialize

import (
	"picture_community/global"
	"picture_community/utils"
)

func IdGeneratorInitialize() {
	global.CommentIDGenerator = utils.NewIDGenerator(3)
	global.PostIDGenerator = utils.NewIDGenerator(2)
	global.ForwardIDGenerator = utils.NewIDGenerator(1)
	global.UserIDGenerator = utils.NewIDGenerator(0)
	global.PostPhotoIDGenerator = utils.NewIDGenerator(4)
}

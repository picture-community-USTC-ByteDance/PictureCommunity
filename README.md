# Picture-Community
This is a community similar to Instagram and Weibo. Using Golang and React to develop.

## 服务器代码更新攻略

1. ssh登录服务器： 

    `ssh root@121.5.1.73` 

    (密码pic123)
    
2. 切换目录

    `cd /usr/local/go/src/PictureCommunity/`

3. 获取最新代码

    `git pull`

4. 查看正在运行的老程序的pid

    `lsof -i:8080`

5. 杀死旧程序

    `kill -9 [pid]`

6. 后台运行更新后的程序

    `go run . &`


注意：不要删除代码然后重新下载新的，否则之前上传的图片也没了
config
======
配置文件放入到conf/ 目录下


配置文件


import (
    "github.com/yaotian/config"
)

config.Config.Get("site.title").(string)

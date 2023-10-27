# project-layout
### Golang 项目的代码组织结构

```bash
.
├── LICENSE         协议文档
├── Makefile        Makefile，执行调用scripts里的脚本
├── README.md       README
├── api             可对外提供的 API 文件
├── assets          项目用到的资源文件，比如icon，logo之类的
├── build           用于打包的CI/CD的文件
├── cmd             存放可执行的文件
├── configs         存放配置文件
├── deploy          k8s用到的文件
├── examples        示例代码
├── githooks        git钩子
├── go.mod          依赖包说明文件
├── init            系统初始化和进程管理的配置
├── internal        私有的代码库（别的项目不能导入使用）
├── main.go         程序入口
├── pkg             公有的代码库（别的项目可以导入使用）
├── scripts         执行构建的脚本，Makefile调用这个目录下的shell脚本
├── test            测试程序和测试数据
├── third_party     第三方辅助工具
├── tools           当前项目用到的工具
├── web             静态web资源，前端代码
└── website         项目的网站数据
```

### 说明

    仅供参考，并不是每个目录都是必须，自行决定

### v1 使用的库
* CLI 命令(spf13/cobra): https://github.com/spf13/cobra
* 配置读取(spf13/viper): https://github.com/spf13/viper
* SQL(jmoiron/sqlx): https://github.com/jmoiron/sqlx
* Log(uber-go/zap): https://github.com/uber-go/zap
* Lumberjack(natefinch/lumberjack.v2): https://github.com/natefinch/lumberjack
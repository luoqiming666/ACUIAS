`main.go` 是程序入口文件，用于创建Gin引擎和运行服务器。

`app` 目录包含了应用程序的主要业务逻辑。

- `models` 存放与数据库表对应的模型结构体。
- `repositories` 包含处理与数据库交互的仓库（Repository）层代码。

- `services` 包含系统的核心业务逻辑，如身份认证服务和区块链服务。
1.包含验证身份服务
2.写入区块链服务

- `controllers` 包含处理HTTP请求的控制器（Controller）层代码。
- `middleware` 包含自定义的中间件，如身份验证中间件。

`config` 目录包含应用程序的配置文件，用于读取和管理配置信息。

`utils` 目录包含应用程序的辅助工具，如日志记录器。

`database` 目录包含数据库迁移文件，用于创建数据库表结构。

`blockchain` 目录包含与联盟链交互的客户端代码。

-----------------------------------------------------
文件职责明细

- main.go：主要的入口文件，用于启动整个应用程序。
- app：应用程序的核心逻辑，可能包含应用程序的初始化、全局变量、路由配置等。
- models：定义应用程序中使用的数据模型。
- user.go：用户模型，用于表示应用程序中的用户信息。
- operation_log.go：操作日志模型，用于记录应用程序中的操作日志信息。
- repositories：用于与数据库进行交互，执行数据持久化操作。
- user_repository.go：用户仓库，提供用户数据的增删改查等数据库操作。
- operation_log_repository.go：操作日志仓库，提供操作日志数据的增删改查等数据库操作。
- services：封装应用程序的业务逻辑。
- auth_service.go：认证服务，提供用户身份验证和权限管理相关的功能。
- blockchain_service.go：区块链服务，封装与区块链相关的操作和业务逻辑。
- controllers：处理来自客户端的请求，调用相应的服务和方法，并返回响应给客户端。
- auth_controller.go：认证控制器，处理与用户认证和授权相关的请求。
- middleware：中间件，处理请求和响应之间的预处理逻辑，例如身份验证、日志记录等。
- auth_middleware.go：认证中间件，验证请求的身份和权限。
- config：应用程序的配置文件。
- config.go：包含应用程序的配置信息，如数据库连接、认证密钥等。
- utils：通用的工具函数或类库。
- logger.go：日志记录器，用于记录应用程序的日志信息。
- database：与数据库相关的文件和目录。
- migration：数据库迁移文件，用于管理数据库结构的版本控制。
- 20230715_create_tables.up.sql：数据库迁移脚本，用于创建表格的SQL语句。
- 20230715_create_tables.down.sql：数据库迁移脚本，用于删除表格的SQL语句。
- blockchain：与区块链相关的文件和目录。
- blockchain_client.go：区块链客户端，封装与区块链节点进行通信的功能。
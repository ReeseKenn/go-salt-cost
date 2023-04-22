## 运行
`go build; ./salt-cost`

数据库创建命令： `docker run -d  -p 5432:5432  --name pg-for-go-salt-cost     -e POSTGRES_USER=salt        -e POSTGRES_PASSWORD=123456      -e POSTGRES_DB=salt_cost_dev      -e PGDATA=/var/lib/postgresql/data/pgdata      -v salt-cost-data:/var/lib/postgresql/data      --network=network1       postgres:14`
google: golang connect psql
connection strings.com 连接字符串的写法
安装驱动 
	_ "github.com/lib/pq"

docker ps
docker exec -it pg-for-go-salt-cost bash
psql -U salt -d salt_cost_go_dev
\d
\d users
ctrl + d退出容器

无法从外部连接数据库 需要不加network(加了会新开一个虚拟局域网)
获取地址 wsl/ubuntu运行`ip addr show eth0 | grep -oP '(?<=inet\s)\d+(\.\d+){3}'`

go get 安装依赖

### 使用MySQL数据库
创建容器 `docker run -d -p 3306:3306 --network=network1 --name mysql-for-go-salt-cost -e MYSQL_DATABASE=salt_cost_go_dev -e MYSQL_USER=salt -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v mysql-go-salt-cost-data:/var/lib/mysql -d mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci`
不使用network1连接
`docker run -d -p 3306:3306 --name mysql-for-go-salt-cost -e MYSQL_DATABASE=salt_cost_go_dev -e MYSQL_USER=salt -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v mysql-go-salt-cost-data:/var/lib/mysql -d mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci`

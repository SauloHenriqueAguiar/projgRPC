module projgRPC

go 1.21 // ou a versão que você está utilizando

toolchain go1.23.2

require (
	google.golang.org/grpc v1.67.1 // ou a versão mais recente
	gorm.io/driver/sqlite v1.5.6 // ou a versão mais recente
	gorm.io/gorm v1.25.12 // ou a versão mais recente
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.35.1
)

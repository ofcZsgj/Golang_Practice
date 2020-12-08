module GolangPractice

go 1.15

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.34.0

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
)

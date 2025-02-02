.PHONY: seed

seed:
	 docker-compose exec app go run internal/commands/seed.go

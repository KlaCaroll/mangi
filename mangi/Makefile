.PHONY: api
api:
	go build -o bin/api ./src/api

.PHONY: db
db:
	rm bin/db.mysql
	cat src/db/schema.sql | mysql bin/db.mysql
	cat src/db/seed.sql | mysql bin/db.mysql
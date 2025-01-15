.PHONY: api
api:
	go build -o bin/api ./api

.PHONY: db
db:
	rm bin/db.mysql
	cat db/schema.sql | mysql bin/db.mysql
	cat db/seed.sql | mysql bin/db.mysql
pgm:
	docker build -t linkai_pg_migrations -f Dockerfile.pgm .

amm:
	docker build -t linkai_am_migrations -f Dockerfile.amm .
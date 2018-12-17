SERVICES = am_migrations pg_migrations db_systemids
APP_ENV = dev
systemids:
	docker build -t linkai_db_systemids -f Dockerfile.db_systemids .
	
pgm:
	docker build -t linkai_pg_migrations -f Dockerfile.pgm .

amm:
	docker build -t linkai_am_migrations -f Dockerfile.amm .

all: pgm amm systemids

pushamm:
	docker tag linkai_am_migrations:latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/am_migrations:latest

pushecr:
	$(foreach var,$(SERVICES),docker tag linkai_$(var):latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest;)
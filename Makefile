SERVICES = am_migrations pg_migrations
APP_ENV = dev

pgm:
	docker build -t linkai_pg_migrations -f Dockerfile.pgm .

amm:
	docker build -t linkai_am_migrations -f Dockerfile.amm .

pushecr:
	$(foreach var,$(SERVICES),docker tag linkai_$(var):latest 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest && docker push 447064213022.dkr.ecr.us-east-1.amazonaws.com/$(var):latest;)
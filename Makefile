.PHONY: test devup devdn dev stop 

test:
		docker compose start && go test ./ -v

devup:
		docker compose up -d

devdn:
		docker compose down

dev:
		docker compose start

stop:
		docker compose stop
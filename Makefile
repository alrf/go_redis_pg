SHELL := /bin/bash

test-post:
	for i in {1..20}; do \
		curl -X POST -H 'Content-Type: application/json' -d "{"\"department"\": "\"Dep$$i"\", "\"section"\": $$i, "\"equipment"\": "\"Sect$$i"\", "\"description"\": "\"Descr$$i"\"}" "http://127.0.0.1:8080/humai/echoservice"; \
	done

test-get:
	curl "http://127.0.0.1:8080/humai/echoservice";

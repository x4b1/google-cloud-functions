.PHONY: mod
mod:
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor

.PHONY: deploy-sum
deploy-sum:
	gcloud functions deploy Sum --runtime go113 --trigger-http

deploy-sum-async:
	gcloud functions deploy SumAsync --runtime go113 --trigger-topic sum-values

.PHONY: deploy-substract
deploy-substract:
	gcloud functions deploy Substract --runtime go113 --trigger-http

.PHONY: check-sum
check-sum:
	@curl --request POST \
  --url {{function_endpoint_url}} \
  --header 'content-type: application/json' \
  --data '{ "first": 1, "second": 1 }'

.PHONY: check-substract
check-substract:
	@curl --request POST \
	--url {{function_endpoint_url}} \
  --header 'content-type: application/json' \
  --data '{ "first": 5, "second": 3 }'

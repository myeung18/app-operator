
ORG = myeung
NAMESPACE ?= ms-test
PROJECT = app-operator
TAG ?= latest


.PHONY: build-image
build-image: compile build

.PHONY: build
build: go-generate
	operator-sdk build quay.io/${ORG}/${PROJECT}:${TAG} --go-build-args "-ldflags -X=github.com/myeung18/app-operator/pkg/controller/appservice.ProductName=ChangeAsProduct--"

.PHONY: push-image
push-image:
	docker push quay.io/myeung/app-operator:${TAG}

.PHONY: compile
compile:
	go build -o=build/_output/bin/app-operator ./cmd/manager/main.go

.PHONY: generate-csv
generate-csv:
	operator-sdk olm-catalog gen-csv --csv-version ${TAG}

.PHONY: go-generate
go-generate:
	go generate ./...

.PHONY: verify-csv
verify-csv:
	operator-courier verify --ui_validate_io deploy/olm-catalog/hawtio-operator

.PHONY: push-csv
push-csv:
	operator-courier push deploy/olm-catalog/hawtio-operator ${QUAY_NAMESPACE} hawtio-operator ${TAG} "${QUAY_TOKEN}"

.PHONY: install
install: install-crds
	kubectl apply -f deploy/service_account.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/role.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/role_binding.yaml -n ${NAMESPACE}
	kubectl apply -f deploy/cluster_role.yaml
	cat deploy/cluster_role_binding.yaml | sed "s/{{NAMESPACE}}/${NAMESPACE}/g" | kubectl apply -f -
	oc create -f deploy/operator.yaml

.PHONY: uninstall
uninstall: uninstall-crds
	kubectl delete -f deploy/service_account.yaml -n ${NAMESPACE}
	kubectl delete -f deploy/role.yaml -n ${NAMESPACE}
	kubectl delete -f deploy/role_binding.yaml -n ${NAMESPACE}
	kubectl delete -f deploy/cluster_role.yaml
	cat deploy/cluster_role_binding.yaml | sed "s/{{NAMESPACE}}/${NAMESPACE}/g" | kubectl delete -f -
	oc delete -f deploy/operator.yaml

.PHONY: install-crds
install-crds:
	kubectl apply -f deploy/crds/app.example.com_appservices_crd.yaml

.PHONY: uninstall-crds
uninstall-crds:
	kubectl delete -f deploy/crds/app.example.com_appservices_crd.yaml
#
# .PHONY: run
# run:
#     operator-sdk up local --namespace=ms-dev --go-ldflags \"-X=appservice.ProductName=XXXYYY\"

.PHONY: deploy
deploy:
	kubectl apply -f deploy/operator.yaml -n ${NAMESPACE}

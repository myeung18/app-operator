oc create -f deploy/crds/app.example.com_appservices_crd.yaml
oc create -f  deploy/service_account.yaml
oc create -f  deploy/cluster_role.yaml
oc create -f  deploy/cluster_role_binding.yaml
oc create -f  deploy/role_binding.yaml
oc create -f  deploy/role.yaml
oc create -f  deploy/operator.yaml

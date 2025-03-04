package k8s

import "kubeants.io/service"

type K8SApi struct {
	PodApi
	NamespaceApi
	NodeApi
	ResourceApi
}

var podService = service.ServiceGroupApp.PodDerviceGroup.PodService
var namespaceService = service.ServiceGroupApp.NamespaceServiceGroup.NamespaceService
var resourceService = service.ServiceGroupApp.ResourceServiceGroup.ResourceService // 所有k8s资源原生接口

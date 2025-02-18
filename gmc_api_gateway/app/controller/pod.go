package controller

import (
	"fmt"
	"gmc_api_gateway/app/common"
	"gmc_api_gateway/app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPods godoc
// @Summary Show detail pods
// @Description get pods Details
// @Accept  json
// @Produce  json
// @Success 200 {object} model.POD
// @Router /pods/{name} [get]
// @Security   Bearer
// @Param name path string true "name of the pods"
// @Param cluster query string true "cluster Name of the pods"
// @Param workspace query string true "workspace Name of the pods"
func GetPods(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:      "pods",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	getData, err := common.DataRequest(params)
	// if err != nil {
	// 	common.ErrorMsg(c, http.StatusNotFound, err)
	// 	return nil
	// }
	if err != nil || common.InterfaceToString(common.FindData(getData, "status", "")) == "Failure" {
		msg := common.ErrorMsg2(http.StatusNotFound, common.ErrNotFound)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": msg,
		})
	}

	ownerReferencesData := common.FindData(getData, "metadata", "ownerReferences")
	var ownerReferencesInfo []model.OwnerReference
	common.Transcode(ownerReferencesData, &ownerReferencesInfo)

	podIPsData := common.FindData(getData, "status", "podIPs")
	var podIPsInfo []model.PodIPs
	common.Transcode(podIPsData, &podIPsInfo)

	containerStatusesData := common.FindData(getData, "status", "containerStatuses")
	var containerStatusesInfo []model.ContainerStatuses
	common.Transcode(containerStatusesData, &containerStatusesInfo)

	podcontainersData := common.FindData(getData, "spec", "containers")
	var podcontainersDataInfo []model.PODCONTAINERS
	common.Transcode(podcontainersData, &podcontainersDataInfo)

	StatusConditionsData := common.FindData(getData, "status", "conditions")
	var StatusConditionsInfo []model.StatusConditions
	common.Transcode(StatusConditionsData, &StatusConditionsInfo)

	// volumeMountsData := common.FindData(getData, "spec.containers", "volumeMounts")
	// var volumeMountsInfo []model.VolumeMounts
	// common.Transcode(volumeMountsData, &volumeMountsInfo)
	involvesData, _ := common.GetModelRelatedList(params)

	pod := model.POD{
		Workspace:         params.Workspace,
		Cluster:           params.Cluster,
		Name:              common.InterfaceToString(common.FindData(getData, "metadata", "name")),
		Namespace:         common.InterfaceToString(common.FindData(getData, "metadata", "namespace")),
		CreationTimestamp: common.InterfaceToTime(common.FindData(getData, "metadata", "creationTimestamp")),
		NodeName:          common.InterfaceToString(common.FindData(getData, "spec", "nodeName")),
		Lable:             common.FindData(getData, "metadata", "labels"),
		Annotations:       common.FindData(getData, "metadata", "annotations"),
		QosClass:          common.InterfaceToString(common.FindData(getData, "status", "qosClass")),
		OwnerReference:    ownerReferencesInfo,
		StatusConditions:  StatusConditionsInfo,
		Status:            common.InterfaceToString(common.FindData(getData, "status", "phase")),
		HostIP:            common.InterfaceToString(common.FindData(getData, "status", "hostIP")),
		PodIP:             common.InterfaceToString(common.FindData(getData, "status", "podIP")),
		PodIPs:            podIPsInfo,
		ContainerStatuses: containerStatusesInfo,
		Podcontainers:     podcontainersDataInfo,
		// VolumeMounts:      volumeMountsInfo,
		Events: getCallEvent(params),
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data":         pod,
		"involvesData": involvesData,
	})
}

// GetPods godoc
// @Summary Show List pods
// @Description get pods List
// @Accept  json
// @Produce  json
// @Success 200 {object} model.POD
// @Router /pods [get]
// @Security   Bearer
// @Param cluster query string false "cluster Name of the pods"
// @Param workspace query string  false "workspace Name of the pods"
func GetAllPods(c echo.Context) error {
	var pods []model.POD
	fmt.Printf("## pods", pods)
	params := model.PARAMS{
		Kind:      "pods",
		Name:      c.Param("name"),
		Cluster:   c.QueryParam("cluster"),
		Workspace: c.QueryParam("workspace"),
		Project:   c.QueryParam("project"),
		User:      c.QueryParam("user"),
		Method:    c.Request().Method,
		Body:      responseBody(c.Request().Body),
	}
	data := GetModelList(params)
	fmt.Printf("####Pod data confirm : %s", data)

	for i, _ := range data {
		var restart int
		containerStatusesData := common.FindData(data[i], "status", "containerStatuses")
		var containerStatusesInfo []model.ContainerStatuses
		common.Transcode(containerStatusesData, &containerStatusesInfo)
		for c := range containerStatusesInfo {
			restart += containerStatusesInfo[c].RestartCount
		}
		var podIP string
		var hostIP string
		if common.InterfaceToString(common.FindData(data[i], "status", "podIP")) == "" {
			podIP = "-"
		} else {
			podIP = common.InterfaceToString(common.FindData(data[i], "status", "podIP"))
		}
		if common.InterfaceToString(common.FindData(data[i], "status", "hostIP")) == "" {
			hostIP = "-"
		} else {
			hostIP = common.InterfaceToString(common.FindData(data[i], "status", "hostIP"))
		}
		pod := model.POD{
			Name:              common.InterfaceToString(common.FindData(data[i], "metadata", "name")),
			Namespace:         common.InterfaceToString(common.FindData(data[i], "metadata", "namespace")),
			Cluster:           common.InterfaceToString(common.FindData(data[i], "clusterName", "")),
			Workspace:         common.InterfaceToString(common.FindData(data[i], "workspaceName", "")),
			UserName:          common.InterfaceToString(common.FindData(data[i], "userName", "")),
			CreationTimestamp: common.InterfaceToTime(common.FindData(data[i], "metadata", "creationTimestamp")),
			Status:            common.InterfaceToString(common.FindData(data[i], "status", "phase")),
			NodeName:          common.InterfaceToString(common.FindData(data[i], "spec", "nodeName")),
			PodIP:             podIP,
			HostIP:            hostIP,
			Restart:           restart,
		}
		if params.User != "" {
			if params.User == pod.UserName {
				pods = append(pods, pod)
			}
		} else {
			pods = append(pods, pod)
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": pods,
	})
}

func CreatePod(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "pods",
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}

	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"info": common.StringToInterface(postData),
	})
}

func DeletePod(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:    "pods",
		Name:    c.Param("name"),
		Cluster: c.QueryParam("cluster"),
		Project: c.QueryParam("project"),
		Method:  c.Request().Method,
		Body:    responseBody(c.Request().Body),
	}

	postData, err := common.DataRequest(params)
	if err != nil {
		common.ErrorMsg(c, http.StatusNotFound, err)
		return nil
	}

	return c.JSON(http.StatusOK, echo.Map{
		"info": common.StringToInterface(postData),
	})
}

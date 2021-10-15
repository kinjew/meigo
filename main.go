package main

import (
	"fmt"
	mgInit "meigo/library/init"
	"meigo/library/log"
	Server "meigo/library/server"
	"meigo/routers"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
)

var ExeDir string

func main() {

	//go strCancat()

	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	ExeDir = filepath.Dir(path)
	//fmt.Println(path) // for example /home/user/main
	//fmt.Println(dir)  // for example /home/user
	// 配置读取加载
	mgInit.ConfInit(ExeDir)

	// 初始化数据库连接
	mgInit.DBInit()
	defer mgInit.DBClose()

	// 初始化路由
	router := routers.InitRouter()

	//监控
	//https://www.cnblogs.com/52fhy/p/11828448.html
	go func() {
		//提供给负载均衡探活
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))

		})

		//prometheus
		http.Handle("/metrics", promhttp.Handler())

		//pprof, go tool pprof -http=:8081 http://$host:$port/debug/pprof/heap
		http.ListenAndServe(":10108", nil)
	}()

	// 启动服务
	if err := router.Run(Server.ServerConf.Port); err != nil {
		log.Error("err", err)
	}

	/*
		// 平滑启动
		if err := graceup.ListenAndServe(Server.ServerConf.Port, router); err != nil {
			log.Error("err", err)
		}
	*/

}

func strCancat() {
	//需要先导入Strings包
	/*
		s1 := "字符串"
		s2 := "拼接"
		var build strings.Builder
		build.WriteString(s1)
		build.WriteString(s2)
		s3 := build.String()
		println(s3)
	*/

	//node信息
	var nodeOne = "1"
	var nodeTwo = "2"
	var nodeThree = "3"
	var nodeFour = "4"
	//node对应的template信息
	templateOne := "template" + nodeOne
	templateTwo := "template" + nodeTwo
	templateThree := "template" + nodeThree
	templateFour := "template" + nodeFour
	var dependencyInput, dagDependencyInputArtifacts, dagDependencyInputParams, currentTemplate, currentOutputArtifacts, currentOutputParams string

	//根据依赖关系定义dag
	//nodeOne无依赖关系
	var dagheader = `  - name: diamond
    dag:
      tasks:
`
	var dagTemplateHeader = `      - name: %s
        template: %s
`
	var dagDependencyTemplateHeader = `      - name: %s
        template: %s
        dependencies: [%s]
`
	//nodeOne是第一个节点，输入依赖于参数传递（取传递进入的参数）

	var artifactItemTemplate = `      - name: %s
        path: /tmp/%s
`
	var parametersItemTemplate = `      - name: %s
        valueFrom:
          path: /tmp/%s
`
	//dag相关
	var artifactArgumentsTemplate = `          - name: %s
            from: "{{tasks.%s.outputs.artifacts.%s}}"
`
	var parametersArgumentsTemplate = `          - name: %s
            value: "{{tasks.%s.outputs.parameters.%s}}"
`

	var templateBodyFirst = `
    inputs:
      parameters:
      - name: message       # parameter declaration
    container:
      image: docker/whalesay:latest
      command: [sh, -c]
      args: ["cowsay %s {{workflow.name}} {{inputs.parameters.message}} | tee /tmp/%s"]
    outputs:
      artifacts:
%s
      parameters:
%s
`

	var templateBodyMiddle = `
    inputs:
      artifacts:
%s
    container:
      image: docker/whalesay:latest
      command: [sh, -c]
      args: ["cowsay %s {{workflow.name}} | tee /tmp/%s"]
    outputs:
      artifacts:
%s
      parameters:
%s
`
	var dagTemplateBodyFirst = `
        arguments:
          artifacts:
%s
          parameters: [{name: message, value: "{{workflow.parameters.message}}"}]
`
	var dagTemplateBodyMiddle = `
        arguments:
          artifacts:
%s
          parameters:
%s
`

	dagTemplateOneHeader := fmt.Sprintf(dagTemplateHeader, templateOne, templateOne)
	/*
	   	var dagtemplateOneBody = `
	           arguments:
	             parameters: [{name: message, value: A}]
	   `
	*/

	//nodeTwo有依赖关系，依赖于nodeOne
	dagTemplateTwoHeader := fmt.Sprintf(dagDependencyTemplateHeader, templateTwo, templateTwo, templateOne)
	/*
	   	var dagtemplateTwoBody = `
	           arguments:
	             parameters: [{name: message, value: B}]
	   `

	*/
	//nodeThree有依赖关系，依赖于nodeOne
	dagTemplateThreeHeader := fmt.Sprintf(dagDependencyTemplateHeader, templateThree, templateThree, templateOne)
	/*
	   	var dagtemplateThreeBody = `
	           arguments:
	             parameters: [{name: message, value: C}]
	   `
	   	//nodeFour依赖关系nodetwo和nodethree


	*/
	var tmpStrSlice []string = []string{templateTwo, templateThree}
	tmpStr := strings.Join(tmpStrSlice, ",")
	dagTemplateFourHeader := fmt.Sprintf(dagDependencyTemplateHeader, templateFour, templateFour, tmpStr)
	/*
	   	var dagtemplateFourBody = `
	           arguments:
	             parameters: [{name: message, value: D}]
	   `

	*/
	var wfHeader string = `
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName:`
	var wfName string = " " + "345" + "-"

	// 构造wf的头部
	wfhead := wfHeader + wfName

	var WfSpecHeader string = `
spec:
  entrypoint: diamond
  templateDefaults:
    #timeout: 30s   # timeout value will be applied to all templates
    #retryStrategy: # retryStrategy value will be applied to all templates
    #  limit: 2
  arguments:
    parameters:
    - name: message
      value: hello world

  templates:
`
	//templateOne不依赖任何节点
	var templateOneHeader = "  - name: " + templateOne
	dependencyInput = ""
	currentTemplate = templateOne
	//生成制品
	currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateOne, templateOne)
	//生成输出参数
	currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateOne, templateOne)
	//生成模版体
	templateOneBody := fmt.Sprintf(templateBodyFirst, currentTemplate, currentTemplate, currentOutputArtifacts, currentOutputParams)
	//生成dag模版体
	dagTemplateOneBody := fmt.Sprintf(dagTemplateBodyFirst, "")

	//templateTwo依赖于templateOne的输入
	var templateTwoHeader = "  - name: " + templateTwo
	dependencyInput = fmt.Sprintf(artifactItemTemplate, templateOne, templateOne)
	currentTemplate = templateTwo
	//生成制品
	currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateTwo, templateTwo)
	//生成输出参数
	currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateTwo, templateTwo)
	//生成模版体
	templateTwoBody := fmt.Sprintf(templateBodyMiddle, dependencyInput, currentTemplate, currentTemplate, currentOutputArtifacts, currentOutputParams)
	//生成dag模版体
	dagDependencyInputArtifacts = fmt.Sprintf(artifactArgumentsTemplate, templateOne, templateOne, templateOne)
	dagDependencyInputParams = fmt.Sprintf(parametersArgumentsTemplate, templateOne, templateOne, templateOne)
	dagTemplateTwoBody := fmt.Sprintf(dagTemplateBodyMiddle, dagDependencyInputArtifacts, dagDependencyInputParams)

	//templateThree依赖于templateOne的输入
	var templateThreeHeader = "  - name: " + templateThree
	dependencyInput = fmt.Sprintf(artifactItemTemplate, templateOne, templateOne)
	currentTemplate = templateThree
	//生成制品
	currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateThree, templateThree)
	//生成输出参数
	currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateThree, templateThree)
	//生成模版体
	templateThreeBody := fmt.Sprintf(templateBodyMiddle, dependencyInput, currentTemplate, currentTemplate, currentOutputArtifacts, currentOutputParams)
	//生成dag模版体
	dagDependencyInputArtifacts = fmt.Sprintf(artifactArgumentsTemplate, templateOne, templateOne, templateOne)
	dagDependencyInputParams = fmt.Sprintf(parametersArgumentsTemplate, templateOne, templateOne, templateOne)
	dagTemplateThreeBody := fmt.Sprintf(dagTemplateBodyMiddle, dagDependencyInputArtifacts, dagDependencyInputParams)

	//templateFour依赖于templateTwo和templateThree的输入
	var templateFourHeader = "  - name: " + templateFour
	var strSlice = []string{}
	dependencyInputTemp := fmt.Sprintf(artifactItemTemplate, templateTwo, templateTwo)
	strSlice = append(strSlice, dependencyInputTemp)
	dependencyInputTemp = fmt.Sprintf(artifactItemTemplate, templateThree, templateThree)
	strSlice = append(strSlice, dependencyInputTemp)
	dependencyInput = strings.Join(strSlice, "")

	currentTemplate = templateFour
	//生成制品
	currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateFour, templateFour)
	//生成输出参数
	currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateFour, templateFour)
	//生成模版体
	templateFourBody := fmt.Sprintf(templateBodyMiddle, dependencyInput, currentTemplate, currentTemplate, currentOutputArtifacts, currentOutputParams)
	//构造制品
	var strSliceDagArtifacts = []string{}
	dagDependencyInputArtifactsTemp := fmt.Sprintf(artifactArgumentsTemplate, templateTwo, templateTwo, templateTwo)
	strSliceDagArtifacts = append(strSliceDagArtifacts, dagDependencyInputArtifactsTemp)
	dagDependencyInputArtifactsTemp = fmt.Sprintf(artifactArgumentsTemplate, templateThree, templateThree, templateThree)
	strSliceDagArtifacts = append(strSliceDagArtifacts, dagDependencyInputArtifactsTemp)
	dagDependencyInputArtifacts = strings.Join(strSliceDagArtifacts, "")
	//构造输入参数
	var strSliceDagParams = []string{}
	dagDependencyInputParamsTemp := fmt.Sprintf(parametersArgumentsTemplate, templateTwo, templateTwo, templateTwo)
	strSliceDagParams = append(strSliceDagParams, dagDependencyInputParamsTemp)
	dagDependencyInputParamsTemp = fmt.Sprintf(parametersArgumentsTemplate, templateThree, templateThree, templateThree)
	strSliceDagParams = append(strSliceDagParams, dagDependencyInputParamsTemp)
	dagDependencyInputParams = strings.Join(strSliceDagParams, "")
	//生成dag模版体
	dagTemplateFourBody := fmt.Sprintf(dagTemplateBodyMiddle, dagDependencyInputArtifacts, dagDependencyInputParams)

	//组合工作流板块
	var strTemplate []string = []string{WfSpecHeader, templateOneHeader, templateOneBody, templateTwoHeader, templateTwoBody, templateThreeHeader, templateThreeBody,
		templateFourHeader, templateFourBody}
	wfTemplate := strings.Join(strTemplate, "")

	var strDagTemplate []string = []string{dagheader, dagTemplateOneHeader, dagTemplateOneBody, dagTemplateTwoHeader, dagTemplateTwoBody, dagTemplateThreeHeader, dagTemplateThreeBody,
		dagTemplateFourHeader, dagTemplateFourBody}
	wfDagTemplate := strings.Join(strDagTemplate, "")

	var wfYamlTmp []string = []string{wfhead, wfTemplate, wfDagTemplate}
	wfYaml := strings.Join(wfYamlTmp, "")
	//println(wfYaml)
	fmt.Fprintf(os.Stdout, wfYaml)

}

/*

func main() {
	//可以求出最大子序列
	var arr = []int{2, 1, -3, 4, -1, 2, 1, -5, 4}
	//res := maxInerList(arr)
	res := maxResult(arr)
	println(res)

}

func maxInerList(arr []int) int {
	len := len(arr)
	var max, begin, end = 0, 0, 0
	for i := 0; i < len; i++ {
		var sum = 0
		for j := i; j < len; j++ {
			println("the i:", i, "the j:", j)
			sum += arr[j]
			println("the max:", max, "the sum:", sum)
			if sum >= max {
				println("the ok i:", i, "the ok j:", j)
				max = sum
				begin = i
				end = j
			}

		}

	}
	//打印结果
	println("the begin index:", begin, "the end index:", end)
	return max
}

func maxResult(arr []int) int {
	len := len(arr)
	var max, thissum = 0, 0
	for i := 0; i < len; i++ {
		if thissum+arr[i] > arr[i] {
			thissum = thissum + arr[i]
		} else {
			thissum = arr[i]
		}
		//max = max(max,thissum)
		if max < thissum {
			max = thissum
		}

		//max = (max > thissum) ? max : thissum
		//rs := math.Max(1.0,2.0)
	}
	return max
}

*/

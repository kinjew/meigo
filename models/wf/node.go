package wf

import (
	"context"
	"encoding/json"
	"fmt"
	"meigo/library/db/common"
	"meigo/library/log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/go-redis/redis/v8"
	ctxExt "github.com/kinjew/gin-context-ext"
)

// Node 实体
type Node struct {
	common.BaseModelV1
	ParentId     string `gorm:"column:parent_id;" json:"parent_id" form:"parent_id" binding:"required"`
	FlowId       int    `gorm:"column:flow_id;" json:"flow_id" form:"flow_id" `
	NodeName     string `gorm:"column:node_name;" json:"node_name" form:"node_name"`
	NodeType     string `gorm:"column:node_type;" json:"node_type" form:"node_type"`
	NodeClassify int    `gorm:"column:node_classify;" json:"node_classify" form:"node_classify"`
	Rules        string `gorm:"column:rules;" json:"rules" form:"rules"`
	Styles       string `gorm:"column:styles;" json:"styles" form:"styles"`
	Creator      string `gorm:"column:creator;" json:"creator" form:"creator"`
	Modifier     string `gorm:"column:modifier;" json:"modifier" form:"modifier"`
	IsDel        int    `gorm:"column:is_del;" json:"is_del" form:"is_del"`
}

// Flow 实体
type Flow struct {
	common.BaseModelV1
	OrgId        string `gorm:"column:org_id;" json:"org_id" form:"org_id"`
	FlowName     string `gorm:"column:flow_name;" json:"flow_name" form:"flow_name"`
	FlowStatus   int    `gorm:"column:flow_status;" json:"flow_status" form:"flow_status"`
	BeginTime    int    `gorm:"column:begin_time;" json:"begin_time" form:"begin_time"`
	EndTime      int    `gorm:"column:end_time;" json:"end_time" form:"end_time"`
	TriggerCount int    `gorm:"column:trigger_count;" json:"trigger_count" form:"trigger_count"`
	Content      string `gorm:"column:content;" json:"content" form:"content"`
	Creator      string `gorm:"column:creator;" json:"creator" form:"creator"`
	Modifier     string `gorm:"column:modifier;" json:"modifier" form:"modifier"`
	IsDel        int    `gorm:"column:is_del;" json:"is_del" form:"is_del"`
}

// FlowYaml 实体
type FlowYaml struct {
	common.BaseModelV1
	FlowId      int    `gorm:"column:flow_id;" json:"flow_id" form:"flow_id" `
	YamlContent string `gorm:"column:yaml_content;" json:"yaml_content" form:"yaml_content"`
	IsDel       int    `gorm:"column:is_del;" json:"is_del" form:"is_del"`
}

/*
ArgoYaml 生产yaml文件
*/

func (n *Node) ArgoYaml(c *ctxExt.Context) (flag bool, err error) {

	//work flow redis prefix
	wf_prefix := viper.GetString("redis.wf_prefix")
	//fmt.Println("wf_prefix: ", viper.GetString("redis.wf_prefix"))

	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:       viper.GetString("redis.addr"),
		Password:   viper.GetString("redis.password"), // no password set
		DB:         viper.GetInt("redis.DB"),          // use default DB
		MaxRetries: viper.GetInt("redis.maxRetries"),
	})
	var ctx = context.Background()

	//node_id与flow_id不能同时为空
	nodeId := c.Query("node_id")
	FlowId := c.Query("flow_id")
	if nodeId == "" && FlowId == "" {
		log.Error("node_id && flow_id err: ", err)
		return false, fmt.Errorf("node_id and flow_id are null")
	}
	//获取节点信息
	var flow_id = 0
	if FlowId == "" {
		var node Node
		nodeIdInt, _ := strconv.Atoi(nodeId)
		err = sqlDB.Table("flow_nodes").Where("id = ?", nodeIdInt).Select("* ").First(&node).Error //Map查询
		if err != nil {
			return false, err
		}
		//获取flow_id
		flow_id = node.FlowId
	} else {
		flow_id, _ = strconv.Atoi(FlowId)
	}
	//获取工作流信息
	var flow Flow
	err = sqlDB.Table("flows").Where("id = ?", flow_id).Select("* ").Scan(&flow).Error //Map查询
	if err != nil {
		return false, err
	}
	//获取所有工作流的所有节点列表
	var list []Node
	err = sqlDB.Table("flow_nodes").Where("flow_id = ?", flow_id).Select("* ").Scan(&list).Error //Map查询
	if err != nil {
		return false, err
	}
	// 节点信息处理与存入redis
	nodeInfoMap := make(map[int]Node)
	for _, v := range list {
		nodeInfoMap[v.ID] = v
		//获取当前节点信息
		json_str, _ := json.Marshal(v)
		//println(strconv.Itoa(v.ID))
		//println(v.ID, wf_prefix+strconv.Itoa(v.ID), json_str)
		_ = rdb.Set(ctx, wf_prefix+strconv.Itoa(v.ID), json_str, time.Duration(86400)*time.Second).Err()
	}
	//根据依赖关系定义dag
	/*
	   工作流信息模版
	*/

	var wfHeader string = `
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName:`
	var wfName = " " + strconv.Itoa(flow.ID) + "-"

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
      value: '{"test":"hello word"}'
  templates:
`
	/*
	   节点信息模版
	*/
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

	/*
			var artifactItemTemplate = `      - name: %s
				        path: /tmp/%s
				`

			var parametersItemTemplate = `      - name: %s
		        valueFrom:
		          path: /tmp/%s
		`
	*/
	//dag相关
	/*
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
		      image: wf:1.0
		      command: ["/app/wf-server/bin/wf"]
		      args: ["-wfUuid={{workflow.name}}","-nodeId=%s","{{inputs.parameters.message}}"]
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
		      image: wf:1.0
		      command: ["/app/wf-server/bin/wf"]
		      args: ["-wfUuid={{workflow.name}}","-nodeId=%s"]
		    outputs:
		      artifacts:
		%s
		      parameters:
		%s
		`
	*/

	var templateBodyFirst = `
    inputs:
      parameters:
      - name: message       # parameter declaration
    container:
      image: wf:1.0
      command: ["/app/wf-server/bin/wf"]
      args: ["-wfUuid={{workflow.name}}","-nodeId=%s","-message={{inputs.parameters.message}}"]
`
	var templateBodyMiddle = `
    container:
      image: wf:1.0
      command: ["/app/wf-server/bin/wf"]
      args: ["-wfUuid={{workflow.name}}","-nodeId=%s"]
`
	var dagTemplateBodyFirst = `
        arguments:
          artifacts:
%s
          parameters: [{name: message, value: "{{workflow.parameters.message}}"}]
`
	/*
			var dagTemplateBodyMiddle = `
		        arguments:
		          artifacts:
		%s
		          parameters:
		%s
		`
	*/
	//变量预定义
	//var dagTemplateOneHeader, dependencyInput, dagDependencyInputArtifacts, dagDependencyInputParams, currentTemplate, currentOutputArtifacts, currentOutputParams string
	var dagTemplateOneHeader string

	var strTemplate, strDagTemplate []string
	strTemplate = append(strTemplate, WfSpecHeader)
	strDagTemplate = append(strDagTemplate, dagheader)

	//节点和节点关系转换
	for _, v := range list {

		//节点名称
		//templateName := v.NodeName + strconv.Itoa(int(v.ID))
		templateName := "template" + strconv.Itoa(int(v.ID))
		//currentTemplate = templateName
		//根节点处理
		if v.ParentId == "" {
			//生成模版体
			//templateOne不依赖任何节点
			var templateOneHeader = "  - name: " + templateName
			/*
				//生成制品
				currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateName, templateName)
				//生成输出参数
				currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateName, templateName)
			*/
			templateOneBody := fmt.Sprintf(templateBodyFirst, strconv.Itoa(int(v.ID)))
			//生成dag模版
			dagTemplateOneHeader = fmt.Sprintf(dagTemplateHeader, templateName, templateName)
			dagTemplateOneBody := fmt.Sprintf(dagTemplateBodyFirst, "")
			//压入切片
			strTemplate = append(strTemplate, templateOneHeader, templateOneBody)
			strDagTemplate = append(strDagTemplate, dagTemplateOneHeader, dagTemplateOneBody)
		} else {
			//解析parent_id
			parentIdSlice := strings.Split(v.ParentId, ",")
			if len(parentIdSlice) <= 1 {
				//获取依赖节点
				ParentIdInt, _ := strconv.Atoi(v.ParentId)
				templateDependency := "template" + strconv.Itoa(nodeInfoMap[ParentIdInt].ID)
				//构造dag模版
				dagTemplateHeader := fmt.Sprintf(dagDependencyTemplateHeader, templateName, templateName, templateDependency)
				/*
					dagDependencyInputArtifacts = fmt.Sprintf(artifactArgumentsTemplate, templateDependency, templateDependency, templateDependency)
					dagDependencyInputParams = fmt.Sprintf(parametersArgumentsTemplate, templateDependency, templateDependency, templateDependency)
					dagTemplateBody := fmt.Sprintf(dagTemplateBodyMiddle, dagDependencyInputArtifacts, dagDependencyInputParams)
				*/

				//构造模版头
				var templateHeader = "  - name: " + templateName
				/*
					dependencyInput = fmt.Sprintf(artifactItemTemplate, templateDependency, templateDependency)
					//生成制品
					currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateName, templateName)
					//生成输出参数
					currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateName, templateName)
				*/
				//生成模版体
				templateBody := fmt.Sprintf(templateBodyMiddle, strconv.Itoa(int(v.ID)))
				//压入切片
				strTemplate = append(strTemplate, templateHeader, templateBody)
				//strDagTemplate = append(strDagTemplate, dagTemplateHeader, dagTemplateBody)
				strDagTemplate = append(strDagTemplate, dagTemplateHeader)
			} else {
				//存在多个父节点
				var tmpStrSlice []string
				for _, item := range parentIdSlice {
					itemInt, _ := strconv.Atoi(item)
					tmpStrSlice = append(tmpStrSlice, "template"+strconv.Itoa(nodeInfoMap[itemInt].ID))
				}
				tmpStr := strings.Join(tmpStrSlice, ",")
				dagTemplateHeader := fmt.Sprintf(dagDependencyTemplateHeader, templateName, templateName, tmpStr)
				var templateHeader = "  - name: " + templateName
				//templateFour依赖于templateTwo和templateThree的输入
				//构造依赖项
				/*
					var strSlice = []string{}
					var strSliceDagArtifacts = []string{}
					var strSliceDagParams = []string{}

						for _, tempItem := range tmpStrSlice {
							dependencyInputTemp := fmt.Sprintf(artifactItemTemplate, tempItem, tempItem)
							strSlice = append(strSlice, dependencyInputTemp)

							dagDependencyInputArtifactsTemp := fmt.Sprintf(artifactArgumentsTemplate, tempItem, tempItem, tempItem)
							strSliceDagArtifacts = append(strSliceDagArtifacts, dagDependencyInputArtifactsTemp)

							dagDependencyInputParamsTemp := fmt.Sprintf(parametersArgumentsTemplate, tempItem, tempItem, tempItem)
							strSliceDagParams = append(strSliceDagParams, dagDependencyInputParamsTemp)

						}
					dependencyInput = strings.Join(strSlice, "")
						//生成制品
						currentOutputArtifacts = fmt.Sprintf(artifactItemTemplate, templateName, templateName)
						//生成输出参数
						currentOutputParams = fmt.Sprintf(parametersItemTemplate, templateName, templateName)
				*/
				//生成模版体
				templateBody := fmt.Sprintf(templateBodyMiddle, strconv.Itoa(int(v.ID)))
				/*
					//构造制品
					dagDependencyInputArtifacts = strings.Join(strSliceDagArtifacts, "")
					//构造输入参数
					dagDependencyInputParams = strings.Join(strSliceDagParams, "")
					//生成dag模版体
					dagTemplateBody := fmt.Sprintf(dagTemplateBodyMiddle, dagDependencyInputArtifacts, dagDependencyInputParams)
				*/
				//压入切片
				strTemplate = append(strTemplate, templateHeader, templateBody)
				strDagTemplate = append(strDagTemplate, dagTemplateHeader)
			}
		}
	}
	//组合工作流板块
	// 切片反转 strTemplate = sliceReverse(strTemplate)
	wfTemplate := strings.Join(strTemplate, "")
	// 切片反转 strDagTemplate = sliceReverse(strDagTemplate)
	wfDagTemplate := strings.Join(strDagTemplate, "")
	var wfYamlTmp = []string{wfhead, wfTemplate, wfDagTemplate}
	wfYaml := strings.Join(wfYamlTmp, "")
	//存储工作流模版
	var flowYaml FlowYaml
	flowYamlTemp := FlowYaml{FlowId: flow_id, YamlContent: wfYaml}
	err = sqlDB.Table("flow_yamls").Where("flow_id = ?", flow_id).Select("* ").First(&flowYaml).Error //Map查询
	if err == nil && flowYaml.ID > 0 {
		//更新流程内容
		flowYamlTemp.UpdatedAt = int(time.Now().Unix())
		err = sqlDB.Table("flow_yamls").Updates(flowYamlTemp).Where("id = ?", flowYaml.ID).Error
		if err != nil {
			return false, err
		}
	} else {
		//新建流程内容
		flowYamlTemp.CreatedAt = int(time.Now().Unix())
		err = sqlDB.Table("flow_yamls").Create(&flowYamlTemp).Error
		if err != nil {
			return false, err
		}
	}
	//println(wfYaml)
	//fmt.Fprintf(os.Stdout, wfYaml)
	return true, err
}

// 切片反转
func sliceReverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

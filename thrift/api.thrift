namespace go api
include "dto/alert.thrift"
include "dto/index.thrift"
include "dto/rule.thrift"
include "dto/task.thrift"
include "dto/common.thrift"

service AlertService{
    common.ErrResponse AddAlert(1:string ruleCode,2:string time,3:i64 roomId)
    alert.AlertsResponse SelectAlert(1:i64 roomId,2:string ruleCode,3:string startTime,4:string endTime)
}

service IndexService{
    index.IndexResponse SelectIndex(1:string code)
    index.IndexsResponse SelectAllIndex()
    index.IndexResponse AddIndex(1:string name,2:string code,3:bool indexType,4:string content)
    common.ErrResponse DeleteIndex(1:string code)
    common.ErrResponse UpdateIndex(1:index.Index index)
    index.MapIndexResponse SelectRoomIndex(1:list<string> code,2:i64 roomId)
}

service RuleService{
    rule.RuleResponse SelectRule(1:string code)
    rule.RulesResponse SelectAllRule()
    rule.RuleResponse AddRule(1:i64 roomId,2:string name,3:string code,4:bool ruleType,5:string content)
    rule.CheckResponse CheckRule(1:rule.Rule rule)
    common.ErrResponse DeleteRule(1:string code)
    common.ErrResponse UpdateRule(1:rule.Rule rule)
}

service TaskService{
    task.TaskResponse SelectTask(1:string code)
    task.TaskResponse AddTask(1:string name,2:string code,3:string ruleCode,4:i64 Frequency)
    task.TasksResponse SelectRoomTask(1:i64 roomId)
    common.ErrResponse DeleteTask(1:string code)
    common.ErrResponse UpdateTask(1:task.Task task)
}
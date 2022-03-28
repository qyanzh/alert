namespace go api
include "dto/alert.thrift"
include "dto/index.thrift"
include "dto/rule.thrift"
include "dto/task.thrift"

service AlertService{
    void AddAlert(1:i64 ruleId,2:string time,3:map<i64,double> indexNum)
    alert.AlertsResponse SelectAlert(1:i64 ruleId,2:string startTime,3:string endTime)
}

service IndexService{
    index.IndexResponse SelectIndex(1:string code)
    index.IndexsResponse SelectAllIndex()
    index.IndexResponse AddIndex(1:string name,2:string code,3:string content,4:i64 timeRange)
    void DeleteIndex(1:string code)
    void UpdateIndex(1:index.Index index)
    index.MapIndexResponse SelectRoomIndex(1:list<string> code,2:i64 roomId)
}

service RuleService{
    rule.RuleResponse SelectRule(1:string code)
    rule.RulesResponse SelectAllRule()
    rule.RuleResponse AddRule(1:i64 roomId,2:string name,3:string code,4:bool ruleType,5:string content)
    rule.CheckResponse CheckRule(1:string rule)
    void DeleteRule(1:string code)
    void UpdateRule(1:rule.Rule rule)
}

service TaskService{
    task.TaskResponse SelectTask(1:string code)
    task.TaskResponse AddTask(1:string name,2:string code,3:i64 ruleId,4:i64 Frequency)
    void UpdateTaskEnable(1:string code,2:bool enable)
    task.TasksResponse SelectRunnableTask()
    void DeleteTask(1:string code)
    void UpdateTask(1:task.Task task)
}
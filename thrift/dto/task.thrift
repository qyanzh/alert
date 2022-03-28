namespace go rpc_dto

struct Task{
    1:i64 Id
    2:string Code
    3:string Name
    4:i64 RuleId
    5:i64 Frequency
    6:bool Enable
    7:string NextTime
    8:i16 LastStatus
}
struct TaskResponse{
    1:Task task
}
struct TasksResponse{
    1:list<Task> tasks
}

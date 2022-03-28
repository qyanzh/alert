namespace go rpc_dto
struct Alert{
    1:i64 Id
    2:string Time
    3:map<i64,double> IndexNum
}

struct AlertsResponse{
    1:list<Alert> alerts
}
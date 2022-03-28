namespace go rpc_dto
struct Alert{
    1:i64 Id
    2:string Time
    3:i64 RoomId
}

struct AlertsResponse{
    1:list<Alert> alerts
}
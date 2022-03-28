namespace go rpc_dto

struct Index{
    1:i64 Id
    2:string Code
    3:string Name
    4:bool Type
    5:string Expr
    6:i64 TimeRange
}
struct IndexResponse{
    1:Index index
}
struct IndexsResponse{
    1:list<Index> indexs
}
struct MapIndexResponse{
    1:map<string,double> indexs
}
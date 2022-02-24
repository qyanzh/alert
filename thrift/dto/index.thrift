namespace go rpc_dto

struct Index{
    1:i64 Id
    2:string Code
    3:string Name
    4:bool Type
    5:string Expr
    6:string TimeRange
}
struct IndexResponse{
    1:Index index
    2:string err
}
struct IndexsResponse{
    1:list<Index> indexs
    2:string err
}
struct MapIndexResponse{
    1:map<string,Index> indexs
    2:string err
}
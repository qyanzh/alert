namespace go rpc_dto

struct Rule{
    1:i64 Id
    2:string Code
    3:i64 RoomId
    4:string Name
    5:bool Type
    6:string Expr
}
struct RuleResponse{
    1:Rule rule
    2:string err
}
struct RulesResponse{
    1:list<Rule> rules
    2:string err
}
struct CheckResponse{
    1:bool result
    2:string err
}
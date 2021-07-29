package load

func Default(value interface{}, foo interface{}) interface{} {
    if value == nil || value.(string) == "" {
        return foo
    }
    return value
}

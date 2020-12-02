Week02 作业题目：

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

这种情况应该用fmt.Errorf中的%w参数，并在上层逻辑用errors.Is检查error是否为sql.ErrNoRows
代码见homework.go
# 学习笔记
 
## 我们在数据库操作的时候，比如dao层中遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error,抛给上层？为什么，应该怎么做请写出代码？

## 要Wrap吧，debug还是得到最底层啊。
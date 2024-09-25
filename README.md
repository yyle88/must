# must
must means assert means require. while the assert/require are using in testcase. must is using in main code.

这个包的作用就是断言，比如该为真的地方返回假，或者该没有err的地方返回错误，就直接崩掉就行。

## 安装
```
go get github.com/yyle88/must
```

## 简单样例
```
res, err := run()
must.Done(err) //假如出现错误-就崩掉
must.Nice(res) //假如结果为空-就崩掉
```

```
res, ok := run()
must.True(ok)  //假如结果非真-就崩掉
must.Nice(res) //假如结果为空-就崩掉
```

## 特别注意
在代码里随意 panic 是不规范的，或者说，是不道德的，就像随地内个啥。

因此这个包只建议在写 demo/test 时用，而不建议在写 service 时用。

## 起名思路

因为这俩是我常用的包，因此我就不用 assert 或者 require 作为模块名啦。
```
"github.com/stretchr/testify/assert"
"github.com/stretchr/testify/require"
```

相比之下我觉得 `must` 是不错的选择。

在汉语式英语里，must表示必然/必须，相对更强硬些，但是在英语里，似乎表示务必，但真搞不定就搞不定吧，这就是非母语编程不得劲的地方。

因此在这里说明下起名的思路。

我只认识4个字母的单词，5个字母的都认识不全，因此我还是倾向于使用4个字母的单词做模块名/函数名，这样用起来也方便。

## 断言三剑客
现在已经有 `done` 和 `sure` 两个基础包，现在再有 `must` 就相当于有三剑客啦。

sure: [软硬兼施](https://github.com/yyle88/sure)

done: [确保结果](https://github.com/yyle88/done)

must: [不对就崩](https://github.com/yyle88/must)

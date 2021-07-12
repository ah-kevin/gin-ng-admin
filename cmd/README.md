# `/cmd`

本项目的主要应用。

每个应用程序的目录名称应与您想要的可执行文件的名称相匹配（例如，/cmd/myapp）。

不要在应用程序目录中放很多代码。 如果您认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。 如果代码不可重用，或者您不希望其他人重用它，请将该代码放在 /internal 目录中。 你会惊讶于其他人会做什么，所以要明确你的意图！

通常有一个小的 main 函数从 /internal 和 /pkg 目录导入和调用代码，没有别的。
Examples:

* https://github.com/heptio/ark/tree/master/cmd (just a really small `main` function with everything else in packages)
* https://github.com/moby/moby/tree/master/cmd
* https://github.com/prometheus/prometheus/tree/master/cmd
* https://github.com/influxdata/influxdb/tree/master/cmd
* https://github.com/kubernetes/kubernetes/tree/master/cmd
* https://github.com/satellity/satellity/tree/master/cmd/satellity
* https://github.com/dapr/dapr/tree/master/cmd

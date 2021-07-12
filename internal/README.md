# `/internal`

私有应用程序和库代码。 这是您不希望其他人在他们的应用程序或库中导入的代码。 请注意，此布局模式由 Go 编译器本身强制执行。 有关更多详细信息，请参阅 [Go 1.4 发行说明](https://golang.org/doc/go1.4#internalpackages) 请注意，您不仅限于顶级内部目录。 您可以在项目树的任何级别拥有多个内部目录。

您可以选择向内部包添加一些额外的结构，以分隔共享和非共享内部代码。 这不是必需的（特别是对于较小的项目），但是有视觉线索显示预期的包用途是很好的。 您的实际应用程序代码可以位于 /internal/app 目录（例如 /internal/app/myapp）中，而这些应用程序共享的代码位于 /internal/pkg 目录（例如 /internal/pkg/myprivlib）中。
Examples:

* https://github.com/hashicorp/terraform/tree/master/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/master/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/master/internal

## `/internal/pkg`

Examples:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg

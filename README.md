# Goravel Modules

> An Modules For A Goravel Extend Package

> 模块化开发中

## Directory Structure

This is a directory standard, but you can change it if you like.

| Directory        | Action           |
| -----------      | --------------   |
| commands         | Store the command files   |
| config           | Store the config files   |
| middleware        | Store the middleware files   |
| example        | Store the example files   |

## Install

1. Add package

```
go get -u github.com/hongyukeji/goravel-modules
```

2. Register service provider

```
// config/app.go
modules "github.com/hongyukeji/goravel-modules"

"providers": []foundation.ServiceProvider{
    ...
    &modules.ServiceProvider{},
}
```

3. Publish Configuration

```
go run . artisan vendor:publish --package=github.com/hongyukeji/goravel-modules
```

4. Testing

```
// config/modules.go 文件中，打开注释
example "github.com/hongyukeji/goravel-modules/example"
&example.ServiceProvider{},

访问 http://127.0.0.1:3000/api/example/ping
```

The console will print `name` and `version`.

5. Use

参考 `github.com/hongyukeji/goravel-modules/example` ，
在 `应用根目录/modules` 目录下创建模块文件夹，
然后注册 `config/modules.go` 文件 `providers` 中。

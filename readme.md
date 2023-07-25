## 我是谁

`pixiu`面板，一个对接了`青龙`、`BBK`的管理面板

### [预览地址](https://pixiu.lxh.io)

## 我能干啥

- [x] 账号密码登录
- [x] 绑定京东账号
- [x] 查看收益记录
- [x] 绑定推送渠道
- [x] 检查是否过期

## 依赖推送组件

1. [wxhelper](https://github.com/ttttupup/wxhelper)(PC微信HOOK)
2. [cqhttp](https://docs.go-cqhttp.org/)(QQ机器人)
3. 其他推送渠道待开发

## 快速启动

1. 下载`docker-compose.yaml`文件和`config.example.yaml`文件。
2. 将`config.example.yaml`文件改名为`config.yaml`，然后放在`docker-compose.yaml`同级目录下的`config`目录(手动新建)
   ，然后酌情修改配置。
3. 执行`docker compose up -d`命令(`docker`需要安装有`docker-compose`插件)。
4. 打开浏览器访问`http://{ip}:1080`即可

## 小tips

- 修改配置文件可不用重启项目，会自动更新
- 青龙需要自行搭建

## 预览图

<table>
  <tr>
    <td><img src="https://s2.loli.net/2023/07/25/YGPwtdQ8ZA1LaTD.png"></td>
    <td><img src="https://s2.loli.net/2023/07/25/FLIP6QgwHv59mEi.png"></td>
  </tr>
  <tr>
    <td><img src="https://s2.loli.net/2023/07/25/62YMUmVi9Duw8k3.png"></td>
    <td><img src="https://s2.loli.net/2023/07/25/rgZWRqO3nycNlLI.png"></td>
  </tr>
  <tr>
    <td><img src="https://s2.loli.net/2023/07/25/s4Qo3tyJHDWV5bk.png"></td>
    <td><img src="https://s2.loli.net/2023/07/25/X4cPOGEoMgUpIla.png"></td>
  </tr>
  <tr>
    <td><img src="https://s2.loli.net/2023/07/25/CMwIO4mpuW1XS39.png"></td>
    <td><img src="https://s2.loli.net/2023/07/25/q2fzvTHBW9yKiox.png"></td>
  </tr>
</table>
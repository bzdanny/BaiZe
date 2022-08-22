## 平台简介

白泽是一套全部开源的快速开发平台，毫无保留给个人及企业免费使用。

* 前端采用Vue3、Element UI。
* 后端采用Gin、Zap、Redis、wire & Jwt。
* 权限认证使用Jwt，支持多终端认证系统。
* 支持加载动态权限菜单，多方式轻松权限控制。
* 高效率开发，使用代码生成器可以一键生成前后端代码。(正在发开)
* 特别鸣谢：[ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue?_from=gitee_search )，
  [element](https://github.com/ElemeFE/element )，[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin )
  ，[eladmin-web](https://github.com/elunez/eladmin-web )。

<p>随手 star ⭐是一种美德。 你们的star就是我的动力</p>

## 内置功能

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 通知公告：系统通知公告信息发布维护。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 在线用户：当前系统中活跃用户状态监控。
11. 服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
12. 定时任务：在线（添加、修改、删除)任务调度。(2.0版本未上线)
13. 系统接口：根据业务代码注释自动生成相关的api接口文档。
14. 代码生成：前后端代码的生成（Go、vue、js、sql）支持CRUD下载 。(2.0版本未上线)
15. k8s管理: (未上线)



## 版本规则

v5.6.7<br>
1位为主版本号（5）：当功能模块有较大的变动，比如增加多个模块或者整体架构发生变化,此版本号由项目决定是否修改。
<br>
2为次版本号（6）：当功能有一定的增加或变化，比如增加了或修改了API接口、数据库结构发生变化、增加自定义视图等功能。此版本号由项目决定是否修改。
<br>
3为阶段版本号(7)：一般是 Bug 修复或是一些小的变动，要经常发布修订版，时间间隔不限。
<br>
主版本号升级请参考更新说明更新修改或添加相应的数据表。
次版本号升级请参考更新说明查看API接口修改情况。
阶段版本号不会影响数据库与api接口，除修复重大bug不更新说明文档

## 在线体验

- admin/admin123

演示地址：https://demo.ibaize.vip
<br>
文档地址：https://doc.ibaize.vip
<br>
gitee地址：https://gitee.com/smell2/BaiZe
<br>
github地址:https://github.com/bzdanny/BaiZe.git

## 演示图

<table>
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110241805797.jpg"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110241806256.jpg"/></td>
    </tr>
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110242322137.png"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110242323820.png"/></td>
    </tr>  
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202112082243214.png"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202112082242154.png"/></td>
    </tr>

</table>

## 白泽管理系统交流群

QQ群： [![加入QQ群](https://img.shields.io/badge/83064682-blue.svg)](https://qm.qq.com/cgi-bin/qm/qr?k=rAIw_VQ_blbSQu0J6fApnm5RbAc2CHbp&jump_from=webapi)
点击按钮入群。
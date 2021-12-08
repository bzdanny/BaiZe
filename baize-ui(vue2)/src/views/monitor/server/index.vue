<template>
  <div class="app-container">
    <el-row>
      <el-col :span="12" class="card-box">
        <el-card>
          <div slot="header"><span>CPU</span></div>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="is-leaf"><div class="cell">属性</div></th>
                  <th class="is-leaf"><div class="cell">值</div></th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td><div class="cell">核心数</div></td>
                  <td><div class="cell">{{ server.cpuNum }}</div></td>
                </tr>  <tr>
                  <td><div class="cell">线程数</div></td>
                  <td><div class="cell">{{ server.cpuNumThread }}</div></td>
                </tr>
                <tr>
                  <td><div class="cell">用户使用率</div></td>
                  <td><div class="cell">{{ server.cpuUsed }}%</div></td>
                </tr>
                <tr>
                  <td><div class="cell">系统使用率</div></td>
                  <td><div class="cell" >{{ server.cpuAvg5 }}%</div></td>
                </tr>
                <tr>
                  <td><div class="cell">当前空闲率</div></td>
                  <td><div class="cell">{{ server.cpuAvg15 }}%</div></td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12" class="card-box">
        <el-card>
          <div slot="header"><span>内存</span></div>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="is-leaf"><div class="cell">属性</div></th>
                  <th class="is-leaf"><div class="cell">内存</div></th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td><div class="cell">总内存</div></td>
                  <td><div class="cell">{{ server.memTotal }}G</div></td>
                </tr>
                <tr>
                  <td><div class="cell">已用内存</div></td>
                  <td><div class="cell" >{{ server.memUsed}}G</div></td>
                </tr>
                <tr>
                  <td><div class="cell">GO使用内存</div></td>
                  <td><div class="cell" >{{ server.goUsed}}G</div></td>
                </tr>
                <tr>
                  <td><div class="cell">剩余内存</div></td>
                  <td><div class="cell" >{{ server.memFree }}G</div></td>
                </tr>
                <tr>
                  <td><div class="cell">使用率</div></td>
                  <td><div class="cell" :class="{'text-danger': server.memUsage > 80}">{{ server.memUsage}}%</div></td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="24" class="card-box">
        <el-card>
          <div slot="header">
            <span>服务器信息</span>
          </div>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <tbody>
                <tr>
                  <td><div class="cell">服务器名称</div></td>
                  <td><div class="cell" >{{ server.sysComputerName }}</div></td>
                  <td><div class="cell">操作系统</div></td>
                  <td><div class="cell" >{{ server.sysOsName }}</div></td>
                </tr>
                <tr>
                  <td><div class="cell">服务器IP</div></td>
                  <td><div class="cell" >{{ server.sysComputerIp }}</div></td>
                  <td><div class="cell">系统架构</div></td>
                  <td><div class="cell">{{ server.sysOsArch}}</div></td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="24" class="card-box">
        <el-card>
          <div slot="header">
            <span>服务信息</span>
          </div>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <tbody>
                <tr>
                  <td><div class="cell">语言环境</div></td>
                  <td><div class="cell" >{{ server.goName }}</div></td>
                  <td><div class="cell">环境版本</div></td>
                  <td><div class="cell" >{{ server.goVersion}}</div></td>
                </tr>
                <tr>
                  <td><div class="cell">启动时间</div></td>
                  <td><div class="cell" >{{ server.goStartTime}}</div></td>
                  <td><div class="cell">运行时长</div></td>
                  <td><div class="cell" >{{ server.goRunTime }}</div></td>
                </tr>
                <tr>
                  <td colspan="1"><div class="cell">安装路径</div></td>
                  <td colspan="3"><div class="cell" >{{ server.goHome }}</div></td>
                </tr>
                <tr>
                  <td colspan="1"><div class="cell">项目路径</div></td>
                  <td colspan="3"><div class="cell" >{{ server.goUserDir }}</div></td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="24" class="card-box">
        <el-card>
          <div slot="header">
            <span>磁盘状态</span>
          </div>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="is-leaf"><div class="cell">盘符路径</div></th>
                  <th class="is-leaf"><div class="cell">总大小</div></th>
                  <th class="is-leaf"><div class="cell">可用大小</div></th>
                  <th class="is-leaf"><div class="cell">已用大小</div></th>
                  <th class="is-leaf"><div class="cell">已用百分比</div></th>
                </tr>
              </thead>
              <tbody v-if="server.diskList">
                <tr v-for="sysFile in server.diskList">
                  <td><div class="cell">{{ sysFile.path }}</div></td>
                  <td><div class="cell">{{ sysFile.total }}</div></td>
                  <td><div class="cell">{{ sysFile.free }}</div></td>
                  <td><div class="cell">{{ sysFile.used }}</div></td>
                  <td><div class="cell" :class="{'text-danger': sysFile.usedPercent > 80}">{{ sysFile.usedPercent }}%</div></td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getServer } from "@/api/monitor/server";

export default {
  name: "Server",
  data() {
    return {
      // 加载层信息
      loading: [],
      // 服务器信息
      server: {}
    };
  },
  created() {
    this.getList();
    this.openLoading();
  },
  methods: {
    /** 查询服务器信息 */
    getList() {
      getServer().then(response => {
        console.log(response.data)
        this.server = response.data;
        this.loading.close();
      });
    },
    // 打开加载层
    openLoading() {
      this.loading = this.$loading({
        lock: true,
        text: "拼命读取中",
        spinner: "el-icon-loading",
        background: "rgba(0, 0, 0, 0.7)"
      });
    }
  }
};
</script>

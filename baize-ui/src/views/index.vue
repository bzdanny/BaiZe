<template>
  <div class="app-container home">
    <el-row :gutter="20">
      <el-col :sm="24" :lg="12" style="padding-left: 20px">
        <h2>白泽后台管理框架</h2>
        <p>
          从JAVA转GO后一直想做一款后台管理系统，看了很多优秀的开源项目但是发现没有合适自己的。于是利用空闲休息时间用若依管理系统开始改写自己写一套GO的后台系统。如此有了白泽管理系统。他可以用于所有的Web应用程序，如网站管理后台，网站会员中心，CMS，CRM，OA等等，当然，您也可以对她进行深度定制，以做出更强系统。所有前端后台代码封装过后十分精简易上手，出错概率低。同时支持移动客户端访问。系统会陆续更新一些实用功能。
        </p>
        <p>
          <b>当前版本:</b> <span>v{{ version }}</span>
        </p>
        <p>
          <el-tag type="danger">&yen;免费开源</el-tag>
        </p>
        <p>
          <el-button
              type="primary"
              size="mini"
              icon="MostlyCloudy"
              plain
              @click="goTarget('https://gitee.com/smell2/BaiZe')"
          >访问码云
          </el-button
          >
          <el-button
              size="mini"
              icon="HomeFilled"
              plain
              @click="goTarget('http://www.ibaize.vip')"
          >访问主页
          </el-button
          >
        </p>
      </el-col>

      <el-col :sm="24" :lg="12" style="padding-left: 50px">
        <el-row>
          <el-col :span="12">
            <h2>技术选型</h2>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="6">
            <h4>后端技术</h4>
            <ul>
              <li>Gin</li>
              <li>JWT</li>
              <li>sqlx</li>
              <li>snowflake</li>
              <li>zap</li>
              <li>fsnotify</li>
              <li>...</li>
            </ul>
          </el-col>
          <el-col :span="6">
            <h4>前端技术</h4>
            <ul>
              <li>Vue</li>
              <li>Vuex</li>
              <li>Element-ui</li>
              <li>Axios</li>
              <li>Sass</li>
              <li>Quill</li>
              <li>...</li>
            </ul>
          </el-col>
        </el-row>
      </el-col>
    </el-row>
    <el-divider/>
    <el-row :gutter="20">
      <el-col :xs="24" :sm="24" :md="12" :lg="12">
        <el-card class="update-log">
          <div slot="header" class="clearfix">
            <span>联系信息</span>
          </div>
          <div class="body">
            <p>
              <i class="el-icon-s-promotion"></i> 官网：
              <el-link
                  href="http://www.ibaize.vip"
                  target="_blank"
              >http://www.ibaize.vip</el-link
              >
            </p>
            <p>
              <i class="el-icon-user-solid"></i> QQ群：

              <a href="https://qm.qq.com/cgi-bin/qm/qr?k=rAIw_VQ_blbSQu0J6fApnm5RbAc2CHbp&jump_from=webapi"
                 target="_blank"
              > 83064682</a
              >
            </p>
          </div>

          <div class="body">
            <img
                src="https://gitee.com/smell2/BaiZe/raw/imgs/202110242200142.jpg"
                style=" width:100%"
                alt="donate"

            />
            <span style="display: inline-block; height: 30px; line-height: 30px"
            >你可以请作者喝杯咖啡表示鼓励</span
            >
          </div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="24" :md="12" :lg="12">
        <el-card class="update-log">
          <div slot="header" class="clearfix">
            <span>更新日志</span>
          </div>
          <el-collapse accordion>
            <el-collapse-item  title="v1.1.0 - 2021-12-26">
              <ol>
                <li>添加swagger</li>
              </ol>
            </el-collapse-item>
          </el-collapse>
          <el-collapse accordion>
            <el-collapse-item  title="v1.0.0 - 2021-12-08">
              <ol>
                <li>白泽系统正式发布</li>
              </ol>
            </el-collapse-item>
          </el-collapse>


        </el-card>
      </el-col>

    </el-row>
  </div>
</template>

<script setup name="Index">
const version = ref('1.0.0')

function goTarget(url) {
  window.open(url, '__blank')
}
</script>

<style scoped lang="scss">
.home {
  blockquote {
    padding: 10px 20px;
    margin: 0 0 20px;
    font-size: 17.5px;
    border-left: 5px solid #eee;
  }

  hr {
    margin-top: 20px;
    margin-bottom: 20px;
    border: 0;
    border-top: 1px solid #eee;
  }

  .col-item {
    margin-bottom: 20px;
  }

  ul {
    padding: 0;
    margin: 0;
  }

  font-family: "open sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
  font-size: 13px;
  color: #676a6c;
  overflow-x: hidden;

  ul {
    list-style-type: none;
  }

  h4 {
    margin-top: 0px;
  }

  h2 {
    margin-top: 10px;
    font-size: 26px;
    font-weight: 100;
  }

  p {
    margin-top: 10px;

    b {
      font-weight: 700;
    }
  }

  .update-log {
    ol {
      display: block;
      list-style-type: decimal;
      margin-block-start: 1em;
      margin-block-end: 1em;
      margin-inline-start: 0;
      margin-inline-end: 0;
      padding-inline-start: 40px;
    }
  }
}
</style>



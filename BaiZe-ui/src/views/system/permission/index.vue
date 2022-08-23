<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryRef" :inline="true" v-show="showSearch">
      <el-form-item label="权限名称" prop="permissionName">
        <el-input
            v-model="queryParams.permissionName"
            placeholder="请输入权限名称"
            clearable
            @keyup.enter="handleQuery"
        />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="权限状态" clearable>
          <el-option
              v-for="dict in sys_normal_disable"
              :key="dict.value"
              :label="dict.label"
              :value="dict.value"
          />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
            type="primary"
            plain
            icon="Plus"
            @click="handleAdd"
            v-hasPermi="['system:permission:add']"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
            type="info"
            plain
            icon="Sort"
            @click="toggleExpandAll"
        >展开/折叠
        </el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table
        v-if="refreshTable"
        v-loading="loading"
        :data="permissionList"
        row-key="permissionId"
        :default-expand-all="isExpandAll"
        :tree-props="{ children: 'children', hasChildren: 'hasChildren' }"
    >
      <el-table-column prop="permissionName" label="权限名称" :show-overflow-tooltip="true"
                       ></el-table-column>
      <el-table-column prop="perms" label="权限标识" :show-overflow-tooltip="true"></el-table-column>
      <el-table-column prop="status" label="状态" >
        <template #default="scope">
          <dict-tag :options="sys_normal_disable" :value="scope.row.status"/>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createTime">
        <template #default="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="200" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-button
              type="text"
              icon="Edit"
              @click="handleUpdate(scope.row)"
              v-hasPermi="['system:permission:edit']"
          >修改
          </el-button>
          <el-button
              type="text"
              icon="Plus"
              @click="handleAdd(scope.row)"
              v-hasPermi="['system:permission:add']"
          >新增
          </el-button>
          <el-button
              type="text"
              icon="Delete"
              @click="handleDelete(scope.row)"
              v-hasPermi="['system:Permission:remove']"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加或修改权限对话框 -->
    <el-dialog :title="title" v-model="open" width="680px" append-to-body>
      <el-form ref="permissionRef" :model="form" :rules="rules" label-width="100px">
        <el-row>
          <el-col :span="24">
            <el-form-item label="上级权限">
              <el-tree-select
                  v-model="form.parentId"
                  :data="permissionOptions"
                  :props="{ value: 'permissionId', label: 'permissionName', children: 'children' }"
                  value-key="permissionId"
                  placeholder="选择上级权限"
                  check-strictly
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="权限名称" prop="permissionName">
              <el-input v-model="form.permissionName" placeholder="请输入权限名称"/>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="权限字符"  prop="perms">
              <el-input v-model="form.perms" placeholder="请输入权限标识" />
              <template #label>
                        <span>
                           <el-tooltip
                               content="控制器中定义的权限字符，如：@PreAuthorize(`@ss.hasPermi('system:user:list')`)"
                               placement="top">
                              <el-icon><question-filled/></el-icon>
                           </el-tooltip>
                           权限字符
                        </span>
              </template>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup name="Permission">
import {addPermission, delPermission, getPermission, listPermission, updatePermission} from "@/api/system/permission";


const {proxy} = getCurrentInstance();
const {sys_show_hide, sys_normal_disable} = proxy.useDict("sys_show_hide", "sys_normal_disable");

const permissionList = ref([]);
const open = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const title = ref("");
const permissionOptions = ref([]);
const isExpandAll = ref(false);
const refreshTable = ref(true);
const showChooseIcon = ref(false);
const iconSelectRef = ref(null);

const data = reactive({
  form: {},
  queryParams: {
    permissionName: undefined,
    visible: undefined
  },
  rules: {
    permissionName: [{required: true, message: "权限名称不能为空", trigger: "blur"}],
    perms: [{required: true, message: "权限字符不能为空", trigger: "blur"}],
  },
});

const {queryParams, form, rules} = toRefs(data);

/** 查询权限列表 */
function getList() {
  loading.value = true;
  listPermission(queryParams.value).then(response => {
    permissionList.value = proxy.handleTree(response.data, "permissionId");
    loading.value = false;
  });
}

/** 查询权限下拉树结构 */
function getTreeselect() {
  permissionOptions.value = [];
  listPermission().then(response => {
    const permission = {permissionId: 0, permissionName: "主类目", children: []};
    permission.children = proxy.handleTree(response.data, "permissionId");
    permissionOptions.value.push(permission);
  });
}

/** 取消按钮 */
function cancel() {
  open.value = false;
  reset();
}

/** 表单重置 */
function reset() {
  form.value = {
    permissionId: undefined,
    parentId: 0,
    permissionName: undefined,
    status: "0"
  };
  proxy.resetForm("permissionRef");
}



/** 搜索按钮操作 */
function handleQuery() {
  getList();
}

/** 重置按钮操作 */
function resetQuery() {
  proxy.resetForm("queryRef");
  handleQuery();
}

/** 新增按钮操作 */
function handleAdd(row) {
  reset();
  getTreeselect();
  if (row != null && row.permissionId) {
    form.value.parentId = row.permissionId;
  } else {
    form.value.parentId = 0;
  }
  open.value = true;
  title.value = "添加权限";
}

/** 展开/折叠操作 */
function toggleExpandAll() {
  refreshTable.value = false;
  isExpandAll.value = !isExpandAll.value;
  nextTick(() => {
    refreshTable.value = true;
  });
}

/** 修改按钮操作 */
async function handleUpdate(row) {
  reset();
  await getTreeselect();
  getPermission(row.permissionId).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = "修改权限";
  });
}

/** 提交按钮 */
function submitForm() {
  proxy.$refs["permissionRef"].validate(valid => {
    if (valid) {
      if (form.value.permissionId != undefined) {
        updatePermission(form.value).then(response => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
        });
      } else {
        addPermission(form.value).then(response => {
          proxy.$modal.msgSuccess("新增成功");
          open.value = false;
          getList();
        });
      }
    }
  });
}

/** 删除按钮操作 */
function handleDelete(row) {
  proxy.$modal.confirm('是否确认删除名称为"' + row.permissionName + '"的数据项?').then(function () {
    return delPermission(row.permissionId);
  }).then(() => {
    getList();
    proxy.$modal.msgSuccess("删除成功");
  }).catch(() => {
  });
}

getList();
</script>

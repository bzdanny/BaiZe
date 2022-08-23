import request from '@/utils/request'

// 查询权限列表
export function listPermission(query) {
  return request({
    url: '/system/permission/list',
    method: 'get',
    params: query
  })
}

// 查询权限详细
export function getPermission(permissionId) {
  return request({
    url: '/system/permission/' + permissionId,
    method: 'get'
  })
}

// 查询权限下拉树结构
export function treeselect() {
  return request({
    url: '/system/permission/treeselect',
    method: 'get'
  })
}

// 根据角色ID查询权限下拉树结构
export function rolePermissionTreeselect(permissionId) {
  return request({
    url: '/system/permission/rolePermissionTreeselect/' + permissionId,
    method: 'get'
  })
}

// 新增权限
export function addPermission(data) {
  return request({
    url: '/system/permission',
    method: 'post',
    data: data
  })
}

// 修改权限
export function updatePermission(data) {
  return request({
    url: '/system/permission',
    method: 'put',
    data: data
  })
}

// 删除权限
export function delPermission(permission) {
  return request({
    url: '/system/permission/' + permission,
    method: 'delete'
  })
}
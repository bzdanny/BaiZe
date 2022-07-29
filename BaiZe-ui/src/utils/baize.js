/**
 * 通用js方法封装处理
 * Copyright (c) 2021 baize
 */
export function handleProps(data, id, label) {
  let config = {
    id: id || 'id',
    label: label || 'label',
    parentId: 'parentId',
    childrenList: 'children'
  }

  var childrenListMap = {}
  var nodeIds = {}
  var tree = []

  for (let d of data) {
    let parentId = d[config.parentId]
    if (childrenListMap[parentId] == null) {
      childrenListMap[parentId] = []
    }
    nodeIds[d[config.id]] = props(d[config.id],d[config.label])
    childrenListMap[parentId].push(props(d[config.id],d[config.label]))
  }

  for (let d of data) {
    let parentId = d[config.parentId]
    if (nodeIds[parentId] == null) {
      tree.push(props(d[config.id],d[config.label]))
    }
  }
  for (let t of tree) {
    adaptToChildrenList(t)
  }

  function adaptToChildrenList(o) {
    if (childrenListMap[o['id']] !== null) {
      o[config.childrenList] = childrenListMap[o['id']]
    }
    if (o[config.childrenList]) {
      for (let c of o[config.childrenList]) {
        adaptToChildrenList(c)
      }
    }
  }

  return tree
}

function props(id, label) {
  let o = new Object()
  o.id = id
  o.label = label
  return o
}

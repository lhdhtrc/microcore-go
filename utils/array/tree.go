package array

import "reflect"

func Tree[T interface{}](list []T) []T {
	var tree []T

	// 预分配地图，提高性能 (可选)
	parentMap := make(map[string][]T, len(list))

	for _, node := range list {
		nodeValue := reflect.ValueOf(node).Elem() // 获取节点的反射值
		parentId := nodeValue.FieldByName("ParentId").String()

		if parentId == "" {
			tree = append(tree, node)
		} else {
			// 直接使用 map 查找和追加，提高效率
			parentMap[parentId] = append(parentMap[parentId], node)
		}
	}

	for _, node := range list {
		nodeValue := reflect.ValueOf(node).Elem()
		childrenField := nodeValue.FieldByName("Children")

		// 处理无效或非切片类型的 Children 字段
		if !childrenField.IsValid() || childrenField.Kind() != reflect.Slice {
			continue
		}

		// 直接从预建的 map 获取子节点
		children, ok := parentMap[nodeValue.FieldByName("ID").String()]

		// 处理没有子节点的情况
		if !ok {
			continue
		}

		// 使用 AppendSlice 设置子节点，效率更高
		childrenField.Set(reflect.AppendSlice(childrenField, reflect.ValueOf(children)))
	}

	return tree
}

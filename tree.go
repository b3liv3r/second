package tree

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// TreeNode представляет узел бинарного дерева
type TreeNode struct {
	Key    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// insertBalancedTree вставляет элемент key в сбалансированное бинарное дерево
func insertBalancedTree(root **TreeNode, parent *TreeNode, key int, addedKeys map[int]bool) {
	if *root == nil {
		*root = &TreeNode{Key: key, Parent: parent}
		addedKeys[key] = true
		return
	}

	if key <= (*root).Key {
		insertBalancedTree(&((*root).Left), *root, key, addedKeys)
	} else {
		insertBalancedTree(&((*root).Right), *root, key, addedKeys)
	}
}

// addElementsToTree добавляет 5 случайных элементов к дереву
func addElementsToTree(tree **TreeNode, addedKeys map[int]bool) {
	for i := 0; i < 5; {
		element := rand.Intn(1000)
		if _, exists := addedKeys[element]; !exists {
			insertBalancedTree(tree, nil, element, addedKeys)
			i++
		}
	}
}

// PrintTreeInOrder возвращает строку с элементами дерева в порядке возрастания
func PrintTreeInOrder(root *TreeNode) string {
	var result string
	printTreeInOrder(root, &result)
	return result
}

func printTreeInOrder(root *TreeNode, result *string) {
	if root != nil {
		if root.Parent != nil {
			*result += fmt.Sprintf("%d --> %d\n", root.Parent.Key, root.Key)
		}
		printTreeInOrder(root.Left, result)
		printTreeInOrder(root.Right, result)
	}
}

// BuildTreeFromInput строит сбалансированное бинарное дерево из входных данных
func BuildTreeFromInput(input string) (*TreeNode, string) {
	var keys []int
	addedKeys := make(map[int]bool)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " --> ")

		if len(parts) == 2 {
			key, _ := strconv.Atoi(parts[1])
			keys = append(keys, key)
			addedKeys[key] = true
		}
	}

	// Сортируем ключи
	sort.Ints(keys)

	// Строим сбалансированное бинарное дерево
	root := buildBalancedTree(keys, nil, addedKeys)

	// Выводим бинарное сбалансированное дерево в порядке возрастания в строку
	result := `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---
{{<// mermaid //>}}
` + PrintTreeInOrder(root) + `{{<// /mermaid //>}}`

	return root, result
}

func buildBalancedTree(keys []int, parent *TreeNode, addedKeys map[int]bool) *TreeNode {
	if len(keys) == 0 {
		return nil
	}

	// Находим середину массива
	mid := len(keys) / 2

	// Создаем узел для середины массива
	node := &TreeNode{Key: keys[mid], Parent: parent}
	addedKeys[keys[mid]] = true

	// Рекурсивно строим левое и правое поддерево
	node.Left = buildBalancedTree(keys[:mid], node, addedKeys)
	node.Right = buildBalancedTree(keys[mid+1:], node, addedKeys)

	return node
}

// ResultAfterAdding возвращает результат после добавления 5 случайных элементов
func ResultAfterAdding(root *TreeNode) string {
	addedKeys := make(map[int]bool)
	addElementsToTree(&root, addedKeys)

	// Вывод обновленного бинарного дерева в порядке возрастания
	resultAfterAdding := PrintTreeInOrder(root)

	return `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---
{{< mermaid >}}
graph TD
` +
		resultAfterAdding + `{{< /mermaid >}}`
}

func main() {
	input := `---
menu:
  after:
    name: binary_tree
    weight: 2
title: Построение сбалансированного бинарного дерева
---
{{< mermaid >}}
graph TD
620 --> 139
{{< /mermaid >}}`

	// Строим сбалансированное бинарное дерево из входных данных
	root, _ := BuildTreeFromInput(input)

	// Получаем результат после добавления 5 случайных элементов
	resultAfterAdding := ResultAfterAdding(root)

	// Выводим результат
	fmt.Println(resultAfterAdding)
}

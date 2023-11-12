package tree

import (
	"fmt"
	"math/rand"
	"strings"
)

type Node struct {
	ID    int
	Name  string
	Form  string
	Links []*Node
}

func CreateGraph() []*Node {
	numNodes := rand.Intn(26) + 5
	nodes := make([]*Node, numNodes)

	for i := 0; i < numNodes; i++ {
		nodes[i] = &Node{
			ID:    i + 1,
			Name:  fmt.Sprintf("Node%d", i+1),
			Form:  getRandomForm(),
			Links: []*Node{},
		}
	}

	// Создаем связи с вероятностью 0.2 для каждой пары узлов
	for i := 0; i < numNodes; i++ {
		for j := i + 1; j < numNodes; j++ {
			if rand.Float64() < 0.05 || len(nodes[i].Links) == 0 {
				nodes[i].Links = append(nodes[i].Links, nodes[j])
			} else if len(nodes[j].Links) == 0 && rand.Float64() < 0.5 {
				nodes[j].Links = append(nodes[j].Links, nodes[i])
			}
		}
	}

	return nodes
}

func getRandomForm() string {
	forms := []string{"((Circle))", "[Square rect]", "(Round Rect)", "{Rhombus}"}
	return forms[rand.Intn(len(forms))]
}

func GenerateDOTFormat(graph []*Node) string {
	var builder strings.Builder
	builder.WriteString("graph LR\n")

	for _, node := range graph {
		builder.WriteString(fmt.Sprintf("    %s%s\n", node.Name, node.Form))
		for _, link := range node.Links {
			builder.WriteString(fmt.Sprintf("    %s --> %s\n", node.Name, link.Name))
		}
	}

	return `---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---
{{< mermaid >}}` + builder.String() + `
{{< /mermaid >}}`
}

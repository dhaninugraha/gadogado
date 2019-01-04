package gadogado

import (
	"github.com/dhaninugraha/gadogado/internal/utils"
	"golang.org/x/net/html"
	"io"
)

var typeMap = map[html.NodeType]string{
    html.ErrorNode: "ErrorNode",
    html.TextNode: "TextNode",
    html.DocumentNode: "DocumentNode",
    html.ElementNode: "ElementNode",
    html.CommentNode: "CommentNode",
    html.DoctypeNode: "DoctypeNode",
}

type Node struct {
	Tag			string				`json:"tag,omitempty"`
	NodeType	string				`json:"node_type"`
	Attrs		map[string]string	`json:"attrs,omitempty"`
	Text		string				`json:"text,omitempty"`
	Children	[]Node				`json:"children,omitempty"`
}

func newNode() Node {
	return Node{Children: []Node{}, Attrs: map[string]string{}}
}

func iterateNodes(n *html.Node, parent *Node, excludedTags *dummyMap) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		parent.Tag = n.Data
		parent.NodeType = typeMap[n.Type]

		attrs := make(map[string]string)
		for _, a := range n.Attr {
			attrs[a.Key] = a.Val
		}
		parent.Attrs = attrs

		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			trimmed := utils.TrimAll(n.FirstChild.Data)
			if trimmed != "" {
				parent.Text = n.FirstChild.Data
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			if !excludedTags.exists(child.Data) {
				childNode := newNode()
				parent.Children = append(parent.Children, childNode)
				iterateNodes(child, &parent.Children[len(parent.Children) - 1], excludedTags)
			}
		}
	}
}

func ExcludeTags(tags ...string) *dummyMap {
	excluded := newDummyMap()

	for _, tag := range tags {
		excluded.addKey(tag)
	}

	return excluded
}

func Make(r io.Reader, excludedTags *dummyMap) (*Node, error) {
	doc, err := html.Parse(r)
	node := newNode()
	if err == nil {
		iterateNodes(doc, &node, excludedTags)

		if len(node.Children) > 0 {
			return &node, nil
		}
	}

	return &node, err
}
package model

import (
	"path"
	"path/filepath"
	"sort"
	"strings"
)

const (
	// TreeIndex is index page tree node
	TreeIndex = "index"
	// TreePost is a post node
	TreePost = "post"
	// TreePage is a page node with content
	TreePage = "page"
	// TreePageNode is a empty page as tree node, not read page
	TreePageNode = "page-node"
	// TreeArchive is node of archive page
	TreeArchive = "archive"
	// TreePostList is node of list page of posts
	TreePostList = "post-list"
	// TreePostTag is node of list posts belongs to a tag
	TreePostTag = "post-tag"
	// TreeTag is node of tag page, no used now
	TreeTag = "tag"
	// TreeXML is xml file node
	TreeXML = "xml"
	// TreeDir is a directory node, use to mark directory
	TreeDir = "dir"
)

// Tree describe the position of one file in all compiled files
type Tree struct {
	Title    string
	Link     string
	I18n     string
	Type     string
	Sort     int
	Dest     string
	URL      string
	children []*Tree
	parent   *Tree
}

type treeSlice []*Tree

// implement sort.Sort interface
func (t treeSlice) Len() int           { return len(t) }
func (t treeSlice) Less(i, j int) bool { return t[i].Sort < t[j].Sort }
func (t treeSlice) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

// NewTree create new tree node records with dest prefix string
func NewTree(dest string) *Tree {
	return &Tree{
		Link: "/",
		Type: TreeIndex,
		I18n: "tree",
		Sort: 0,
		Dest: dest,
		URL:  "",
	}
}

// FullURL return full url of this node from top
func (t *Tree) FullURL() string {
	if t.Type == TreePageNode {
		return "#"
	}
	parents := t.Parents()
	u := ""
	for _, p := range parents {
		u = path.Join(u, p.Link)
	}
	return path.Join(u, t.Link)
}

// HasChildren return true if Tree's children is not empty
func (t *Tree) HasChildren() bool {
	return len(t.children) > 0
}

// Children return nodes of tree by url link in next child level, not all chilrens in all levels
func (t *Tree) Children(link ...string) []*Tree {
	if len(link) == 0 {
		return t.children
	}
	if len(link) == 1 && link[0] == "" {
		return t.children
	}
	if t2 := t.subTree(link[0]); t2 != nil {
		return t2
	}
	return nil
}

// Dirs returns nodes of dir nodes by link
func (t *Tree) Dirs(link ...string) []*Tree {
	children := t.Children(link...)
	if len(children) == 0 {
		return nil
	}
	dirs := []*Tree{}
	for _, c := range children {
		if c.Type == TreeDir {
			dirs = append(dirs, c)
		}
	}
	return dirs
}

// Nodes return nodes of page-node by link
func (t *Tree) Nodes(link ...string) []*Tree {
	children := t.Children(link...)
	if len(children) == 0 {
		return nil
	}
	nodes := []*Tree{}
	for _, c := range children {
		if c.Type == TreePageNode {
			nodes = append(nodes, c)
		}
	}
	sort.Sort(treeSlice(nodes))
	return nodes
}

// Child return the node by link,
// if not exist, return nil
func (t *Tree) Child(link ...string) *Tree {
	if len(link) == 0 || link[0] == "" {
		return t
	}
	if isSameURL(t.URL, link[0]) {
		return t
	}
	for _, c := range t.children {
		if isSameURL(c.URL, link[0]) {
			return c
		}
		if c2 := c.Child(link[0]); c2 != nil {
			return c2
		}
	}
	return nil
}

// Pages return nodes of page by link
func (t *Tree) Pages(link ...string) []*Tree {
	children := t.Children(link...)
	if len(children) == 0 {
		return nil
	}
	nodes := []*Tree{}
	for _, c := range children {
		if c.Type == TreePage {
			nodes = append(nodes, c)
		}
	}
	sort.Sort(treeSlice(nodes))
	return nodes
}

// Posts return nodes of post by link
func (t *Tree) Posts(link ...string) []*Tree {
	children := t.Children(link...)
	if len(children) == 0 {
		return nil
	}
	nodes := []*Tree{}
	for _, c := range children {
		if c.Type == TreePost {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

// IsValid return whether the node is compiled or not
func (t *Tree) IsValid() bool {
	return t.Type != ""
}

// Parent return parent node of this node
func (t *Tree) Parent() *Tree {
	return t.parent
}

// Parents return all parents of the node,
// begin at top
func (t *Tree) Parents() []*Tree {
	var t2 = t.Parent()
	if t2 == nil {
		return []*Tree{}
	}
	res := []*Tree{t2}
	for {
		t2 = t2.Parent()
		if t2 == nil {
			break
		}
		res = append([]*Tree{t2}, res...)
	}
	return res
}

func (t *Tree) subTree(link string) []*Tree {
	//link = strings.TrimSuffix(link, path.Ext(link))
	linkData := strings.Split(strings.Trim(link, "/"), "/")
	if len(linkData) == 0 {
		return nil
	}
	if len(linkData) == 1 && linkData[0] == "" {
		return t.children
	}
	for _, c := range t.children {
		if c.Link == linkData[0] {
			if len(linkData) == 1 {
				if path.Ext(linkData[0]) != "" {
					return []*Tree{c}
				}
			}
			return c.subTree(strings.Join(linkData[1:], "/"))
		}
	}
	return nil
}

// Print print tree nodes as readable string
func (t *Tree) Print(prefix string) {
	println(prefix+"/"+t.Link, t.I18n, t.Title, "@"+t.Type, t.URL)
	for _, c := range t.children {
		c.Print(prefix + "---")
	}
}

// Add add tree node
func (t *Tree) Add(link, title, linkType string, s int) {
	link = filepath.ToSlash(link)
	link = strings.TrimPrefix(link, t.Dest+"/")
	linkData := strings.SplitN(link, "/", 2)
	if len(linkData) == 0 {
		return
	}
	isFind := false
	for _, c := range t.children {
		if c.Link == linkData[0] {
			isFind = true
			if linkType == TreePageNode && len(linkData) == 2 && linkData[1] == "" {
				c.Type = linkType
				c.Title = title
				c.Sort = s
				break
			}
			if linkType == TreePageNode && len(linkData) == 1 && linkData[0] == link {
				c.Type = linkType
				c.Title = title
				c.Sort = s
				break
			}
			if len(linkData) > 1 {
				c.Add(strings.Join(linkData[1:], "/"), title, linkType, s)
			}
			break
		}
	}
	if !isFind {
		tree := &Tree{
			Title:  title,
			Link:   linkData[0],
			I18n:   t.I18n + "." + linkData[0],
			Type:   linkType,
			Sort:   s,
			URL:    path.Join(t.URL, linkData[0]),
			parent: t,
		}
		if len(linkData) > 1 {
			tree.Type = TreeDir
			tree.Title = ""
			if linkType == TreePageNode {
				tree.Type = linkType
				tree.Title = title
				tree.Sort = s
			}
			if linkData[1] != "" {
				tree.Add(strings.Join(linkData[1:], "/"), title, linkType, s)
			}
		}
		t.children = append(t.children, tree)
	}
	sort.Sort(treeSlice(t.children))
}

func isSameURL(url1, url2 string) bool {
	if url1 == url2 {
		return true
	}
	url1 = strings.Trim(url1, "/")
	url2 = strings.Trim(url2, "/")
	return url1 == url2
}

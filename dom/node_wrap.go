package dom

// Auto-generated by internal/cmd/genwrapnode/genwrapnode.go. DO NOT EDIT!

import (
	"fmt"
	"unsafe"

	"github.com/lestrrat-go/libxml2/clib"
	"github.com/lestrrat-go/libxml2/types"
)

func wrapNamespaceNode(ptr unsafe.Pointer) *Namespace {
	var n Namespace
	n.ptr = ptr
	return &n
}

func wrapAttributeNode(ptr unsafe.Pointer) *Attribute {
	var n Attribute
	n.ptr = ptr
	return &n
}

func wrapCDataSectionNode(ptr unsafe.Pointer) *CDataSection {
	var n CDataSection
	n.ptr = ptr
	return &n
}

func wrapCommentNode(ptr unsafe.Pointer) *Comment {
	var n Comment
	n.ptr = ptr
	return &n
}

func wrapElementNode(ptr unsafe.Pointer) *Element {
	var n Element
	n.ptr = ptr
	return &n
}

func wrapTextNode(ptr unsafe.Pointer) *Text {
	var n Text
	n.ptr = ptr
	return &n
}

func wrapPiNode(ptr unsafe.Pointer) *Pi {
	var n Pi
	n.ptr = ptr
	return &n
}

// WrapNode is a function created with the sole purpose of allowing
// go-libxml2 consumers that can generate a C.xmlNode pointer to
// create libxml2.Node types, e.g. go-xmlsec.
func WrapNode(n unsafe.Pointer) (types.Node, error) {
	switch typ := clib.XMLGetNodeTypeRaw(n); typ {
	case clib.AttributeNode:
		return wrapAttributeNode(n), nil
	case clib.CDataSectionNode:
		return wrapCDataSectionNode(n), nil
	case clib.CommentNode:
		return wrapCommentNode(n), nil
	case clib.ElementNode:
		return wrapElementNode(n), nil
	case clib.TextNode:
		return wrapTextNode(n), nil
	case clib.PiNode:
		return wrapPiNode(n), nil
	default:
		return nil, fmt.Errorf("unknown node: %d", typ)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Node in a binary tree
type Node struct {
	key   string
	value int
	nodeL *Node
	nodeR *Node
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func check(err error) {
	if err != nil {
		log.Fatal("Erreur:", err)
	}
}

// Create a new file whose the name is pass a parameter
func createFile(filename string) {
	if fileExists(filename) == false {
		file, err := os.Create(filename)
		check(err)
		fmt.Println("File", file)
	}
}

// Create a hashmap containing as key a character and as value the number of occurrence
func countOccurrences(msg string) map[string]int {
	hashMap := make(map[string]int)
	tChar := strings.Split(msg, "")
	for _, char := range tChar {
		if hashMap[char] == 0 {
			hashMap[char] = 1
		} else {
			hashMap[char]++
		}
	}
	return hashMap
}

func sortNodeByValue(tNode []Node) {
	sort.SliceStable(tNode, func(i, j int) bool {
		return tNode[i].value < tNode[j].value
	})
}

func createHuffmanTree(HashMap map[string]int) *Node {
	if len(HashMap) == 0 {
		return nil
	}
	var tNode []Node
	// Create Node object with HashMap
	for key, value := range HashMap {
		tNode = append(tNode, Node{key, value, nil, nil})
	}
	sortNodeByValue(tNode)
	for len(tNode) > 1 {
		node1, node2 := tNode[0], tNode[1]
		parent := &Node{key: "", nodeL: nil, nodeR: nil, value: node1.value + node2.value}
		if node1.value < node2.value {
			parent.nodeL = &node1
			parent.nodeR = &node2
		} else {
			parent.nodeL = &node2
			parent.nodeR = &node1
		}
		tNode = append(tNode[:0], tNode[2:]...)
		tNode = append(tNode, *parent)
		sortNodeByValue(tNode)
	}
	return &tNode[0]
}
func traverseTree(NodeRoot *Node, HashMapBinary map[string]string, str string) map[string]string {
	if NodeRoot.nodeL != nil {
		HashMapBinary = traverseTree(NodeRoot.nodeL, HashMapBinary, str+"0")
	}
	if NodeRoot.nodeR != nil {
		HashMapBinary = traverseTree(NodeRoot.nodeR, HashMapBinary, str+"1")
	}
	if NodeRoot.key != "" {
		HashMapBinary[NodeRoot.key] = str
	}
	return HashMapBinary
}
func convertBinaryStringToUInt8(binaryString string) []uint8{
	var sUint8 []uint8
	str := ""
	for i, char := range binaryString {
		str += string(byte(char))
		if len(str) == 8 || len(binaryString)-1 == i {
			nb, err := strconv.ParseInt(str, 2, 32)
			check(err)
			newUint := uint8(nb)
			sUint8 = append(sUint8, newUint)
			str = ""
		}
	}
	return sUint8
}
func main() {
	fmt.Println("Huffman encoding...")
	var HashMap map[string]int
	HashMapBinary := make(map[string]string)
	filename := "huffman.bin"
	filenameCompare := "huffmanCompare.bin" // Compare the size of huffman message not encoded
	huffmanMessage := "aaaaabbbbbbbbbccccccccccccdddddddddddddeeeeeeeeeeeeeeeefffffffffffffffffffffffffffffffffffffffffffff"
	createFile(filename)
	createFile(filenameCompare)
	HashMap = countOccurrences(huffmanMessage)
	for key := range HashMap {
		HashMapBinary[key] = ""
	}
	NodeRoot := createHuffmanTree(HashMap)
	// Create a new Hashmap having for key the letter et for value the encoded letter
	HashMapBinary = traverseTree(NodeRoot, HashMapBinary, "")
	// replace each letter in huffmanMessage by his representation in the NodeRoot to get an encoded message
	var binaryString string
	for _, char := range huffmanMessage {
		binaryString += HashMapBinary[string(char)]
	}
	sUint8 := convertBinaryStringToUInt8(binaryString)
	fmt.Println("Binary Array: ", sUint8)
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	check(err)
	fCompare, err := os.OpenFile(filenameCompare, os.O_WRONLY, 0644)
	check(err)
	wFile, err := f.Write(sUint8)
	check(err)
	wFileCompare, err := fCompare.WriteString(huffmanMessage)
	check(err)
	fmt.Println("wFile", wFile)
	fmt.Println("wFileCompare", wFileCompare)
}

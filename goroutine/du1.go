package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	verbose := flag.Bool("v", false, "show verbose progress messages.")

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 定期输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUssage(nfiles, nbytes)
		}
		printDiskUssage(nfiles, nbytes)
	}

	// 输出结果
	//var nfiles, nbytes int64
	//for size := range fileSizes {
	//	nfiles++       // 多少个文件
	//	nbytes += size // 文件大小
	//}
	//printDiskUssage(nfiles, nbytes)
}

func printDiskUssage(nfiles, nbytes int64) {
	var d string
	var z float64
	switch n := float64(nbytes); {
	case n/1e9 > 1:
		d = "GB"
		z = n / 1e9
	case n/1e6 > 1:
		d = "MB"
		z = n / 1e6
	case n/1e3 > 1:
		d = "KB"
		z = n / 1e3
	default:
		d = "B"
		z = n
	}
	fmt.Printf("%d files %.3f %s\n", nfiles, z, d)
}

// 递归 如果是目录就执行递归，否则把文件大小发送给 fileSizes通道
// root, &n, fileSizes
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 读取路径的文件对象
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

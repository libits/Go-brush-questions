package main

type Port int32
type umap map[string]user1

func (um *umap) del(key string) {
	delete(*um, key)
}

func main() {

}

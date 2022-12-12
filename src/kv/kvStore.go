package kv

type KvStore struct {
	storage Storage
}

func GetDataPath(path string) string {
	dataPathBuf := path + STORAGE_FILE_PREFIX + ".data"
	return dataPathBuf
}

func OpenKvStore(path string) (*KvStore, error) {
	storage, err := OpenSimplifiedBitcask(path)
	if err != nil {
		return nil, err
	}
	return &KvStore{storage}, nil
}

func (s *KvStore) Get(key string) (string, error) {
	return s.storage.get(key)
}

func (s *KvStore) Set(key string, val string) error {
	return s.storage.put(key, val)
}

func (s *KvStore) Remove(key string) error {
	return s.storage.remove(key)
}

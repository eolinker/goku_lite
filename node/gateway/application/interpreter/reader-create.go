package interpreter

import (
	"bytes"
	"strconv"
	"strings"
)

//ReaderCreateFunc readerCreateFunc
type ReaderCreateFunc func(head []byte, body []byte) (Reader, error)

const (
	//Body body
	Body = "body"
	//Header header
	Header = "header"
	//Restful restful
	Restful = "restful"
	//Query query
	Query = "query"
	//Cookie cookie
	Cookie = "cookie"
)

var (
	creators     map[string]ReaderCreateFunc
	readers      map[string][]byte
	pathSplitSeq = []byte(".")
)

func init() {
	creators = make(map[string]ReaderCreateFunc)
	readers = make(map[string][]byte)

	readers[Body[:4]] = []byte(Body)
	readers[Header[:4]] = []byte(Header)
	readers[Restful[:4]] = []byte(Restful)
	readers[Cookie[:4]] = []byte(Cookie)
	readers[Query[:4]] = []byte(Query)

	creators[Body] = ReaderCreateFunc(func(head []byte, body []byte) (Reader, error) {

		index := 0
		if len(head) > len(Body) {
			indexData := head[len(Body):]
			i, err := strconv.Atoi(string(indexData))
			if err != nil {
				return nil, err
			}
			index = i
		}

		pathData := bytes.Split(body, pathSplitSeq)

		path := make([]string, 0, len(pathData))
		for _, p := range pathData {
			path = append(path, string(p))
		}

		return &_BodyReader{
			Index: index,
			Path:  path,
			Name:  string(body),
		}, nil
	})
	creators[Header] = ReaderCreateFunc(func(head []byte, body []byte) (Reader, error) {
		index := 0
		if len(head) > len(Header) {
			indexData := head[len(Header):]
			i, err := strconv.Atoi(string(indexData))
			if err != nil {
				return nil, err
			}
			index = i
		}

		return &_HeaderReader{
			Index: index,
			Key:   string(body),
		}, nil
	})
	creators[Query] = ReaderCreateFunc(func(head []byte, body []byte) (Reader, error) {
		if !bytes.Equal(head, []byte(Query)) {
			return nil, GrammarError(string(head))
		}
		return &_QueryReader{
			Key: string(body),
		}, nil
	})
	creators[Restful] = ReaderCreateFunc(genResfult)
	creators[Cookie] = ReaderCreateFunc(func(head []byte, body []byte) (Reader, error) {
		index := 0
		if len(head) > len(Cookie) {
			indexData := head[len(Cookie):]
			i, err := strconv.Atoi(string(indexData))
			if err != nil {
				return nil, err
			}
			index = i
		}

		return &_CookieReader{
			Index: index,
			Name:  string(body),
		}, nil
	})
}

func genResfult(head []byte, body []byte) (Reader, error) {

	if !bytes.Equal(head, []byte(Restful)) {
		return nil, GrammarError(string(head))
	}
	return &_RestFulReader{
		Key: string(body),
	}, nil
}

func genReader(line []byte) (Reader, error) {

	kindex := bytes.IndexAny(line, ".")

	if kindex == -1 {
		return nil, GrammarError(line)
	}

	key := line[:kindex]
	keyPre := strings.ToLower(string(key[:4]))

	cmd, has := readers[keyPre]
	if !has {
		return nil, GrammarError(string(line))
	}
	if !bytes.HasPrefix(key, cmd) {
		return nil, GrammarError(string(line))
	}

	create := creators[string(cmd)]

	return create(bytes.ToLower(key), line[kindex+1:])

}

package gomock

import "jasper/test/gomock/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}

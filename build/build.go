package build

// go build -ldflags "-X 'github.com/Anveena/ezTools/build.GitInfo=$(git show -s --format='git版本:%H%n提交者:%an%n提交信息:%s')' -X 'github.com/Anveena/ezTools/build.Date=$(date +'%Y年%m月%d日 %T')'"
var (
	GitInfo = ""
	Date    = ""
)

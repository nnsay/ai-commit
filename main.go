package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// 获取 scope 名称
func getScopeNames() string {
	scopes := ""
	_, err := os.Stat("./nx.json")
	if err == nil {
		scopeCmd := exec.Command("npx", "nx", "show", "projects")
		scopeOut, err := scopeCmd.Output()
		if err != nil {
			return scopes
		}
		scopes = strings.ReplaceAll(string(scopeOut), "\n", " ")
	}
	return scopes
}

func debugLog(msg string) {
	isDebug := len(os.Getenv("DEBUG")) > 0
	if isDebug {
		fmt.Printf("[DEBUG] %s", msg)
	}
}

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	diffArgs := []string{"diff"}
	diffArgs = append(diffArgs, os.Args[1:]...)
	diffCom := exec.Command("git", diffArgs...)
	diffOut, _ := diffCom.Output()
	diffText := string(diffOut)

	if len(diffText) == 0 {
		fmt.Println("[ERROR] git diff result is empty")
		return
	}
	// question := fmt.Sprintf("请根据下面的git diff结果, 结合nx monorepo的代码组织特点, 编写一条符合约定式提交规则的commit信息: \n%s", diffText)
	restriction1 := "1. 符合约定式提交的格式, 如: <type>: <description>\n  其中type的合法值有:fix feat chore docs build ci style refactor perf test"
	names := getScopeNames()
	if len(names) > 0 {
		restriction1 = fmt.Sprintf(`1. 符合约定式提交的格式, 如: <type>(scope): <description>
	1.1 其中type的合法值有:fix feat chore docs build ci style refactor perf test
	1.2 其中scope的合法值有: %s
	1.3 可能有多个scope时取变更文件最多的scope`,
			names)
	}
	question := fmt.Sprintf(`请根据下面的git diff结果, 编写一条commit信息, 编写要求有:
%s
2. 英文小写;
3. 长度不超过20个单词.
git diff 结果如下:
%s
`, restriction1, diffText)
	debugLog(fmt.Sprintf("question:\n%s\n", question))
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
	}
	commitMsg := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
	fmt.Printf("AI Commit Message Result:\n	%s\n", commitMsg)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Above message is ok? continue or put your custom message at here")
	userInput, _ := reader.ReadString('\n')
	if userInput != "" && userInput != "y" && userInput != "yes" && userInput != "\n" {
		commitMsg = userInput
	}
	debugLog(fmt.Sprintf("the commit mssage: %s\n", commitMsg))
	changeCmd := exec.Command("git", "diff", "--name-only", "--staged")
	changeFiles, _ := changeCmd.Output()
	if len(changeFiles) == 0 {
		debugLog("execute git add because no changes")
		addCmd := exec.Command("git", "add", ".")
		addCmd.Run()
	}
	commitCmd := exec.Command("git", "commit", "-m", commitMsg)
	commitCmd.Run()
}
